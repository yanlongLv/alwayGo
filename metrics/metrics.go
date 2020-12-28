package metrics

//Metric ..
type Metric interface {
	Add(int64)
	Value() int64
}

//Opts ...
type Opts struct{}
