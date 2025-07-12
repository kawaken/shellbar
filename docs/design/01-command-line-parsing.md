# コマンドライン引数パース設計

## 概要

コマンドライン引数パースの実装について議論しました。

### 最終的な設計方針

- ひとまず `flag` パッケージで対応

#### 全体構造

```go
// main.go
func main() {
    shellbar := &Shellbar{}
    if err := shellbar.Run(); err != nil {
        os.Exit(1)
    }
}

// shellbar.go  
type Shellbar struct{}

func (s *Shellbar) Run() error {
    cmd, err := parseArgs()
    if err != nil {
        return err
    }
    return cmd.Run()
}

// command.go
type Command struct {
    run func() error // 内部コマンドの場合は関数、外部の場合はexec実行
}

func (c *Command) Run() error {
    return c.run()
}
```

#### Builder的な関数でコマンド作成

```go
// 外部コマンド
func NewExternalCommand(name string, args []string) *Command {
    return &Command{
        run: func() error {
            cmd := exec.Command(name, args...)
            // PTY制御とかの処理
            return cmd.Run()
        },
    }
}
```
