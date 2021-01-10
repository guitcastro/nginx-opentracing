FROM ubuntu:18.04

ENV ELASTIC_APM_SERVER_URL ***REMOVED***
ENV ELASTIC_APM_SECRET_TOKEN ***REMOVED***

WORKDIR /app
ADD . /app
ENV GOPATH="/app"
RUN set -x \
  && apt-get update \
  && apt-get install --no-install-recommends --no-install-suggests -y \
              golang \
              ca-certificates \
              git \
              curl 
RUN apt-get install -y gcc
              
 RUN export PATH=$PATH:$GOPATH/bin \
  && mkdir -p $GOPATH/bin \
  && curl https://glide.sh/get | sh \
  && cd src/hello-backend \
  && glide up \
  && glide install \
  && go build -o /app/server server.go

EXPOSE 9001
CMD ["/app/server"]
