# ==============================
# ビルドの用意
# ==============================
FROM golang:1.15.6 AS builder

WORKDIR /go/src/app
COPY . .

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build -v -o /go/bin/app ./pkg/server/...

# ==============================
# イメージ作成
# ==============================
FROM gcr.io/distroless/base AS runner
COPY --from=builder /go/bin/app /app

ENTRYPOINT ["/app"]