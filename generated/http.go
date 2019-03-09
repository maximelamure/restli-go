package generated

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

func get(ctx context.Context, URI string, timeout time.Duration, codec reader, data interface{}) error {
	req, errGet := http.NewRequest("GET", URI, nil)
	if errGet != nil {
		return errors.New("unexpected error creating NewRequest GET " + URI + " " + errGet.Error())
	}

	req.Header = getRequestHeaders(ctx, codec)

	resp, errDo := do(ctx, req, timeout, 1)
	if errDo != nil {
		return errors.New("unexpected error executing GET " + URI + " " + errDo.Error())
	}

	defer resp.Body.Close()
	if err := codec.Read(resp.Body, &data); err != nil {
		return err
	}

	return nil
}

func getRequestHeaders(ctx context.Context, codec reader) http.Header {
	headers := http.Header{}
	headers.Set("X-RestLi-Protocol-Version", "2.0.0")
	headers.Set("Content-Type", codec.ContentType())

	return headers
}

func do(ctx context.Context, req *http.Request, timeout time.Duration, retry int) (resp *http.Response, err error) {
	var c http.Client
	c.Timeout = timeout
	retryTimeout := timeout / 2

	for attempt := 0; attempt <= retry; attempt++ {
		resp, err = c.Do(req)
		if err != nil {
			c.Timeout = retryTimeout
			// In go 1.6, replace the switch with `if err, ok := err.(net.Error); ok && err.Timeout() { ... }
			switch err := err.(type) {
			case *url.Error:
				if netErr, ok := err.Err.(net.Error); ok && netErr.Timeout() {
					log.Fatal(netErr, "retry ", req.URL)
				} else {
					return nil, err
				}
			case net.Error:
				if err.Timeout() {
					log.Fatal(err, "retry ", req.URL)
				} else {
					return nil, err
				}
			default:
				return nil, err
			}
		} else {
			break
		}
	}
	return
}
