package fig

import "math"

type figures int

const (
	square   figures = iota // квадрат
	circle                  // круг
	triangle                // равносторонний треугольник
)

func Area(fig figures) (func(float64) float64, bool) {
	switch fig {
	case square:
		return func(x float64) float64 {
			return x * x
		}, true
	case circle:
		return func(x float64) float64 {
			return math.Pi * x * x
		}, true
	case triangle:
		return func(x float64) float64 {
			return math.Sqrt(3) / 4 * x * x
		}, true
	}

	return nil, false
}
