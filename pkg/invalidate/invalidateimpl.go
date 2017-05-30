package invalidate

import "github.com/safeguardproperties/gomodels/cache"

//Impl
type Impl struct {
	Invalidator Invalidator
}

//InvalidatePriSec
func (i *Impl) InvalidatePriSec(primaryContext string, secondaryContext string) *cache.Result {
	return &cache.Result{}
}

//InvalidatePri
func (i *Impl) InvalidatePri(primaryContext string) *cache.Result {
	return &cache.Result{}
}

//Invalidate
func (i *Impl) Invalidate() *cache.Result {
	return &cache.Result{}
}
