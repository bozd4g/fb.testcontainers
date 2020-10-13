FROM golang:1.15-alpine

ENV CGO_ENABLED=0

WORKDIR /src
COPY ../../go.mod ./
RUN go mod download

COPY ../.. .
RUN CGO_ENABLED=0 go build -o /bin/app .

ENTRYPOINT ["/bin/app"]

FROM scratch
COPY --from=0 /bin/app /bin/app
ENTRYPOINT ["/bin/app"]