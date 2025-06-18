package slot

// Paylines 定義了所有支付線的位置
// 每一條線由 5 個座標組成 [column, row]
var Paylines = [][][2]int{
	{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}}, // 第一行
	{{0, 1}, {1, 1}, {2, 1}, {3, 1}, {4, 1}}, // 第二行（中間）
	{{0, 2}, {1, 2}, {2, 2}, {3, 2}, {4, 2}}, // 第三行
}

// CalculateWin 根據盤面計算中獎金額
func CalculateWin(board [][]string) int {
	totalWin := 0

	for _, line := range Paylines {
		symbol := board[line[0][0]][line[0][1]] // 第一個符號作為參考
		match := 1

		for i := 1; i < len(line); i++ {
			next := board[line[i][0]][line[i][1]] // 取得下一個位置的符號
			if next == symbol || next == "Wild" {
				match++ // 符號相同（或是 Wild）就算連線
			} else {
				break // 不同就中斷
			}
		}

		// 查找 Paytable 是否符合連線條件，加入分數
		if win, ok := Paytable[symbol][match]; ok {
			totalWin += win
		}
	}

	return totalWin
}
