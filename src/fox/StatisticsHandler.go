package fox
import(
	"net/http"
	"time"
)

// The variables to hold the statistics
var parallelRequestCount int = 0
var timeOfLastNOK time.Time = time.Now()
var timeOfLastOK time.Time = time.Now()
var nodeName string

// An extension of the ResponseWriter that can record the
// HTTP status codes sent out
// Design taken from Gorilla
type statLogger struct {
	w 		http.ResponseWriter
	status	int
}

type statResponseWriter interface {
	http.ResponseWriter
	Status() int
}

func makeLogger(w http.ResponseWriter) statResponseWriter{
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
	return size, err
}

func (l *statLogger) WriteHeader(s int){
	l.w.WriteHeader(s)
	l.status = s
}

func (l *statLogger) Status() int{
	return l.status
}

// Actual statistics handler
func StatHandler(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		sw := makeLogger(w)

		parallelRequestCount++
		inner.ServeHTTP(sw, r)
		parallelRequestCount--

		// Anything beyond 300 is NOK
		if sw.Status() >= 300{
			timeOfLastNOK = time.Now()
		}else {
			timeOfLastOK = time.Now()
		}
	})
}
