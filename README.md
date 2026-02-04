# Chrome Web Store API v2 Go Client

Chrome Web Store API v2 の Go クライアントライブラリです。

## インストール

```bash
go get github.com/H0R15H0/chrome-webstore-api-v2
```

## 認証設定

Chrome Web Store API を使用するには、OAuth 2.0 クレデンシャルが必要です。

### 1. Google Cloud Console でプロジェクトを作成

1. [Google Cloud Console](https://console.cloud.google.com/) にアクセス
2. 新しいプロジェクトを作成するか、既存のプロジェクトを選択
3. Chrome Web Store API を有効化

### 2. OAuth 2.0 クレデンシャルを取得

1. APIs & Services > Credentials に移動
2. Create Credentials > OAuth client ID を選択
3. Application type: Desktop app を選択
4. Client ID と Client Secret をメモ

### 3. リフレッシュトークンを取得

OAuth 2.0 Playground や独自の認証フローを使用してリフレッシュトークンを取得します。

必要なスコープ:
- `https://www.googleapis.com/auth/chromewebstore` - 読み書きアクセス
- `https://www.googleapis.com/auth/chromewebstore.readonly` - 読み取り専用

## 使用例

### クライアントの初期化

```go
package main

import (
    "context"
    "github.com/H0R15H0/chrome-webstore-api-v2/chromewebstore"
)

func main() {
    ctx := context.Background()

    // 認証情報からクライアントを作成
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

fmt.Printf("State: %s\n", status.State)
fmt.Printf("Version: %s\n", status.Version)
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

fmt.Printf("Upload status: %s\n", resp.StatusCode)
```

### 公開

```go
// 全ユーザーに公開
resp, err := client.Publishers.Items.Publish(itemName).Context(ctx).Do()
if err != nil {
    log.Fatal(err)
}

// 信頼できるテスターのみに公開
resp, err := client.Publishers.Items.Publish(itemName).
    Context(ctx).
    PublishTarget(chromewebstore.PublishTargetTrustedTesters).
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

## API リファレンス

### Client

| メソッド | 説明 |
|---------|------|
| `NewClient(httpClient)` | HTTP クライアントから新しいクライアントを作成 |
| `NewClientFromCredentials(ctx, config)` | 認証情報から新しいクライアントを作成 |

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
| `ItemStateDraft` | 下書き |
| `ItemStatePendingReview` | レビュー待ち |
| `ItemStatePublished` | 公開済み |
| `ItemStateRejected` | 却下 |
| `ItemStateTakenDown` | 削除済み |
| `ItemStateInReview` | レビュー中 |
| `ItemStatePendingPublish` | 公開待ち |

#### PublishTarget

| 値 | 説明 |
|---|------|
| `PublishTargetDefault` | 全ユーザー |
| `PublishTargetTrustedTesters` | 信頼できるテスターのみ |

## エラーハンドリング

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

## ライセンス

MIT License
