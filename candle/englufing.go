package candle

import "github.com/devmsa-cl/go-candle-patterns/types"

// IsBullishEngulfed check if the current candle is engulfed by the previous one
// For a candle to be bullish engulfed, the following conditions must be met:
// 1. The previous candle is bearish.
// 2. The current candle is bullish.
// 3. The top of the current candle is higher than the top of the previous candle.
// 4. The bottom of the current candle is lower than the bottom of the previous candle.
func IsBullishEngulfed(previousCandle types.Candle, currentCandle types.Candle) bool {

	currentTop, currentBottom := BodyEnd(currentCandle)
	previousTop, previousBottom := BodyEnd(previousCandle)

	if IsBearishCandle(previousCandle) &&
		IsBullishCandle(currentCandle) &&
		currentTop > previousTop &&
		currentBottom < previousBottom {
		return true
	}

	return false
}

// IsBearishEngulfed check if the current candle is engulfed by the previous one
// For a candle to be Bearish engulfed, the following conditions must be met:
// 1. The previous candle is Bullish.
// 2. The current candle is Bearish.
// 3. The top of the current candle is higher than the top of the previous candle.
// 4. The bottom of the current candle is lower than the bottom of the previous candle.
func IsBearishEngulfed(previousCandle types.Candle, currentCandle types.Candle) bool {

	currentTop, currentBottom := BodyEnd(currentCandle)
	previousTop, previousBottom := BodyEnd(previousCandle)

	if IsBullishCandle(previousCandle) &&
		IsBearishCandle(currentCandle) &&
		currentTop < previousTop &&
		currentBottom < previousBottom {
		return true
	}

	return false
}
