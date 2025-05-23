ARG GO_VERSION=1
FROM golang:${GO_VERSION}-alpine as builder

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /run-app .

FROM alpine:latest

COPY --from=builder /run-app /usr/local/bin/
COPY --from=builder /usr/src/app/templates/ ./templates/
COPY --from=builder /usr/src/app/css/ ./css/
COPY --from=builder /usr/src/app/js/ ./js/

CMD ["run-app"]
