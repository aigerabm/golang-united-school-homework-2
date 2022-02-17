package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type SideNumType int

const (
	SidesTriangle SideNumType = 3
	SidesSquare   SideNumType = 4
	SidesCircle   SideNumType = 0
)

func CalcSquare(sideLen float64, sidesNum SideNumType) float64 {
	if sidesNum == SidesCircle {
		return math.Pow(sideLen, 2) * math.Pi
	} else if sidesNum == SidesTriangle {
		return math.Sqrt(math.Pow(sideLen, 2)-math.Pow(sideLen/2, 2)) * sideLen / 2
	} else if sidesNum == SidesSquare {
		return math.Pow(sideLen, 2)
	} else {
		return 0
	}
}
