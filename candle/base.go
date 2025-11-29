package candle

import (
	"math"

	"github.com/devmsa-cl/go-candle-patterns/types"
)

// new comment
func IsBearishCandle(candle types.Candle) bool {
	return candle.Close < candle.Open
}

func IsBullishCandle(candle types.Candle) bool {
	return candle.Close > candle.Open
}

func BodyLen(candle types.Candle) float64 {
	return math.Abs(candle.Close - candle.Open)
}
func CandleLen(candle types.Candle) float64 {
	return math.Abs(candle.High - candle.Low)
}

// BodyHalfLen returns the length of the body divided by 2.
func BodyHalfLen(candle types.Candle) float64 {
	return BodyLen(candle) / 2
}

/*
return the wick length
*/
func WickLen(candle types.Candle) float64 {
	return math.Abs(candle.High - math.Max(candle.Close, candle.Open))
}
func TailLen(candle types.Candle) float64 {
	return math.Abs(math.Min(candle.Close, candle.Open) - candle.Low)
}

/*
CandleWick returns the high and low of a candle.
The two values are returned as a tuple (high, low)
*/
func CandleWick(candle types.Candle) (float64, float64) {
	return candle.High, candle.Low
}

// BodyEnd returns the top and bottom of the candle.
//
// For bullish candles, the body is [top:close, bottom:open]
//
// For bearish candles, the body is [top:open, bottom:close]
// return two value (top, bottom)
func BodyEnd(c types.Candle) (float64, float64) {
	if IsBullishCandle(c) {
		return c.Close, c.Open
	}
	return c.Open, c.Close
}
