package set

import (
	google_protobuf "github.com/golang/protobuf/ptypes/any"
	"github.com/safeguardproperties/gomodels/cache"
)

//Setter has method to set value for given primaryContext, secondaryContext, key
type Setter interface {
	Set(primaryContext string, secondaryContext string, key string, value *google_protobuf.Any, expiry int64) *cache.Result
}

//Client
type Client struct {
}

//New
func New() *Client {
	return &Client{}
}

//Set caches value for provided primaryContext, secondaryContext, key
func (c *Client) Set(primaryContext string, secondaryContext string, key string, value *google_protobuf.Any, expiry int64) *cache.Result {
	return &cache.Result{}
}
