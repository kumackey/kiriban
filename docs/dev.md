# 開発メモ

TODO: makefileにする

## テスト

```bash
go test -v ./...
```

## モック生成

```bash
mockgen -source=../internal/domain/issue_commenter.go -destination=../internal/domain/mock/issue_commenter.go -package=domainmock
```