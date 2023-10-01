FROM golang:alpine as builder

ADD . /go/src/active-charity-backend
WORKDIR /go/src/active-charity-backend
RUN go mod download

COPY . ./

RUN chmod +x wait-for-postgres.sh

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /active-charity-backend ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /active-charity-backend ./active-charity-backend
RUN mkdir ./config
COPY ./config/default-config.yaml ./config

EXPOSE 65001

ENTRYPOINT ["./active-charity-backend"]
