FROM golang:1.20-alpine

WORKDIR /backend

COPY . .

RUN go mod download && go mod verify
RUN go build -o application ./cmd/url-sortener/main.go
RUN chmod +x application    # Add this line to make the binary executable

EXPOSE 8080

CMD [ "./application" ]
