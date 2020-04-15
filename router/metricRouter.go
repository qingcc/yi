package router

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	_ "net/http/pprof"
)

//开启一个端口采集数据，可用于分析性能（cpu，内存， goroutine使用情况等等）
//访问 http://localhost:8081/debug/pprof/ or http://addr/debug/pprof/
func InitMetrics(addr string) {
	if addr == "" {
		addr = ":8081"
	}
	metricsRouter := http.DefaultServeMux
	metricsRouter.Handle("/metrics", promhttp.Handler())
	metricsRouter.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	adminServer := &http.Server{
		Addr:    addr,
		Handler: metricsRouter,
	}

	if err := adminServer.ListenAndServe(); err != nil {
		println("ListenAndServe metrics: ", err.Error())
	}
}
