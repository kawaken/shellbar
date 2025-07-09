package config

// Config はShellbarの設定を表す構造体
type Config struct {
	Format      string             `toml:"format"`
	RefreshRate string             `toml:"refresh_rate"`
	Defaults    DefaultConfig      `toml:"defaults"`
	Commands    map[string]Command `toml:"commands"`
}

// DefaultConfig はコマンドのデフォルト設定
type DefaultConfig struct {
	Timeout  string `toml:"timeout"`
	Interval string `toml:"interval"`
}

// Command は個別のコマンド設定
type Command struct {
	Command  string `toml:"command"`
	Interval string `toml:"interval"`
	Timeout  string `toml:"timeout"`
}

// NewDefaultConfig はデフォルト設定を返す
func NewDefaultConfig() *Config {
	return &Config{
		Format:      "{path} {git_branch} | {date} {time}",
		RefreshRate: "1s",
		Defaults: DefaultConfig{
			Timeout:  "3s",
			Interval: "5s",
		},
		Commands: make(map[string]Command),
	}
}