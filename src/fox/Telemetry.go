package fox

import(
	"net/http"
	"github.com/rcrowley/go-metrics"
	"time"
)

// Telemetry keeps all the meters for a given route
type Telemetry struct{
	rCount metrics.Meter
	sCount metrics.Meter
	fCount metrics.Meter
	
	tmr metrics.Histogram
	
	inner http.Handler
}

// ServeHTTP delegates the serving to the inner handler and updates telemetry
func (t Telemetry)ServeHTTP(w http.ResponseWriter, r *http.Request){
	t.rCount.Mark(1)	
	sw := makeLogger(w)

	start := time.Now()
	t.inner.ServeHTTP(sw, r)
	t.tmr.Update(int64(time.Since(start)/time.Millisecond))

	if sw.Status() >= 300{
		t.fCount.Mark(1)
	}else {
		t.sCount.Mark(1)
	}

}

// NewTelemetry generates a new telemetry handler that is properly initialized for the name
func NewTelemetry(i http.Handler, name string)Telemetry{
	r := metrics.NewMeter()
	s := metrics.NewMeter()
	f := metrics.NewMeter()
	
	tme := metrics.NewHistogram(metrics.NewExpDecaySample(1028, 0.015))
	
	metrics.Register(name + "_requests", r)
	metrics.Register(name + "_success", s)
	metrics.Register(name + "_failure", f)
	metrics.Register(name + "_time", tme)
	
	t := Telemetry{inner: i, rCount:r, sCount:s, fCount:f, tmr: tme}
	return t
}