package chessboard

// File stores if a square is occupied by a piece.
type File []bool

// Chessboard contains a map of eight Files, accessed with keys from "A" to "H."
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var result int
	for _, occupied := range cb[file] {
		if occupied {
			result++
		}
	}
	return result
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank > 8 || rank < 1 {
		return 0
	}

	var result int
	for _, v := range cb {
		if v[rank-1] {
			result++
		}
	}
	return result
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var result int
	for _, v := range cb {
		result += len(v)
	}
	return result
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var result int
	for _, v := range cb {
		for _, occupied := range v {
			if occupied {
				result++
			}
		}
	}
	return result
}
