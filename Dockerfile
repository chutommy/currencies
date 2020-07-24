FROM golang:alpine AS builder
LABEL maintainer="tommychu2256@gmail.com"

# set enviroment
ENV GO111MODULE="on"
ENV CGO_ENABLE=0
ENV GOOS="linux"
ENV GOARCH="amd64"

# set working directory
WORKDIR /service

# install dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# build binary
COPY . .
RUN go mod tidy
RUN go build -o main .


FROM alpine:latest AS production

# set enviroment
ENV PORT=10502

# prepare service
RUN apk update && apk add ca-certificates
WORKDIR /service

# prepare binary
COPY --from=builder /service/main .
COPY --from=builder /service/config.yaml .

# run service
EXPOSE $PORT
ENTRYPOINT ["./main"]
