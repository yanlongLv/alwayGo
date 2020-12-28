package rolling

import (
	"fmt"
)

//Iterator ..
type Iterator struct {
	count          int
	interatedCount int
	cur            *Bucket
}

//Next ..
func (i *Iterator) Next() bool {
	return i.count != i.interatedCount
}

//Bucket ...
func (i *Iterator) Bucket() Bucket {
	if !(i.Next()) {
		panic(fmt.Errorf("stat/metric : iteration out of range iteratedCount:%d count %d,", i.interatedCount, i.count))
	}
	bucket := *i.cur
	i.interatedCount++
	i.cur = i.cur.Next()
	return bucket
}
