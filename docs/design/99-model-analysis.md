# モデル分析：設計書とコード実装の比較

## 概要

shellbarプロジェクトの設計書とコード実装を比較し、モデル構造の現状と相違点を分析しました。

## 現在のモデル構造

### 1. 実装済みのモデル

#### **Shellbar** (`shellbar.go`)

```go
type Shellbar struct{}
```

- **役割**: メインエントリーポイント
- **実装状況**: ✅ 完了
- **設計書との一致**: ✅ 一致

#### **Command** (`command.go`)

```go
type Command struct {
    run func() error
}
```

- **役割**: コマンド実行の抽象化
- **実装状況**: ✅ 完了（ビルトインコマンドのみ）
- **設計書との一致**: ✅ 一致

#### **Config** (`config/config.go`)

```go
type Config struct {
    Format      string             `toml:"format"`
    RefreshRate string             `toml:"refresh_rate"`
    Defaults    DefaultConfig      `toml:"defaults"`
    Commands    map[string]Command `toml:"commands"`
}
```

- **役割**: メイン設定構造体
- **実装状況**: ✅ 完了
- **設計書との相違**: ⚠️ 部分的相違（後述）

## 主要な相違点と矛盾

### 1. 設定構造の相違

#### **Commands vs Status**

- **設計書**: `[status.名前]` セクションでステータス設定を定義
- **実装**: `Commands map[string]Command` でコマンド設定を管理

```toml
# 設計書での想定
[status.git_branch]
interval = "3s"
timeout = "3s"

[status.git_status]
command = "git status --porcelain | wc -l"
interval = "2s"
timeout = "5s"
```

```go
// 実装での構造
Commands map[string]Command `toml:"commands"`
```

**影響**: 設定ファイルの構造が異なるため、設計書通りの設定ファイルが読み込めない

### 2. ビルトインコマンドの取り扱い

#### **判定ロジックの相違**

- **設計書**: `commandフィールドの有無`で判定
- **実装**: まだ判定ロジックが未実装

#### **設定方法の相違**

- **設計書**: ビルトインコマンドは設定不要、フォーマットで直接指定
- **実装**: 設定構造はあるが、ビルトインコマンドの自動判定未実装

### 3. コマンド実行機能の実装状況

#### **外部コマンド実行**

- **設計書**: 外部コマンドの実行機能を想定
- **実装**: `NewExternalCommand`は placeholder（未実装）

```go
// 実装（未完了）
func NewExternalCommand(name string, args []string) *Command {
    return &Command{
        run: func() error {
            fmt.Printf("sbar: executing %s %v (not implemented yet)\n", name, args)
            return nil
        },
    }
}
```

#### **ビルトインコマンド**

- **設計書**: path, date, time, git_branch の4つを定義
- **実装**: version, help のみ実装済み

### 4. エラーハンドリングの実装状況

- **設計書**: 詳細なエラーハンドリング戦略を定義
- **実装**: 基本的なエラーハンドリングのみ

## 未実装・不整合の項目

### 1. 設定ファイル読み込み機能

- **状況**: 設定構造体は定義済みだが、ファイル読み込み機能が未実装
- **影響**: TOMLファイルからの設定読み込みができない

### 2. ステータスバー表示機能

- **状況**: フォーマット文字列展開機能が未実装
- **影響**: `{path} {git_branch} | {date} {time}` の形式が処理できない

### 3. ビルトインコマンドの実装

- **状況**: path, date, time, git_branch の実装が未完了
- **影響**: デフォルト設定での動作ができない

### 4. 設定ファイル検索機能

- **設計書**: XDG準拠のパス検索を定義
- **実装**: 未実装

### 5. エラーログ機能

- **設計書**: XDG準拠のログディレクトリを定義
- **実装**: 未実装

## 優先対応項目

### Phase 1: 基本機能実装

1. **設定ファイル読み込み機能**: `config` パッケージに LoadConfig 関数追加
2. **ビルトインコマンド実装**: path, date, time, git_branch の実装
3. **ステータスバー表示機能**: フォーマット文字列展開機能
4. **設定構造の統一**: `Commands` を `Status` に変更、または設計書を実装に合わせる

### Phase 2: 外部コマンド対応

1. **外部コマンド実行機能**: `NewExternalCommand` の実装
2. **ビルトイン/外部判定ロジック**: command フィールド有無での判定
3. **エラーハンドリング**: 設計書通りのエラー処理実装

### Phase 3: 運用機能

1. **設定ファイル検索**: XDG準拠のパス検索
2. **エラーログ**: ログディレクトリとファイル出力
3. **バリデーション**: 設定値の検証機能

## 推奨対応方針

### 1. 設計書とコードの統一

設定構造について、以下のいずれかを選択：

**オプションA**: コードを設計書に合わせる

- `Commands` → `Status` に変更
- `[status.名前]` セクション対応

**オプションB**: 設計書をコードに合わせる

- `[commands.名前]` セクションで統一
- 現在の実装を前提とした設計書更新

### 2. 段階的実装

現在の基本構造は良好なため、Phase 1から順次実装を進める

### 3. テスト駆動開発

設計書の仕様に基づいたテストを先に作成し、実装を進める
