npm run build
GOOS=linux go build
docker build -t minh2804/client .
go clean
