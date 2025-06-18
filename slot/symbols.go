package slot

// Reels 為每個轉輪的符號組成。
// 假設為 5 個轉輪，每個轉輪內可能出現的符號都一樣。
// var Reels = [][]string{
//     {"A", "K", "Q", "J", "10", "9", "Wild", "Scatter"},
//     {"A", "K", "Q", "J", "10", "9", "Wild", "Scatter"},
//     {"A", "K", "Q", "J", "10", "9", "Wild", "Scatter"},
//     {"A", "K", "Q", "J", "10", "9", "Wild", "Scatter"},
//     {"A", "K", "Q", "J", "10", "9", "Wild", "Scatter"},
// }

// 更貼近實際場景
var Reels = [][]string{
    // Reel 1
    {"A", "10", "K", "Q", "J", "10", "9", "Wild", "A", "K", "Q", "J", "10", "Scatter", "A", "10", "Q", "9", "Wild", "K"},
    
    // Reel 2
    {"Q", "A", "Wild", "9", "K", "J", "Scatter", "10", "J", "Q", "Wild", "K", "A", "10", "9", "Q", "Scatter", "K", "J"},
    
    // Reel 3
    {"10", "9", "K", "Q", "A", "Scatter", "Wild", "J", "A", "10", "Q", "Wild", "9", "J", "K", "Q", "Scatter", "A"},
    
    // Reel 4
    {"A", "K", "Wild", "Scatter", "10", "9", "Q", "J", "K", "10", "A", "Wild", "9", "J", "Q", "Scatter", "K", "J", "10"},
    
    // Reel 5
    {"Wild", "A", "Q", "K", "J", "Scatter", "10", "9", "Wild", "Q", "A", "J", "K", "Scatter", "9", "10", "K", "Wild"},
}
