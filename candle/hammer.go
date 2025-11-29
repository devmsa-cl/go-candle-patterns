package candle

import "github.com/devmsa-cl/go-candle-patterns/types"

type HammerOptions struct {
	BodySizeThreshold float64 `json:"bodySizeThreshold,omitempty" validate:"omitempty,gt=0"`
	WickThreshold     float64 `json:"WickThreshold,omitempty" validate:"omitempty,gt=0"`
}
type InvertedHammerOptions struct {
	BodySizeThreshold float64 `json:"bodySizeThreshold,omitempty" validate:"omitempty,gt=0"`
	TailThreshold     float64 `json:"tailThreshold,omitempty" validate:"omitempty,gt=0"`
}

func HammerDefaultOptions() HammerOptions {
	return HammerOptions{
		BodySizeThreshold: 2,
		WickThreshold:     0.1,
	}
}
func InvertedHammerDefaultOptions() InvertedHammerOptions {
	return InvertedHammerOptions{
		BodySizeThreshold: 2,
		TailThreshold:     0.1,
	}
}

// IsHammer returns true if the candle is a Hammer candle.
// For a Hammer candle to be valid:
// 1. Bullish candle
// 2. Tail (lower wick) at least 2x the body
// 3. Upper wick is not more than 5% of the tail

// IsHammer is a bullish candle and hammer
func IsHammer(candle types.Candle, opt HammerOptions) bool {
	body := BodyLen(candle)
	tail := TailLen(candle)
	wick := WickLen(candle)

	// Conditions for hammer:
	// 1. Bullish candle
	// 2. Tail (lower wick) at least 2x the body
	// 3. Upper wick is not more than 10% of the tail (or as per threshold)

	isBullishCandle := IsBullishCandle(candle)
	hasLongLowerWick := tail >= body*opt.BodySizeThreshold
	upperWickRatio := wick / tail

	isValidWick := upperWickRatio <= opt.WickThreshold

	return isBullishCandle && hasLongLowerWick && isValidWick

}

// Hammer use to find  hammer or hanging man
func Hammer(candle types.Candle, opt HammerOptions) bool {
	body := BodyLen(candle)
	tail := TailLen(candle)
	wick := WickLen(candle)

	// Conditions for hammer:
	// 1. Tail (lower wick) at least 2x the body
	// 2. Upper wick is not more than 10% of the tail (or as per threshold)

	hasLongLowerWick := tail >= body*opt.BodySizeThreshold
	upperWickRatio := wick / tail

	isValidWick := upperWickRatio <= opt.WickThreshold

	return hasLongLowerWick && isValidWick

}

// IsInvertedHammer is a bullish candle and inverted hammer
func IsInvertedHammer(candle types.Candle, opt InvertedHammerOptions) bool {
	body := BodyLen(candle)
	tail := TailLen(candle)
	wick := WickLen(candle)

	// Condition for inverted hammer:
	// 1. Bullish candle
	// 2. Body (upper wick) at least 2x the body
	// 3. Upper tail is not more than 10% of the tail (or per threshold)
	isBullishC := IsBullishCandle(candle)
	hasLongUpperWick := wick >= body*opt.BodySizeThreshold
	lowerTailRatio := tail / wick

	isValidTail := lowerTailRatio <= opt.TailThreshold
	return isBullishC && hasLongUpperWick && isValidTail
}

// InvertedHammer use to find shooting start, inverted hammer
func InvertedHammer(candle types.Candle, opt InvertedHammerOptions) bool {
	body := BodyLen(candle)
	tail := TailLen(candle)
	wick := WickLen(candle)

	// Condition for inverted hammer:
	// 1 Body (upper wick) at least 2x the body
	// 2. Upper tail is not more than 10% of the tail (or per threshold)
	hasLongUpperWick := wick >= body*opt.BodySizeThreshold
	lowerTailRatio := tail / wick

	isValidTail := lowerTailRatio <= opt.TailThreshold
	return hasLongUpperWick && isValidTail
}
