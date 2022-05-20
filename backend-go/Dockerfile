FROM golang:1.17 AS builder
WORKDIR /code
COPY go.* ./
RUN go mod download && go mod verify
COPY . ./
RUN make # GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" .

FROM gcr.io/distroless/base:debug
COPY --from=builder /code/cmd/kemenag-scraper/kemenag-scraper /go/bin/kemenag-scraper
COPY --from=builder /code/cmd/http-hikarie-islamy-be/http-hikarie-islamy-be /go/bin/http-hikarie-islamy-be
COPY --from=builder /code/configs/backend.toml /etc/hikarie-islamy/backend.toml
ENTRYPOINT [ "/go/bin/http-hikarie-islamy-be" ]
