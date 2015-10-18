package gogression

type Arithmetic struct {
	First      int64
	Difference int64
}

func NewArithmetic(a, n int64) *Arithmetic {
	return &Arithmetic{a, n}
}

func (a Arithmetic) Item(n int64) int64 {
	return a.First + (n-int64(1))*a.Difference
}

func (a Arithmetic) Sum(n int64) int64 {
	return n * (a.First + a.Item(n)) / 2
}
