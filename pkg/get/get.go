package get

import "github.com/safeguardproperties/gomodels/cache"

//Getter has method to get value for given primaryContext, secondaryContext, key
type Getter interface {
	Get(primaryContext string, secondaryContext string, key string) *cache.Result
}

//Client
type Client struct {
}

//New
func New() *Client {
	return &Client{}
}

//Get retrieves cached value for provided primaryContext, secondaryContext, key
func (c *Client) Get(primaryContext string, secondaryContext string, key string) *cache.Result {
	return &cache.Result{}
}
