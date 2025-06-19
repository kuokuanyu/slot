package slot

// 不同欄位長度（增加隨機性與控制中獎率）
// 符號權重不同（高價符號少見）
// 仍保留「連續三格抽出」的邏輯（真實 slot 大多使用）
// 通常為 30～100 個 symbol。
// 有些大型機種甚至會到 200 個 symbol，因為這樣可以精細控制中獎機率（Hit Rate）與賠率分佈。
// 更貼近實際場景
var Reels = [][]string{
	// Reel 1 - 52 個符號
	{"A", "10", "K", "Q", "J", "10", "9", "Wild", "A", "K", "Q", "J", "10", "Scatter",
	 "A", "10", "Q", "9", "Wild", "K", "J", "9", "Q", "A", "K", "Wild",
	 "10", "9", "J", "Scatter", "Q", "Wild", "K", "A", "10", "Q", "J", "K",
	 "10", "9", "Wild", "Scatter", "A", "K", "Q", "J", "10", "9", "K", "A", "J", "Q"},

	// Reel 2 - 60 個符號
	{"Q", "A", "Wild", "9", "K", "J", "Scatter", "10", "J", "Q", "Wild", "K", "A", "10",
	 "9", "Q", "Scatter", "K", "J", "Q", "10", "9", "Wild", "K", "A", "J", "Scatter",
	 "Q", "K", "J", "10", "A", "Wild", "9", "K", "Q", "J", "A", "10", "9", "K",
	 "Wild", "Q", "J", "Scatter", "A", "K", "10", "9", "Wild", "Q", "A", "J", "10", "K", "Q", "9", "J", "Scatter"},

	// Reel 3 - 48 個符號（中間轉輪較短）
	{"10", "9", "K", "Q", "A", "Scatter", "Wild", "J", "A", "10", "Q", "Wild",
	 "9", "J", "K", "Q", "Scatter", "A", "Wild", "K", "J", "Q", "10", "9",
	 "Scatter", "K", "A", "J", "Q", "10", "Wild", "9", "K", "J", "Q", "A",
	 "10", "9", "Wild", "K", "A", "Q", "J", "10", "9", "Wild", "Scatter", "Q"},

	// Reel 4 - 56 個符號
	{"A", "K", "Wild", "Scatter", "10", "9", "Q", "J", "K", "10", "A", "Wild",
	 "9", "J", "Q", "Scatter", "K", "J", "10", "A", "Q", "Wild", "9", "K",
	 "J", "Q", "10", "Scatter", "A", "K", "Wild", "9", "10", "Q", "J", "A",
	 "K", "10", "Wild", "Scatter", "J", "Q", "A", "10", "9", "K", "Q", "Wild",
	 "10", "J", "Q", "A", "K", "Wild", "9", "Scatter"},

	// Reel 5 - 63 個符號（最長）
	{"Wild", "A", "Q", "K", "J", "Scatter", "10", "9", "Wild", "Q", "A", "J", "K", "Scatter",
	 "9", "10", "K", "Wild", "J", "Q", "10", "9", "A", "K", "Wild", "Q", "J", "10",
	 "Scatter", "A", "K", "9", "Q", "J", "Wild", "10", "A", "Q", "K", "Scatter", "J", "9",
	 "Wild", "A", "K", "10", "Q", "J", "Scatter", "10", "9", "K", "A", "Wild", "Q", "J",
	 "10", "Scatter", "A", "K", "Q", "J", "Wild"},
}

// Reels 為每個轉輪的符號組成。
// 假設為 5 個轉輪，每個轉輪內可能出現的符號都一樣。
// var Reels = [][]string{
//     {"A", "K", "Q", "J", "10", "9", "Wild", "Scatter"},
//     {"A", "K", "Q", "J", "10", "9", "Wild", "Scatter"},
//     {"A", "K", "Q", "J", "10", "9", "Wild", "Scatter"},
//     {"A", "K", "Q", "J", "10", "9", "Wild", "Scatter"},
//     {"A", "K", "Q", "J", "10", "9", "Wild", "Scatter"},
// }
