package util

import (
	"net"
	"net/http"
	"net/url"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/rcrowley/go-metrics"
)

// The variables to hold the statistics
var parallelRequestCount int = 0
var timeOfLastNOK time.Time = time.Now()
var timeOfLastOK time.Time = time.Now()

var c metrics.Meter

// An extension of the ResponseWriter that can record the
// HTTP status codes sent out
// Design taken from Gorilla
type statLogger struct {
	w      http.ResponseWriter
	size   int
	status int
}

type statResponseWriter interface {
	http.ResponseWriter
	Status() int
	Size() int
}

func MakeLogger(w http.ResponseWriter) statResponseWriter {
	var l statResponseWriter = &statLogger{w: w}
	return l
}

func (l *statLogger) Header() http.Header {
	return l.w.Header()
}

func (l *statLogger) Write(b []byte) (int, error) {
	// 200 if nothing has been set
	if l.status == 0 {
		l.status = http.StatusOK
	}
	size, err := l.w.Write(b)
	l.size += size
	return size, err
}

func (l *statLogger) WriteHeader(s int) {
	l.w.WriteHeader(s)
	l.status = s
}

func (l *statLogger) Status() int {
	return l.status
}

func (l *statLogger) Size() int {
	return l.size
}

// HTTP logging handler
func LoggingHandler(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sw := MakeLogger(w)
		inner.ServeHTTP(sw, r)

		logHTTP(r, *r.URL, time.Now(), sw.Status(), sw.Size())
	})
}

// buildCommonLogLine builds a log entry for req in Apache Common Log Format.
// ts is the timestamp with which the entry should be logged.
// status and size are used to provide the response HTTP status and size.
// Copied from gorilla/handlers/handlers.go, converted to strings instead of byte[]
func logHTTP(req *http.Request, url url.URL, ts time.Time, status int, size int) {
	host, _, err := net.SplitHostPort(req.RemoteAddr)

	if err != nil {
		host = req.RemoteAddr
	}

	uri := req.RequestURI

	// Requests using the CONNECT method over HTTP/2.0 must use
	// the authority field (aka r.Host) to identify the target.
	// Refer: https://httpwg.github.io/specs/rfc7540.html#CONNECT
	if req.ProtoMajor == 2 && req.Method == "CONNECT" {
		uri = req.Host
	}
	if uri == "" {
		uri = url.RequestURI()
	}

	log.WithFields(log.Fields{
		"host": host,
		"method": req.Method,
		"uri": uri,
		"proto": req.Proto,
		"status": status,
		"size": size,
	}).Info(status)
}

// Stats is a handler for displaying API statistics
func StatsHandler(w http.ResponseWriter, r *http.Request) {
	SendHeaders(w)
	w.WriteHeader(http.StatusOK)
	metrics.WriteJSONOnce(metrics.DefaultRegistry, w)
}

func SendHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}
