# Chrome Web Store API v2 CLI

Chrome Web Store API v2 を操作するための CLI ツールです。Go ライブラリとしても使用できます。

## CLI インストール

```bash
go install github.com/H0R15H0/chrome-webstore-api-v2/cmd/cws@latest
```

## 認証設定

Chrome Web Store API を使用するには、以下のいずれかの認証方法を設定します：

- **アクセストークン** - CI/CD や自動化向け（推奨）
- **OAuth 2.0** - 個人利用やテスト向け

---

## 方法 1: アクセストークン認証（推奨）

gcloud CLI や他の方法で取得したアクセストークンを直接使用します。CI/CD パイプラインや自動化に最適です。

### 1. Google Cloud Console でプロジェクトを設定

1. [Google Cloud Console](https://console.cloud.google.com/) にアクセス
2. 新しいプロジェクトを作成するか、既存のプロジェクトを選択
3. 「APIs & Services」→「Library」から **Chrome Web Store API** を有効化

### 2. gcloud CLI でアクセストークンを取得

```bash
# gcloud CLI をインストール（未インストールの場合）
# https://cloud.google.com/sdk/docs/install

# 認証（初回のみ）
gcloud auth login

# アクセストークンを取得
gcloud auth print-access-token --scopes=https://www.googleapis.com/auth/chromewebstore
```

### 3. 環境変数を設定

```bash
export CHROME_WEBSTORE_ACCESS_TOKEN="$(gcloud auth print-access-token --scopes=https://www.googleapis.com/auth/chromewebstore)"
export CHROME_WEBSTORE_PUBLISHER_ID="your-publisher-id"
export CHROME_WEBSTORE_ITEM_ID="your-item-id"
```

**注意**: アクセストークンは 1 時間で期限切れになります。CI/CD では毎回新しいトークンを取得してください。

### CI/CD での使用例（GitHub Actions）

```yaml
- name: Setup gcloud
  uses: google-github-actions/setup-gcloud@v2
  with:
    credentials_json: ${{ secrets.GCP_SERVICE_ACCOUNT_KEY }}

- name: Publish to Chrome Web Store
  run: |
    export CHROME_WEBSTORE_ACCESS_TOKEN="$(gcloud auth print-access-token --scopes=https://www.googleapis.com/auth/chromewebstore)"
    cws upload extension.zip
    cws publish
```

---

## 方法 2: OAuth 2.0 認証

OAuth 2.0 は個人利用やテストに適しています。リフレッシュトークンを使用するため、アクセストークンの期限切れを気にする必要がありません。

### 1. OAuth 同意画面を設定

1. 「APIs & Services」→「OAuth consent screen」に移動
2. User Type: **External** を選択して「Create」
3. 必要な情報を入力（アプリ名、サポートメール等）
4. スコープは設定せずに進む
5. テストユーザーに Chrome Web Store のアイテムを所有する Google アカウントを追加

### 2. OAuth 2.0 クレデンシャルを作成

1. 「APIs & Services」→「Credentials」に移動
2. 「Create Credentials」→「OAuth client ID」を選択
3. **Application type: Web application** を選択
4. 「Authorized redirect URIs」に以下を追加:
   ```
   https://developers.google.com/oauthplayground
   ```
5. 「Create」をクリック
6. **Client ID** と **Client Secret** をメモ

### 3. リフレッシュトークンを取得

[OAuth 2.0 Playground](https://developers.google.com/oauthplayground/) を使用してリフレッシュトークンを取得します。

1. 右上の **歯車アイコン** をクリック
2. 「**Use your own OAuth credentials**」にチェック
3. 手順 2 で取得した **Client ID** と **Client Secret** を入力
4. 左側「Step 1」のスコープ入力欄に以下を入力:
   ```
   https://www.googleapis.com/auth/chromewebstore
   ```
5. 「**Authorize APIs**」をクリック
6. Chrome Web Store のアイテムを所有している Google アカウントで認証
7. 「Step 2」で「**Exchange authorization code for tokens**」をクリック
8. 表示された **Refresh token** をコピー

### 4. 環境変数を設定

```bash
export CHROME_WEBSTORE_CLIENT_ID="your-client-id"
export CHROME_WEBSTORE_CLIENT_SECRET="your-client-secret"
export CHROME_WEBSTORE_REFRESH_TOKEN="your-refresh-token"
export CHROME_WEBSTORE_PUBLISHER_ID="your-publisher-id"
export CHROME_WEBSTORE_ITEM_ID="your-item-id"
```

---

## Publisher ID と Item ID の確認方法

- **Publisher ID**: [Chrome Web Store Developer Dashboard](https://chrome.google.com/webstore/devconsole) の URL に含まれる ID
  - 例: `https://chrome.google.com/webstore/devconsole/12345678-abcd-...` の `12345678-abcd-...` 部分
- **Item ID**: 拡張機能の ID（32文字の英小文字）
  - Developer Dashboard で拡張機能を選択した際の URL や、公開 URL に含まれる

## CLI コマンド

| コマンド | 説明 |
|---------|------|
| `cws fetch-status` | アイテムのステータスを取得 |
| `cws upload <file.zip>` | 拡張機能をアップロード |
| `cws publish` | アイテムを公開 |
| `cws cancel-submission` | 保留中の申請をキャンセル |
| `cws set-published-deploy-percentage <percentage>` | デプロイ率を設定 |

## CLI 使用例

```bash
# ステータスを取得
cws fetch-status

# JSON 形式で出力
cws fetch-status --json

# 特定のプロジェクションを指定
cws fetch-status --projection DRAFT

# 拡張機能をアップロード
cws upload extension.zip

# 即時公開
cws publish --type immediate

# 段階的ロールアウト
cws publish --type staged --deploy-percentage 10

# 申請をキャンセル
cws cancel-submission

# デプロイ率を 50% に設定
cws set-published-deploy-percentage 50

# フラグで ID を指定
cws fetch-status --publisher-id my-publisher --item-id my-item
```

---

## Go ライブラリとして使用

CLI だけでなく、Go ライブラリとしてプログラムから直接使用することもできます。

### インストール

```bash
go get github.com/H0R15H0/chrome-webstore-api-v2
```

### クライアントの初期化

#### アクセストークン認証

```go
package main

import (
    "github.com/H0R15H0/chrome-webstore-api-v2/chromewebstore"
)

func main() {
    // アクセストークンからクライアントを作成
    // トークンは gcloud auth print-access-token などで取得
    client := chromewebstore.NewClientFromAccessToken("your-access-token")

    // アイテム名を作成
    itemName := chromewebstore.NewItemName("publisher-id", "item-id")

    // ...
}
```

#### OAuth 2.0 認証

```go
package main

import (
    "context"
    "github.com/H0R15H0/chrome-webstore-api-v2/chromewebstore"
)

func main() {
    ctx := context.Background()

    // OAuth 2.0 認証情報からクライアントを作成
    client := chromewebstore.NewClientFromCredentials(ctx, chromewebstore.AuthConfig{
        ClientID:     "your-client-id",
        ClientSecret: "your-client-secret",
        RefreshToken: "your-refresh-token",
    })

    // アイテム名を作成
    itemName := chromewebstore.NewItemName("publisher-id", "item-id")

    // ...
}
```

### ステータスの取得

```go
status, err := client.Publishers.Items.FetchStatus(itemName).Context(ctx).Do()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Item ID: %s\n", status.ItemID)
if status.PublishedItemRevisionStatus != nil {
    fmt.Printf("State: %s\n", status.PublishedItemRevisionStatus.State)
    for _, ch := range status.PublishedItemRevisionStatus.DistributionChannels {
        fmt.Printf("Version: %s (Deploy: %d%%)\n", ch.CrxVersion, ch.DeployPercentage)
    }
}
```

### 拡張機能のアップロード

```go
file, err := os.Open("extension.zip")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

resp, err := client.Media.Upload(itemName).
    Context(ctx).
    Media(file, "application/zip").
    Do()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Upload state: %s\n", resp.UploadState)
fmt.Printf("Version: %s\n", resp.CrxVersion)
```

### 公開

```go
// 即時公開
resp, err := client.Publishers.Items.Publish(itemName).
    Context(ctx).
    PublishType(chromewebstore.PublishTypeImmediate).
    Do()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("State: %s\n", resp.State)

// 段階的ロールアウト
resp, err := client.Publishers.Items.Publish(itemName).
    Context(ctx).
    PublishType(chromewebstore.PublishTypeStaged).
    DeployPercentage(10).
    Do()
```

### デプロイ率の設定

```go
resp, err := client.Publishers.Items.SetPublishedDeployPercentage(itemName).
    Context(ctx).
    DeployPercentage(50).
    Do()
```

### 申請のキャンセル

```go
resp, err := client.Publishers.Items.CancelSubmission(itemName).Context(ctx).Do()
```

### エラーハンドリング

```go
status, err := client.Publishers.Items.FetchStatus(itemName).Context(ctx).Do()
if err != nil {
    if apiErr, ok := err.(*chromewebstore.APIError); ok {
        if apiErr.IsNotFound() {
            // アイテムが見つからない
        } else if apiErr.IsUnauthorized() {
            // 認証エラー
        } else if apiErr.IsRateLimited() {
            // レート制限
        }
    }
    log.Fatal(err)
}
```

## API リファレンス

### Client

| メソッド | 説明 |
|---------|------|
| `NewClient(httpClient)` | HTTP クライアントから新しいクライアントを作成 |
| `NewClientFromAccessToken(accessToken)` | アクセストークンからクライアントを作成 |
| `NewClientFromCredentials(ctx, config)` | OAuth 2.0 認証情報から新しいクライアントを作成 |

### ItemsService

| メソッド | 説明 |
|---------|------|
| `FetchStatus(name)` | アイテムのステータスを取得 |
| `Publish(name)` | アイテムを公開 |
| `CancelSubmission(name)` | 保留中の申請をキャンセル |
| `SetPublishedDeployPercentage(name)` | デプロイ率を設定 |

### MediaService

| メソッド | 説明 |
|---------|------|
| `Upload(name)` | 拡張機能パッケージをアップロード |

### 型

#### ItemState

| 値 | 説明 |
|---|------|
| `ItemStatePendingReview` | レビュー待ち |
| `ItemStateStaged` | ステージング |
| `ItemStatePublished` | 公開済み |
| `ItemStatePublishedToTesters` | テスターに公開済み |
| `ItemStateRejected` | 却下 |
| `ItemStateCancelled` | キャンセル済み |

#### UploadState

| 値 | 説明 |
|---|------|
| `UploadStateSucceeded` | 成功 |
| `UploadStateInProgress` | 処理中 |
| `UploadStateFailed` | 失敗 |
| `UploadStateNotFound` | 見つからない |

#### PublishType

| 値 | 説明 |
|---|------|
| `PublishTypeImmediate` | 即時公開 |
| `PublishTypeStaged` | 段階的ロールアウト |

## ライセンス

MIT License
