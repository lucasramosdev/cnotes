FROM golang:1.24.1-alpine3.20 as base
RUN apk update
WORKDIR /src/cnotes
COPY go.mod go.sum ./
COPY . .
RUN ls
RUN go build -o cnotes ./cmd/app

FROM alpine:3.20 as binary
COPY --from=base /src/cnotes/cnotes .
EXPOSE 8080
CMD ["./cnotes"]