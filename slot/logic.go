package slot

// Paylines 定義玩家獲勝「連線」的規則線條。
// 定義哪些符號排列位置構成有效連線（可得分）
// 每一條線由 5 個座標組成 [column, row]
// 高度彈性	可以精確定義任何連線路徑（不侷限直線），適合商業 slot 的複雜設計需求。
// 符合真實機台邏輯	真實 slot 遊戲有很多 zigzag、diagonal、V 形、反 V 形等 Payline，這種 [列, 行] 設計能完全表達。
// 以下是一個較符合商業邏輯的 Paylines 陣列設計，包含：
// 三條水平線（上中下）
// 兩條對角線（左上到右下、左下到右上）
// 多條 Zigzag（Z字形／W字形）斜線
var Paylines = [][][2]int{
    // 1. Top Row
    {{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}}, // 上排直線

    // 2. Middle Row
    {{0, 1}, {1, 1}, {2, 1}, {3, 1}, {4, 1}}, // 中排直線

    // 3. Bottom Row
    {{0, 2}, {1, 2}, {2, 2}, {3, 2}, {4, 2}}, // 下排直線

    // 4. Diagonal Left Top → Right Bottom
    {{0, 0}, {1, 1}, {2, 2}, {3, 1}, {4, 0}}, // V字型對角

    // 5. Diagonal Left Bottom → Right Top
    {{0, 2}, {1, 1}, {2, 0}, {3, 1}, {4, 2}}, // 倒V字型對角

    // 6. Zigzag Top → Mid → Top → Mid → Top
    {{0, 0}, {1, 1}, {2, 0}, {3, 1}, {4, 0}}, // Z字型

    // 7. Zigzag Bottom → Mid → Bottom → Mid → Bottom
    {{0, 2}, {1, 1}, {2, 2}, {3, 1}, {4, 2}}, // 倒Z字型

    // 8. W Shape
    {{0, 0}, {1, 2}, {2, 1}, {3, 2}, {4, 0}}, // W字型

    // 9. M Shape
    {{0, 2}, {1, 0}, {2, 1}, {3, 0}, {4, 2}}, // M字型

    // 10. Outward Arrow
    {{0, 1}, {1, 0}, {2, 1}, {3, 2}, {4, 1}}, // 箭頭展開形

    // 11. Inward Arrow
    {{0, 1}, {1, 2}, {2, 1}, {3, 0}, {4, 1}}, // 箭頭收束形
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
