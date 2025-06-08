package tests

import (
	"testing"
)

func TestMain(t *testing.T) {
	// 基本的なテストケース
	t.Run("Basic Test", func(t *testing.T) {
		// ここにテストケースを追加
		expected := true
		actual := true
		if expected != actual {
			t.Errorf("Expected %v but got %v", expected, actual)
		}
	})
} 