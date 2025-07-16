# ネーミング設計

## 概要

ステータスバーの構成要素を明確に分離し、ユーザーにとって直感的なネーミング方針を定める。

## 基本方針

### 用語定義

- **Status**: 設定・情報源を表す（ステータスバーに表示する情報の設定）
  - 何を表示するか
  - いつ更新するか
  - どのくらいの間隔で実行するか
- **Command**: 実行処理を表す（実際にその情報を取得する処理）
  - 実行時の動作
  - 具体的な処理内容

### ネーミング規則

- **設定ファイル**: `[status.名前]` セクションでステータス設定を定義
- **コードベース**: `Command` インターフェースで実行処理を実装
- **判定ロジック**: `command`フィールドの有無でビルトイン/外部を判定

### 具体例

#### 設定ファイル例（TOML）

```toml
# ビルトインコマンド（commandフィールドなし）
[status.git_branch]
interval = "3s"
timeout = "3s"

# 外部コマンド（commandフィールドあり）
[status.cpu_usage]
command = "top -bn1 | grep 'Cpu(s)' | awk '{print $2}'"
interval = "5s"
timeout = "2s"
```

#### コード例（Go）

```go
// Commandインターフェース
type Command interface {
    Run(ctx context.Context) error
}

// ビルトインコマンドの実装例
type GitBranchCommand struct{}

func (g *GitBranchCommand) Run(ctx context.Context) error {
    // git branch情報を取得する処理
    return nil
}

// 設定からCommandを生成する例
func createCommand(name string, config StatusConfig) Command {
    if config.Command == "" {
        // ビルトインコマンドの場合
        switch name {
        case "git_branch":
            return &GitBranchCommand{}
        // ...
        }
    }
    // 外部コマンドの場合
    return NewExternalCommand(config.Command)
}
```
