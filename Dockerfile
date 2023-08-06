FROM golang:1.20-alpine

WORKDIR /backend 

COPY . .

RUN apk add --no-cache gcc musl-dev

RUN go mod download && go mod verify
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o /bin/main ./cmd/url-sortener/main.go 

EXPOSE 8080

CMD ["main"]