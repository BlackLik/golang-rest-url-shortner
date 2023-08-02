# Windows
set GOOS=linux
set GOARCH=amd64
set CGO_CFLAGS=-g -O2 -w
set CGO_ENABLED=1
go build ./cmd/url-sortener/main.go; ./main.exe;

# Linux
# IF [[ "$OSTYPE" == "linux-gnu"* ]];
# then 
#     go build ./cmd/url-sortener/main.go; ./main;
# fi