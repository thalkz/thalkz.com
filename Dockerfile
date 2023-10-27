FROM golang:1.21-alpine as builder
RUN apk add build-base
WORKDIR app/
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o /main

FROM alpine as runner
RUN mkdir /home/page
WORKDIR /home
COPY ./static ./static
COPY --from=builder /main ./main
ENTRYPOINT ["./main"]