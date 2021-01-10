FROM golang:1.15.2-alpine
ENV ELASTIC_APM_SERVER_URL ***REMOVED***
ENV ELASTIC_APM_SECRET_TOKEN ***REMOVED***

WORKDIR /src
ENV CGO_ENABLED=0
RUN apk add --no-cache \
      git \
      ca-certificates \
      tzdata \
 && update-ca-certificates
COPY . .
RUN go build -o /app .

ENTRYPOINT ["/app"]