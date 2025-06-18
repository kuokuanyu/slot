package slot

// Paytable 定義了各個符號的中獎規則
// 格式為 [符號][連線數] = 中獎分數
var Paytable = map[string]map[int]int{
	"A":    {3: 10, 4: 20, 5: 50},
	"K":    {3: 5, 4: 10, 5: 30},
	"Q":    {3: 5, 4: 10, 5: 25},
	"Wild": {3: 15, 4: 50, 5: 100}, // Wild 可替代其他符號
}
