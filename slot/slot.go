package slot

import (
	"fmt"
	"math/rand"
)

// 定義 slot 的符號
var symbols = []string{"🍒", "🍋", "🔔", "⭐", "🍇", "7️⃣"}

// 建立 3x3 的轉輪盤
func Spin() [][]string {
	board := make([][]string, 3)
	for i := 0; i < 3; i++ {
		board[i] = make([]string, 3)
		for j := 0; j < 3; j++ {
			board[i][j] = symbols[rand.Intn(len(symbols))]
		}
	}
	return board
}

// 印出 slot 結果
func PrintBoard(board [][]string) {
	for _, row := range board {
		for _, val := range row {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
}

// 判斷是否中 Jackpot（中間橫列三個一樣就算中獎）
func IsJackpot(board [][]string) bool {
	mid := board[1]
	return mid[0] == mid[1] && mid[1] == mid[2]
}
