# yatter

## How to use

```sh
# Create and start containers
docker-compose up

# デタッチモードで起動
# Detached mode: Run containers in the background
docker-compose up -d
```

## Swagger UI

APIのリソースを可視化して操作できるツール。
OpenAPI仕様(標準化されたAPIの記述方法)から自動的に作成される。

http://localhost:8081

### Swagger UIの見方

openapi.ymlで記述した情報を見ることができるようになる

- health endpoint
- アプリケーションの健全性を伝えるアプリケーション内のHTTPエンドポイント

- endpoint
  - URL/URIのこと