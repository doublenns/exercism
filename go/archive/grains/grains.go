package grains

import (
	"errors"
	"fmt"
	"math"
)

const chessboardSquareCount = 64

func Square(number int) (uint64, error) {
	if number <= 0 {
		return uint64(0), errors.New("can't have 0 or less chess board squares")
	} else if number > chessboardSquareCount {
		return uint64(0), fmt.Errorf("there are only %d squares on a chess board", chessboardSquareCount)
	}
	result := math.Pow(2, float64(number-1))
	return uint64(result), nil
}

func Total() uint64 {
	var result uint64
	for i := 1; i <= chessboardSquareCount; i++ {
		counter, _ := Square(i)
		result += counter
	}
	return result
}
