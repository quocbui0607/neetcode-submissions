func isValidSudoku(board [][]byte) bool {
    rowIndx := make(map[byte]map[byte]bool)
    colIndx := make(map[byte]map[byte]bool)
    boxes := make(map[byte]map[byte]bool)

    for r:=byte(0); r < 9; r++ {
        for c:=byte(0); c < 9; c++ {
            if rowIndx[r] == nil {
                rowIndx[r] = map[byte]bool{}
            }

            if colIndx[c] == nil {
                colIndx[c] = map[byte]bool{}
            }

            boxIndx := (r/3)*3 + c/3
            if boxes[boxIndx] == nil {
                boxes[boxIndx] = map[byte]bool{}

            }
            val := board[r][c]
            if val == '.' {
                continue
            }

            if rowIndx[r][val] || colIndx[c][val] || boxes[boxIndx][val] {
                return false
            }

            rowIndx[r][val] =true
            colIndx[c][val] =true
            boxes[boxIndx][val] =true

           
        }
    }

    return true
}
