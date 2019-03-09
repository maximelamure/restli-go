package client

import (
	"context"

	"github.com/maximelamure/restli-go/common"
)

type Client interface {
	Load(requests ...Request)
	Fetch() Result
	FindOrders(ctx context.Context, orderID int, buyerID int, opts ...common.CallOption) *Request
	FindUsers(ctx context.Context, companyID int, opts ...common.CallOption) *Request
}

type profileClient struct {
}

func NewClient(ctx context.Context) Client {
	return &profileClient{}
}

func (p *profileClient) Load(requests ...Request) {

}

func (p *profileClient) Fetch() Result {
	return Result{}
}

func (p *profileClient) FindOrders(ctx context.Context, orderID int, buyerID int, opts ...common.CallOption) *Request {
	//todo: save the filters
	return &Request{}
}

func (p *profileClient) FindUsers(ctx context.Context, companyID int, opts ...common.CallOption) *Request {
	//todo: save the filters
	return &Request{}
}

// Simple functions
func FindOrders(ctx context.Context, orderID int, buyerID int, opts ...common.CallOption) ([]Order, error) {
	//todo: do the http call
	return make([]Order, 0), nil
}

func FindUsers(ctx context.Context, companyID int) ([]User, error) {
	//todo: do the http call
	return make([]User, 0), nil
}
