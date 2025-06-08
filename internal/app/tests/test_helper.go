package tests

import (
	"os"
	"testing"
)

// SetupTestDB はテスト用のデータベース接続を設定します
func SetupTestDB(t *testing.T) {
	// テスト用の環境変数を設定
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_NAME", "hensai-kun_test")
}

// TeardownTestDB はテスト用のデータベース接続をクリーンアップします
func TeardownTestDB(t *testing.T) {
	// クリーンアップ処理を実装
}

// TestMain はテストのセットアップとクリーンアップを行います
func TestMain(m *testing.M) {
	// テストの前処理
	// テストの実行
	code := m.Run()
	// テストの後処理
	os.Exit(code)
} 