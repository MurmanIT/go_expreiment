FROM golang:alpine
RUN mkdir /app
COPY hw.go /app
WORKDIR /app
RUN go build -o /app/hw hw.go
ENTRYPOINT ["/app/hw"]