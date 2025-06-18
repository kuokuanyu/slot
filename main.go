package main

import (
	"fmt"
	"log"
	"slot/slot"
)

func main() {
	board := slot.Spin() // 產生盤面

	// log.Println("測試spin function: ")

	// for _, b := range board {
	// 	log.Println(b)
	// }
	// // log.Println(board)

	log.Println("測試顯示牌面")

	// 顯示盤面（印出來方便查看）
	for row := 0; row < 3; row++ {
		for col := 0; col < 5; col++ {
			fmt.Printf("%4s", board[col][row])
		}
		fmt.Println()
	}

	// win := slot.CalculateWin(board) // 計算分數
	// fmt.Printf("🎉 Win amount: %d\n", win)
}
