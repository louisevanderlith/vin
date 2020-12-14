cd cmd/
go mod tidy
GOOS="linux" CGO_ENABLED="0" go build