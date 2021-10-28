package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number <= 0 {
		return 0, errors.New("can't have 0 or less chess board squares")
	} else if number > 64 {
		return 0, errors.New("there are only 64 squares on a chess board")
	}
	result := math.Pow(2, float64(number-1))
	return uint64(result), nil
}

func Total() uint64 {
	var result uint64
	const chessboardSquareCount = 64
	for i := 1; i <= chessboardSquareCount; i++ {
		counter, _ := Square(i)
		result += counter
	}
	return result
}
