FROM golang:1.20.2 as base

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/

FROM scratch
COPY --from=base /app/server /server
ENTRYPOINT [ "/server" ]
EXPOSE 3000
