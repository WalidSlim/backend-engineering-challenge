package movingaverage

type MovingAverage struct {
	windowSize   int
	data         []float64
	sum          float64
	nonZeroCount int
}

// NewMovingAverage creates a new MovingAverage instance with the specified window size.
func NewMovingAverage(windowSize int) *MovingAverage {
	return &MovingAverage{
		windowSize:   windowSize,
		data:         make([]float64, 0, windowSize),
		sum:          0.0,
		nonZeroCount: 0,
	}
}

// AddValue adds a new data point to the moving average and returns the updated average.
func (ma *MovingAverage) AddValue(value float64) float64 {
	if len(ma.data) == ma.windowSize {
		// Remove the oldest value from the sum when the window is full.
		ma.sum -= ma.data[0]
		if ma.data[0] != 0 {
			ma.nonZeroCount -= 1
		}
		ma.data = ma.data[1:]
	}

	// Add the new value to the data slice and update the sum.
	ma.data = append(ma.data, value)
	ma.sum += value
	if value != 0 {
		ma.nonZeroCount += 1
	}
	//Adding 1 in case non zero count = 0 for the sake of avoiding 0 division
	var denominator int
	if ma.nonZeroCount == 0 {
		denominator = 1
	} else {
		denominator = ma.nonZeroCount
	}
	return ma.sum / float64(denominator)
}
