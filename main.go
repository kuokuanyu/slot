package main

import (
	"fmt"
	"log"
	"os"
	"slot/slot"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 載入 .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	// log.Println("印出env中的PORT:", port)
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	router.GET("/lucky", func(c *gin.Context) {
		board := slot.Spin()
		// log.Println("測試spin function: ")

		total, details := slot.CalculateWin(board, slot.Paylines, slot.Paytable, "Wild", "Scatter")

		// 將盤面轉換為可讀格式（row-major 3x5）
		displayBoard := make([][]string, 3) // 3 row
		for row := 0; row < 3; row++ {
			displayBoard[row] = make([]string, 5)
			for col := 0; col < 5; col++ {
				displayBoard[row][col] = board[col][row]
			}
		}

		// 整理連線結果
		lineResults := []gin.H{}
		for line, win := range details {
			if line == -1 {
				// -1 代表 Scatter 獲勝分數，目前沒有在Paytable設置Scatter分數
				// lineResults = append(lineResults, gin.H{
				// 	"type":   "scatter",
				// 	"score":  win,
				// 	"detail": "Scatter 獲勝",
				// })
			} else if win == 0 {
				lineResults = append(lineResults, gin.H{
					"type":   "bonus",
					"line":   line + 1,
					"score":  0,
					"detail": "Bonus 觸發但無獲勝分數",
				})
			} else {
				lineResults = append(lineResults, gin.H{
					"type":   "line",
					"line":   line + 1,
					"score":  win,
					"detail": fmt.Sprintf("連線 %d 獲勝 %d 分", line+1, win),
				})
			}
		}

		// 翻轉陣列
		formattedBoard := FormatBoardText(board)

		c.JSON(200, gin.H{
			"board":     formattedBoard,
			"total":     total,
			"win_lines": lineResults,
		})
	})

	fmt.Printf("🎰 Slot API running on port %s\n", port)
	router.Run(":" + port)

}

// FormatBoardText 將 board（欄為主的資料）轉成 row 為主的文字格式，方便顯示盤面
func FormatBoardText(board [][]string) string {
	var result string
	for row := 0; row < 3; row++ {
		for col := 0; col < 5; col++ {
			symbol := board[col][row]
			result += fmt.Sprintf("%5s", symbol)
		}
		result += "\n"
	}

	// fmt.Println("翻轉牌面: ")
	// fmt.Println(result)
	return result
}

// 這是測試用的程式碼片段，主要用來測試 slot package 的功能
// 產生盤面
// 這裡的盤面是 5 個轉輪，每個轉輪有 3 行
// example:
// [   9 Wild Scatter]
// [   K    A   10]
// [   Scatter    Q   10]
// [   K   10    A]
// [   A    K Wild]
// board := slot.Spin()

// log.Println("測試spin function: ")

// for _, b := range board {
// 	fmt.Printf("%4s", b)
// }
// fmt.Println()

// log.Println("翻轉牌面: ")

// 顯示翻轉盤面（印出來方便查看）
// example:
//  A   K   9   K   9
//  K   Q   J  10Wild
//  Q   9   K   A   A
// for row := 0; row < 3; row++ {
// 	for col := 0; col < 5; col++ {
// 		fmt.Printf("%-10s", board[col][row]) // - 表示左對齊
// 	}
// 	fmt.Println()
// }

// total, details := slot.CalculateWin(board, slot.Paylines, slot.Paytable, "Wild", "Scatter")
// fmt.Printf("總分: %d\n", total)
// for line, win := range details {
// 	if line == -1 {
// 		// -1 代表 Scatter 獲勝分數，目前沒有在Paytable設置Scatter分數
// 		fmt.Println("Scatter 獲勝分數:", win)
// 	} else if win == 0 {
// 		fmt.Printf("連線 %d 沒有獲勝分數，但有觸發Bouns\n", line)
// 	} else {
// 		fmt.Printf("連線 %d 獲勝分數: %d\n", line+1, win)
// 	}
// }
