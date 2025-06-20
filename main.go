package main

import (
	"fmt"
	"slot/slot"
)

func main() {
	// 產生盤面
	// 這裡的盤面是 5 個轉輪，每個轉輪有 3 行
	// example:
	// [   9 Wild Scatter]
	// [   K    A   10]
	// [   Scatter    Q   10]
	// [   K   10    A]
	// [   A    K Wild]
	board := slot.Spin()

	// log.Println("測試spin function: ")

	// for _, b := range board {
	// 	fmt.Printf("%4s", b)
	// }
	// fmt.Println()

	// log.Println("測試顯示牌面: ")

	// 顯示盤面（印出來方便查看）
	// example:
	//  A   K   9   K   9
	//  K   Q   J  10Wild
	//  Q   9   K   A   A
	for row := 0; row < 3; row++ {
		for col := 0; col < 5; col++ {
			fmt.Printf("%4s", board[col][row])
		}
		fmt.Println()
	}

	// win := slot.CalculateWin(board) // 計算分數
	// fmt.Printf("🎉 Win amount: %d\n", win)
}
