package getset

import (
	"github.com/divyag9/cache/pkg/get"
	"github.com/divyag9/cache/pkg/set"
)

//GetSet
type GetSet struct {
	Getter get.Getter
	Setter set.Setter
}

//GetAndSet
func (gs *GetSet) GetAndSet(primaryContext string, secondaryContext string, key string, expiry int64, valueGetter func()) {

}
