goctl api -o auth.api
goctl api go -api auth.api -dir auth
go mod init
go mod tidy
go run auth.go -f etc/auth-api.yaml
