package client

import (
	"encoding/json"
)

type Request struct {
	next []*Request
}

func NewRequest(next ...*Request) {

}

func (r *Request) then(next ...*Request) {
	r.next = next
}

type Response struct {
	Data  json.RawMessage
	Error error
}
