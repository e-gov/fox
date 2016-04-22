package util
import(
	"net"
	"net/http"
	"net/url"
	"time"
	"fmt"
	logging "github.com/op/go-logging"
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
	w 		http.ResponseWriter
	size int
	status	int
}

type statResponseWriter interface {
	http.ResponseWriter
	Status() int
	Size() int
}

func MakeLogger(w http.ResponseWriter) statResponseWriter{
	var l statResponseWriter = &statLogger{w: w}
	return l
}

func (l *statLogger) Header() http.Header{
	return l.w.Header()
}

func (l *statLogger) Write(b []byte) (int, error){
	// 200 if nothing has been set
	if l.status == 0 {
		l.status = http.StatusOK
	}
	size, err := l.w.Write(b)
	l.size += size
	return size, err
}

func (l *statLogger) WriteHeader(s int){
	l.w.WriteHeader(s)
	l.status = s
}

func (l *statLogger) Status() int{
	return l.status
}

func (l *statLogger) Size() int{
	return l.size
}

// HTTP logging handler
func LoggingHandler(inner http.Handler, log *logging.Logger) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		sw := MakeLogger(w)
		inner.ServeHTTP(sw, r)
		log.Info(buildCommonLogLine(r, *r.URL, time.Now(), sw.Status(), sw.Size()))
	})
}


// buildCommonLogLine builds a log entry for req in Apache Common Log Format.
// ts is the timestamp with which the entry should be logged.
// status and size are used to provide the response HTTP status and size.
// Copied from gorilla/handlers/handlers.go, converted to strings instead of byte[]
func buildCommonLogLine(req *http.Request, url url.URL, ts time.Time, status int, size int) string {
	username := "-"
	if url.User != nil {
		if name := url.User.Username(); name != "" {
			username = name
		}
	}

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

	return fmt.Sprintf("%s %s [%s] \"%s %s %s\" %d %d", 
		host, 
		username, 
		ts.Format("02/Jan/2006:15:04:05 -0700"), 
		req.Method, 
		uri, 
		req.Proto, 
		status, 
		size)
}
