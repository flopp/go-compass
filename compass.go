package compass

import "math"

type Direction int

const (
	N   Direction = 0
	NNE Direction = 1
	NE  Direction = 2
	ENE Direction = 3
	E   Direction = 4
	ESE Direction = 5
	SE  Direction = 6
	SSE Direction = 7
	S   Direction = 8
	SSW Direction = 9
	SW  Direction = 10
	WSW Direction = 11
	W   Direction = 12
	WNW Direction = 13
	NW  Direction = 14
	NNW Direction = 15
)

func (d Direction) String() string {
	switch d {
	case N:
		return "N"
	case NNE:
		return "NNE"
	case NE:
		return "NE"
	case ENE:
		return "ENE"
	case E:
		return "E"
	case ESE:
		return "ESE"
	case SE:
		return "SE"
	case SSE:
		return "SSE"
	case S:
		return "S"
	case SSW:
		return "SSW"
	case SW:
		return "SW"
	case WSW:
		return "WSW"
	case W:
		return "W"
	case WNW:
		return "WNW"
	case NW:
		return "NW"
	case NNW:
		return "NNW"
	}
	return "???"
}

type Resolution int

const (
	Resolution4  Resolution = 0
	Resolution8  Resolution = 1
	Resolution16 Resolution = 2
)

func normalizeAngle(angleDeg float64) float64 {
	angleDeg = math.Mod(angleDeg, 360.0)
	if angleDeg < 0 {
		angleDeg += 360.0
	}
	return angleDeg
}

// GetDirection returns the nearest compass direction for the given angle (in degrees);
// "resolution" selects the set of considered directions:
// * Resolution4 => return one of N, E, S, W
// * Resolution8 => return one of N, NE, E, SE, S, SW, W, NW
// * Resolution16 => return one of N, NNE, NE, ENE, E, ESE, SE, SSE, S, SSW, SW, WSW, W, WNW, NW, NNW
func GetDirection(angleDeg float64, resolution Resolution) Direction {
	switch resolution {
	case Resolution4:
		return Direction(4 * int(normalizeAngle(angleDeg+45.0)/90.0))
	case Resolution8:
		return Direction(2 * int(normalizeAngle(angleDeg+22.5)/45.0))
	case Resolution16:
		return Direction(int(normalizeAngle(angleDeg+11.25) / 22.5))
	}
	return Direction(int(normalizeAngle(angleDeg+11.25) / 22.5))
}
