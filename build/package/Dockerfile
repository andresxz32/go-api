FROM golang:1.17-alpine3.15 AS builder

ENV GO111MODULE=on 

WORKDIR /go/src/github.com/andresxz32/go-api
 
COPY . .

ENV TZ America/Chicago

#dependencies
RUN go mod tidy 

WORKDIR /go/src/github.com/andresxz32/go-api/cmd/go-api
 
RUN GOOS=linux go build -a -installsuffix cgo -o app .


#Second stage for a smaller image
FROM alpine:3.15
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/andresxz32/go-api/cmd/go-api/app .
CMD ["./app"]