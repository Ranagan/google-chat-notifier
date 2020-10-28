# To Run

## Build yourself
```bash
go build main.go

./main -url=<webhook URL> -message=<message to send> -threadKey=<optional thread key>
```

## Docker
```bash
docker build -t google-chat-notifier

docker run google-chat-notifier -url=<webhook URL> -message=<message to send> -threadKey=<optional thread key>
```