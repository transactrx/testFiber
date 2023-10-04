FROM golang:1.21-alpine AS builder

#copy source into builder container
COPY . /app

#install dependencies
RUN cd /app && go mod tidy

RUN cd /app/cmd/testApp && go build -o /app/testApp

FROM alpine:3

RUN mkdir /app

COPY --from=builder /app/testApp /app/testApp

ENTRYPOINT ["/app/testApp"]
