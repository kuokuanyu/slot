package main

import (
	"fmt"
	"math/rand"
	"slot/slot"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 每次執行都隨機

	board := slot.Spin() // 產生轉輪結果

	fmt.Println("🎰 Slot 結果：")
	slot.PrintBoard(board)

	if slot.IsJackpot(board) {
		fmt.Println("🎉 恭喜你中 JACKPOT！")
	} else {
		fmt.Println("😢 很可惜，沒有中獎")
	}
}
