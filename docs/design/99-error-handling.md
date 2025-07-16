# エラーハンドリング設計

## 基本方針

ステータスバーとしての役割を重視した統一的なエラーハンドリング

### エラー分類と対応

| 状況             | ステータスバー表示      | ログ出力       |
| ---------------- | ----------------------- | -------------- |
| **正常取得**     | 結果を表示              | なし           |
| **データなし**   | 空文字を表示            | なし           |
| **実行時エラー** | `{variable_name}`を表示 | 詳細エラー情報 |

### 具体的な動作例

#### Git関連コマンド（git_branch）

```bash
# 正常（gitリポジトリ内）
~/projects/myapp main | 15:30:45

# データなし（gitリポジトリ外）
~/projects/myapp | 15:30:45

# 実行時エラー（.git破損、権限エラーなど）
~/projects/myapp {git_branch} | 15:30:45
```

#### Path関連コマンド（path）

```bash
# 正常
~/projects/myapp main | 15:30:45

# 実行時エラー（権限エラーなど）
{path} main | 15:30:45
```

#### 時刻関連コマンド（date, time）

```bash
# 正常
~/projects/myapp main | 15:30:45

# 実行時エラー（システムエラー）
~/projects/myapp main | {time}
```

## エラーログ仕様

### ログファイル配置

XDG Base Directory仕様に従い、以下の優先順位でログディレクトリを決定：

1. `$XDG_DATA_HOME/shellbar/` （通常は `~/.local/share/shellbar/`）
2. `$HOME/.local/share/shellbar/` （XDG_DATA_HOMEが未設定の場合）

### ログファイル名

- ファイル名：`shellbar.log`
- ローテーション：サイズまたは日付ベース（将来実装）

### ログフォーマット

```text
[YYYY-MM-DD HH:MM:SS] [ERROR] [command_name] error_message
```

### ログ例

```text
[2025-07-11 15:30:45] [ERROR] [git_branch] failed to read .git/HEAD: permission denied
[2025-07-11 15:30:46] [ERROR] [path] failed to get working directory: permission denied
```

## 設定関連のエラーハンドリング戦略

### 設定読み込みエラー

- **ファイル不存在**: デフォルト設定で動作開始
- **TOML解析エラー**: エラーメッセージ表示後、デフォルト設定で継続
- **バリデーションエラー**: 具体的なエラーメッセージで停止

### バリデーションエラーの種類

- **時間形式エラー**: 正しいフォーマット例を提示
- **フォーマット整合性エラー**: 未定義の変数名を具体的に指摘
- **必須フィールドエラー**: どのステータスのcommandが未定義かを明示

### 実行時エラー

- **コマンド失敗**: エラー内容をキャッシュして表示継続
- **タイムアウト**: 前回の結果を使用、エラー状態も表示
