FROM golang:1.17.0-alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /root/go/src/url_shortener

COPY . .

RUN go mod tidy
RUN go build main.go

# Export necessary port
EXPOSE 8081

CMD ["/root/go/src/url_shortener/main"]