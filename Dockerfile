FROM golang AS builder
COPY . /app
WORKDIR /app
RUN go build -v -o wormhole

FROM golang
COPY --from=builder /app/wormhole /app/wormhole
ENTRYPOINT ["/app/wormhole"]  
