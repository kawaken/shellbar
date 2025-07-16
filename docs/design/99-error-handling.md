# エラーハンドリング

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
[YYYY-MM-DD HH:MM:SS] [ERROR] [status_name] error_message
```

### ログ例

```text
[2025-07-11 15:30:45] [ERROR] [git_branch] failed to read .git/HEAD: permission denied
[2025-07-11 15:30:46] [ERROR] [path] failed to get working directory: permission denied
```

