package slot

import "testing"

// TestCalculateWin 測試 CalculateWin 函式
// 測試不同情況下的中獎計算，沒有 Wild 和 Scatter
func TestCalculateWin(t *testing.T) {
	board := [][]string{
		{"A", "K", "Q"},
		{"A", "K", "Q"},
		{"A", "K", "Q"},
		{"A", "K", "Q"},
		{"A", "K", "Q"},
	}

	paylines := [][][2]int{
		{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}}, // Top row
	}

	paytable := map[string]map[int]int{
		"A": {3: 10, 4: 20, 5: 50},
	}

	expected := 50
	score, _ := CalculateWin(board, paylines, paytable, "Wild", "Scatter")
	if score != expected {
		t.Errorf("expected %d, got %d", expected, score)
	}
}

// TestCalculateWinWithScatter 測試 CalculateWin 函式
// 測試有 Wild 符號的情況
func TestCalculateWinWithWild(t *testing.T) {
	board := [][]string{
		{"Wild", "K", "Q"},
		{"A", "K", "Q"},
		{"Wild", "K", "Q"},
		{"A", "K", "Q"},
		{"Wild", "K", "Q"},
	}

	paylines := [][][2]int{
		{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}}, // Top row
	}

	paytable := map[string]map[int]int{
		"A": {3: 10, 4: 20, 5: 50},
	}

	expected := 50 // Wild替代了A，形成5連線
	score, _ := CalculateWin(board, paylines, paytable, "Wild", "Scatter")
	if score != expected {
		t.Errorf("expected %d, got %d", expected, score)
	}
}

// TestCalculateWin_MultipleCases 測試多個 CalculateWin 函式
func TestCalculateWin_MultipleCases(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]string
		expected int
	}{
		{
			name: "A連線",
			board: [][]string{
				{"A", "K", "Q"},
				{"A", "K", "Q"},
				{"A", "K", "Q"},
				{"A", "K", "Q"},
				{"A", "K", "Q"},
			},
			expected: 50,
		},
		{
			name: "沒中獎",
			board: [][]string{
				{"A", "K", "Q"},
				{"K", "A", "Q"},
				{"Q", "A", "K"},
				{"J", "K", "Q"},
				{"10", "9", "8"},
			},
			expected: 0,
		},
	}

	paylines := [][][2]int{
		{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}}, // Top row
	}
	paytable := map[string]map[int]int{
		"A": {3: 10, 4: 20, 5: 50},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score, _ := CalculateWin(tt.board, paylines, paytable, "Wild", "Scatter")
			if score != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, score)
			}
		})
	}
}
