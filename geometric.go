package gogression
import "math"

type Geometic struct {
	First int64
	Ratio int64
}

func NewGeometic(g, r int64) *Geometic {
	return &Geometic{g, r}
}

func (g Geometic) Item(n int64) int64 {
	return g.First * int64(math.Pow(float64(g.Ratio), float64(n - int64(1))))
}

func (g Geometic) Sum(n int64) int64 {
	return (g.First*(int64(1) - int64(math.Pow(float64(g.Ratio), float64(n)))))/(int64(1)-g.Ratio)
}
