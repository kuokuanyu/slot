package main

import (
	"fmt"
	"log"
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

	log.Println("翻轉牌面: ")

	// 顯示翻轉盤面（印出來方便查看）
	// example:
	//  A   K   9   K   9
	//  K   Q   J  10Wild
	//  Q   9   K   A   A
	for row := 0; row < 3; row++ {
		for col := 0; col < 5; col++ {
			fmt.Printf("%-10s", board[col][row]) // - 表示左對齊
		}
		fmt.Println()
	}

	fmt.Println("計算中獎分數: ")

	total, details := slot.CalculateWin(board, slot.Paylines, slot.Paytable, "Wild", "Scatter")
	fmt.Printf("總分: %d\n", total)
	for line, win := range details {
		if line == -1 {
			// -1 代表 Scatter 獲勝分數，目前沒有在Paytable設置Scatter分數
			fmt.Println("Scatter 獲勝分數:", win)
		} else if win == 0 {
			fmt.Printf("連線 %d 沒有獲勝分數，但有觸發Bouns\n", line)
		} else {
			fmt.Printf("連線 %d 獲勝分數: %d\n", line+1, win)
		}
	}
}
