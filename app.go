package main

import (
	"github.com/grafana/pyroscope-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

const (
	name       = "app"
	serverProf = "pyroscope:4040"
	listenProm = ":9100"
)

func main() {
	go startProm()

	for range time.Tick(1 * time.Second) {
		log.Println("tick")
	}
}

func startProm() {
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(listenProm, nil))
}

func startProf() error {
	_, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: name,

		// replace this with the address of pyroscope server
		ServerAddress:     serverProf,
		BasicAuthUser:     "",
		BasicAuthPassword: "",

		// you can disable logging by setting this to nil
		//Logger: pyroscope.StandardLogger,
		Logger: nil,

		ProfileTypes: []pyroscope.ProfileType{
			// these profile types are enabled by default:
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,

			// these profile types are optional:
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})

	return err
}
