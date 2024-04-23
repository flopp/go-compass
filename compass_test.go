package compass

import (
	"fmt"
	"testing"
)

func TestGetDirection(t *testing.T) {
	type Test struct {
		angle      float64
		resolution Resolution
		expected   Direction
	}
	for _, test := range []Test{
		{0, Resolution4, N},
		{90, Resolution4, E},
		{180, Resolution4, S},
		{270, Resolution4, W},

		{0, Resolution8, N},
		{45, Resolution8, NE},
		{90, Resolution8, E},
		{135, Resolution8, SE},
		{180, Resolution8, S},
		{225, Resolution8, SW},
		{270, Resolution8, W},
		{315, Resolution8, NW},

		{0, Resolution16, N},
		{22.5, Resolution16, NNE},
		{45, Resolution16, NE},
		{67.5, Resolution16, ENE},
		{90, Resolution16, E},
		{102.5, Resolution16, ESE},
		{135, Resolution16, SE},
		{157.5, Resolution16, SSE},
		{180, Resolution16, S},
		{202.5, Resolution16, SSW},
		{225, Resolution16, SW},
		{247.5, Resolution16, WSW},
		{270, Resolution16, W},
		{292.5, Resolution16, WNW},
		{315, Resolution16, NW},
		{337.5, Resolution16, NNW},

		{-90, Resolution4, W},
		{3600, Resolution4, N},
	} {
		// test main angle
		direction := GetDirection(test.angle, test.resolution)
		if direction != test.expected {
			t.Errorf("direction(%f, %v) = '%v', expected = '%v'", test.angle, test.resolution, direction, test.expected)
		}
		// test angle-1
		direction = GetDirection(test.angle-1, test.resolution)
		if direction != test.expected {
			t.Errorf("direction(%f-1, %v) = '%v', expected = '%v'", test.angle, test.resolution, direction, test.expected)
		}
		// test angle+1
		direction = GetDirection(test.angle+1, test.resolution)
		if direction != test.expected {
			t.Errorf("direction(%f+1, %v) = '%v', expected = '%v'", test.angle, test.resolution, direction, test.expected)
		}
	}
}

func TestGetDirectionRanges(t *testing.T) {
	for angle := 0.0; angle < 360.0; angle += 1.0 {
		d := GetDirection(angle, Resolution4)
		if d != N && d != E && d != S && d != W {
			t.Errorf("direction(%f, Level4) = %v, expected=N/E/S/W", angle, d)
		}

		d = GetDirection(angle, Resolution8)
		if d != N && d != NE && d != E && d != SE && d != S && d != SW && d != W && d != NW {
			t.Errorf("direction(%f, Level8) = %v, expected=N/NE/E/SE/S/SW/W/NW", angle, d)
		}

		d = GetDirection(angle, Resolution16)
		if d != N && d != NNE && d != NE && d != ENE && d != E && d != ESE && d != SE && d != SSE && d != S && d != SSW && d != SW && d != WSW && d != W && d != WNW && d != NW && d != NNW {
			t.Errorf("direction(%f, Level16) = %v, expected=N/NNE/NE/ENE/E/ESE/SE/SSE/S/SSW/SW/WSW/W/WNW/NW/NNW", angle, d)
		}
	}

	fmt.Printf("%v\n", GetDirection(13, Resolution4))
	fmt.Printf("%v\n", GetDirection(13, Resolution8))
	fmt.Printf("%v\n", GetDirection(13, Resolution16))
}
