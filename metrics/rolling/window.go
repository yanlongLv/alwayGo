package rolling

//Bucket ..
type Bucket struct {
	Points []float64
	Count  int64
	next   *Bucket
}

//Append ...
func (b *Bucket) Append(val float64) {
	b.Points = append(b.Points, val)
	b.Count++
}

//Add ..
func (b *Bucket) Add(offset int, val float64) {
	b.Points[offset] += val
	b.Count++
}

//Reset ...
func (b *Bucket) Reset() {
	b.Points = b.Points[:0]
	b.Count = 0
}

//Next ....
func (b *Bucket) Next() *Bucket {
	return b.next
}
