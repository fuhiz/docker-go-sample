FROM golang:1.15

## 作業ディレクトリ
WORKDIR /app

# モジュール管理のファイルをコピー
COPY app/go.mod .
COPY app/go.sum .

# 外部パッケージのダウンロード
RUN go mod download

EXPOSE 9000
