CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
echo "打包完成"
docker run -v $(pwd)/:/app -p 8000:8000 -p 8001:8001 -p 8002:8002 -p 3306:3306 -p 6379:6379 alpine .//app/main
