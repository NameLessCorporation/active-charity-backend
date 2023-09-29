FROM golang:alpine as builder

ADD . /go/src/active-charity-backend
WORKDIR /go/src/active-charity-backend
RUN go mod download

COPY . ./

RUN chmod +x wait-for-postgres.sh

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /active-charity-backend ./cmd/app/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /active-charity-backend ./active-charity-backend
RUN mkdir ./configs
COPY ./configs/default-config.yaml ./configs

EXPOSE 65000

ENTRYPOINT ["./active-charity-backend"]
