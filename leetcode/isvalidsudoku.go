package main

func isValidSudoku(board [][]byte) bool {

	bm := make(map[string]bool)

	for row:=0; row<9; row ++ {
		for col:=0;col<9; col++ {

			cur :=string(board[row][col])
			if cur =="." { continue}
			_,ok1 := bm[cur + "row" + string(row)]
			_,ok2 := bm[cur + "column"+  string(col)]
			_,ok3 := bm[cur + "box" + string(row/3) + "-" + string(col/3)]

			if ok1 ||ok2||ok3{

				return false
			} else {
				bm[cur + "row" + string(row)] = true
				bm[cur + "column"+  string(col)] = true
				bm[cur + "box" + string(row/3) + "-" + string(col/3)]= true
			}
		}

	}
	return true
}