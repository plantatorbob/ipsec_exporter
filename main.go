package main

import (
	"log"
	"net/http"

	"github.com/plantatorbob/strongswan_exporter/app"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	log.Println("Up and Running")
	strongswanCollector := app.NewStrongswanCollector()
	strongswanCollector.Init()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatalln(http.ListenAndServe(":9814", nil)) //nolint:gosec
}
