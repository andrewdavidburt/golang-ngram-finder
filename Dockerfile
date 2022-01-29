FROM golang:1.15

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .
RUN cp /build/moby-dick.txt .

ENTRYPOINT ["/dist/main", "moby-dick.txt"]