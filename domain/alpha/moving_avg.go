package alpha

import "math"

type MovingAverage struct {
	// Immutable fields
	length int
	// Mutable fields
	runningSum       float64
	runningResSquare float64
	valueBuffer      []float64
	residualBuffer   []float64
}

func NewMovingAverage(length int) *MovingAverage {
	return &MovingAverage{
		length: length,

		runningSum:       0,
		runningResSquare: 0,
		valueBuffer:      make([]float64, 0, length),
		residualBuffer:   make([]float64, 0, length),
	}
}

func (ma *MovingAverage) Length() int {
	return ma.length
}

func (ma *MovingAverage) Append(value float64) {
	ma.valueBuffer = append(ma.valueBuffer, value)
	ma.runningSum += value

	if len(ma.valueBuffer) > ma.length {
		ma.runningSum -= ma.valueBuffer[0]
		ma.valueBuffer = ma.valueBuffer[1:]
	}
	// Calculate residual
	movingAverage := ma.runningSum / float64(len(ma.valueBuffer))
	resSquare := math.Pow(value-movingAverage, 2)

	ma.runningResSquare += resSquare
	if len(ma.residualBuffer) > ma.length {
		ma.runningResSquare -= ma.residualBuffer[0]
		ma.residualBuffer = ma.residualBuffer[1:]
	}
}

func (ma *MovingAverage) Mean() float64 {
	if len(ma.valueBuffer) == 0 {
		return 0
	}
	return ma.runningSum / float64(len(ma.valueBuffer))
}

func (ma *MovingAverage) Std() float64 {
	if len(ma.residualBuffer) == 0 {
		return 0
	}
	return math.Sqrt(ma.runningResSquare / float64(len(ma.residualBuffer)-1))
}
