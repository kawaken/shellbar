# 設定ファイル

## 設定構造の方針

### 基本構造

- **フォーマット**: TOML形式
- **階層**: フラットな構造を基本とし、必要に応じてセクションを使用
- **時間間隔の扱い**: 文字列で指定する。Goのtime.Durationのフォーマットに従う([time.ParseDuration](https://pkg.go.dev/time#ParseDuration))

### 設定項目

- **format**: ステータスバーの表示フォーマット（変数展開対応）
- **refresh_rate**: 画面更新の最小間隔
- **defaults**: 各ステータスのデフォルト値（timeout、interval）
- **status**: ステータス情報の定義


### 設定の適用ルール

- **設定ファイルなし**: ビルトインコマンド + デフォルトフォーマット
- **設定ファイルあり**: ビルトインコマンド + ユーザー定義ステータス + ユーザーフォーマット
- **個別設定が最優先**: statusセクションで個別に設定された値を使用
- **defaults次優先**: 個別設定がない場合はdefaultsセクションの値を使用

### デフォルト値

- **format**: `"{path} {git_branch} | {date} {time}"`
- **refresh_rate**: `"1s"` - 画面更新の最小間隔
- **defaults.interval**: `"5s"` - コマンド実行間隔のデフォルト
- **defaults.timeout**: `"3s"` - コマンドタイムアウトのデフォルト

## バリデーション

### 設定読み込み時にチェックする項目

1. **時間形式の妥当性**
   - 対象: refresh_rate、defaults.timeout、defaults.interval、各status.interval、各status.timeout
   - 手法: time.ParseDurationでの検証
   - 許可単位: `ns`、`us`（`µs`）、`ms`、`s`、`m`、`h`（time.ParseDurationでサポートされる単位）
   - 対処: 
     - ビルトインコマンドの場合、特にエラーにせず、デフォルト値を採用し実行する
     - 外部コマンドの場合、実行せずプレースホルダを表示する

2. **フォーマット文字列の整合性**
   - 対象: formatで参照している変数名（{variable_name}形式）
   - 例: `{git_branch}`を使うなら`status.git_branch`が必要
   - 対処: 存在しないstatusが指定されている場合には、プレースホルダをそのまま表示する

3. **必須フィールドの存在**
   - 対象: 外部コマンドを使用するstatusのcommandフィールド
   - 対処: commandが指定されていないstatusは、プレースホルダをそのまま表示する

### コマンド実行時にチェックする項目

1. **コマンドの実行可能性**
   - 対処: プレースホルダを表示

## 設定ファイルの優先順位

1. `./.shellbar.toml` （カレントディレクトリ）
2. `~/.shellbar.toml` （ホームディレクトリ直下）
3. `~/.config/shellbar/shellbar.toml` （XDG準拠）
4. デフォルト設定での動作


## 設定ファイル例

### 基本的な設定例

```toml
format = "{path} {git_branch} {git_status} | {time}"
refresh_rate = "500ms"

[defaults]
interval = "5s"
timeout = "3s"

# ビルトインコマンドを使用するステータス
[status.path]
interval = "5s"
timeout = "3s"

[status.git_branch]
interval = "3s"
timeout = "3s"

[status.time]
interval = "1s"
timeout = "1s"

# 外部コマンドを使用するステータス
[status.git_status]
command = "git status --porcelain | wc -l"
interval = "2s"
timeout = "5s"
```
