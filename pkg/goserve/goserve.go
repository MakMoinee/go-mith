package goserve

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spf13/viper"
)

const (
	maxOpenFDS          = 5000
	defaultOpenFDSLimit = 4000
)

var (
	metricHandlers = http.NotFoundHandler()
)

type Service struct {
	Router           *chi.Mux
	maxOpenFDS       int
	port             string
	cert             string
	key              string
	isHealthy        bool
	tls              bool
	profilingEnabled bool
	info             statusInfo
	shutdownHandlers []func()
}

type statusInfo struct {
	StartupTime time.Time
	mu          sync.Mutex
	Info        map[string]interface{}
}

func (c *Service) SetGlobalMetrics(handler http.Handler) {
	metricHandlers = handler
}

func (c *Service) EnableProfiling(isEnable bool) {
	c.profilingEnabled = isEnable
}

func (c *Service) SetInfo(info map[string]interface{}) {
	c.info.Info = info
}

func (c *Service) Start() error {
	c.initShutdownLoops()
	c.initDefaultEndpoints()
	server := http.Server{
		Addr:    c.port,
		Handler: c.Router,
	}
	var err error
	if c.tls == true {
		err = server.ListenAndServeTLS(c.cert, c.key)
	} else {
		err = server.ListenAndServe()
	}
	return err

}

func (c *Service) initDefaultEndpoints() {
	if c.profilingEnabled == true {
		c.Router.Mount("/debug", middleware.Profiler())
	}

	c.Router.Get("/info", func(w http.ResponseWriter, r *http.Request) {
		representaion, _ := json.Marshal(c.info)
		w.Write(representaion)
	})

	c.Router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		if c.isHealthy == true {
			_, err := w.Write([]byte("ok"))
			if err != nil {

			}
		} else {
			http.Error(w, http.StatusText(500), 500)
		}
	})

	c.Router.Handle("/prometheus", metricHandlers)
}

func (c *Service) setHealthy(isHealthy bool) {
	c.isHealthy = isHealthy
}

func (c *Service) initShutdownLoops() {
	sigChan := make(chan os.Signal, 0)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			select {
			case sig := <-sigChan:
				for _, shutdownHandler := range c.shutdownHandlers {
					shutdownHandler()
				}
				c.setHealthy(false)
				var stats runtime.MemStats
				runtime.ReadMemStats(&stats)
				exitReport := fmt.Sprintf(" | exitSignal = %s | goRoutineCount = %d | memAllocBytes = %d"+
					"| heapAllocBytes = %d | heapObjects = %d | stacksInUse = %d | killedAfter = %s |", sig, runtime.NumGoroutine(), stats.Alloc,
					stats.HeapAlloc, stats.HeapObjects, stats.StackInuse, time.Since(c.info.StartupTime))
				log.Println(exitReport)
				os.Exit(1)
			}
		}

	}()
}

func GetService(r *chi.Mux, port string, cert string, key string, tlsEnabled bool, mInfo map[string]interface{}) *Service {
	customService := Service{
		Router:           r,
		maxOpenFDS:       defaultOpenFDSLimit,
		port:             port,
		cert:             cert,
		key:              key,
		isHealthy:        true,
		tls:              tlsEnabled,
		shutdownHandlers: make([]func(), 0),
		info: statusInfo{
			StartupTime: time.Now(),
			Info:        mInfo,
		},
	}

	return &customService
}

func NewService(port string) *Service {
	// config.Set()
	// tlsEnabled := config.Registry.GetBool("SERVER_SSL_ENABLED")
	tlsEnabled := false // uncomment above if you have config setup in your code
	cert := viper.GetString("cert_file")
	key := viper.GetString("key_file")
	r := chi.NewRouter()

	customService := Service{
		Router:           r,
		maxOpenFDS:       defaultOpenFDSLimit,
		port:             port,
		cert:             cert,
		key:              key,
		isHealthy:        true,
		tls:              tlsEnabled,
		shutdownHandlers: make([]func(), 0),
		info: statusInfo{
			StartupTime: time.Now(),
			Info:        make(map[string]interface{}),
		},
	}

	return &customService
}
