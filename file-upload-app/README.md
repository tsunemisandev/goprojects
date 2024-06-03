# ファイルアップロードアプリ

## 説明
これは、Go、Gin、およびBootstrapを使用して構築されたシンプルなファイルアップロードアプリケーションです。ユーザーはファイルをアップロードし、それらをサムネイルモードまたはリストモードで表示することができます。さらに、ホームページにはローカルIPアドレスを表示するQRコードが表示されます。

## 特徴
- 複数ファイルのアップロード
- アップロードされたファイルをサムネイルまたはリスト（ファイルサイズ付き）で表示
- ローカルIPアドレス付きのQRコード表示
- BootstrapによるモバイルフレンドリーなUI

## セットアップ
1. リポジトリをクローンします：
    ```bash
    git clone <repository_url>
    ```
2. プロジェクトフォルダーに移動します：
    ```bash
    cd file-upload-app
    ```
3. 依存関係をインストールします：
    ```bash
    go mod tidy
    ```
4. プロジェクトルートに `config.yaml` という設定ファイルを作成し、以下の内容を追加します：
    ```yaml
    upload_path: "./uploads"
    ```
5. アプリケーションを実行します：
    ```bash
    go run main.go
    ```

## 使用方法
- ウェブブラウザで `http://localhost:8080` にアクセスします
- アップロードページを使用してファイルをアップロードします
- ホームページでアップロードされたファイルを表示します

## プロジェクト構造
- `main.go`: アプリケーションのエントリーポイント
- `config/`: 設定関連のファイル
- `handlers/`: アプリケーションのルートハンドラ
- `utils/`: ユーティリティ関数
- `templates/`: HTMLテンプレート

## 依存関係
- [Gin](https://github.com/gin-gonic/gin): Goのウェブフレームワーク
- [Viper](https://github.com/spf13/viper): Goの設定管理
- [go-qrcode](https://github.com/skip2/go-qrcode): GoでのQRコード生成
- [Bootstrap](https://getbootstrap.com/): レスポンシブウェブデザインのためのフロントエンドフレームワーク

## ライセンス
このプロジェクトはMITライセンスの下でライセンスされています。
