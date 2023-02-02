FROM golang:1.19-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY todoAPI.go ./
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o todo .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=0 /app/todo ./
COPY templates /app/templates
CMD ["./todo"]