package square

import (
	"math"
)

type sidesType int

const SidesCircle sidesType = 0
const SidesTriangle sidesType = 3
const SidesSquare sidesType = 4

func CalcSquare(sideLen float64, sidesNum sidesType) float64 {
	switch sidesNum {
	case SidesCircle:
		return math.Pi * math.Pow(sideLen, 2)

	case SidesTriangle:
		return math.Pow(sideLen, 2) * math.Sqrt(3) / 4

	case SidesSquare:
		return math.Pow(sideLen, 2)

	default:
		return 0.0
	}
}
