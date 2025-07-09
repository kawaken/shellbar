package config

import (
	"testing"
)

func TestNewDefaultConfig(t *testing.T) {
	t.Run("デフォルト設定が正しく初期化される", func(t *testing.T) {
		cfg := NewDefaultConfig()

		// フォーマットのデフォルト値確認
		expectedFormat := "{path} {git_branch} | {date} {time}"
		if cfg.Format != expectedFormat {
			t.Errorf("Format: got %v, want %v", cfg.Format, expectedFormat)
		}

		// リフレッシュレートのデフォルト値確認
		expectedRefreshRate := "1s"
		if cfg.RefreshRate != expectedRefreshRate {
			t.Errorf("RefreshRate: got %v, want %v", cfg.RefreshRate, expectedRefreshRate)
		}

		// デフォルトコマンド設定の確認
		if cfg.Defaults.Timeout != "3s" {
			t.Errorf("Defaults.Timeout: got %v, want 3s", cfg.Defaults.Timeout)
		}
		if cfg.Defaults.Interval != "5s" {
			t.Errorf("Defaults.Interval: got %v, want 5s", cfg.Defaults.Interval)
		}

		// ビルトインコマンドは設定に含まれないことを確認
		if len(cfg.Commands) != 0 {
			t.Errorf("Commands: got %d commands, want 0 (builtins are not in config)", len(cfg.Commands))
		}
	})
}