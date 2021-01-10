package main

import (
	"flag"
	opentracing "github.com/opentracing/opentracing-go"
	"go.elastic.co/apm/module/apmhttp"
	"log"
	"net/http"
	"time"
)

const (
	serviceName   = "hello-server"
	hostPort      = "0.0.0.0:0"
	debug         = false
	sameSpan      = false
	traceID128Bit = true
)

func handler(w http.ResponseWriter, r *http.Request) {
	wireContext, _ := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(r.Header))
	span := opentracing.StartSpan(
		"/",
		opentracing.ChildOf(wireContext))
	defer span.Finish()
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is " + tm))
}

func main() {
	flag.Parse()
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LocalAgentHostPort:  *collectorHost + ":" + *collectorPort,
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}

	var myHandler http.Handler = handler
	tracedHandler := apmhttp.Wrap(myHandler)
	http.HandleFunc("/", tracedHandler)
	http.ListenAndServe(":9001", nil)
}
