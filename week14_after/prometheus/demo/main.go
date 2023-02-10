package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go_8_mage/week14_after/prometheus/demo/collector"
	"net/http"
)

func main() {
	//	新建注册表
	register := prometheus.NewRegistry()

	//	进程采集器
	register.MustRegister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))
	register.MustRegister(prometheus.NewGoCollector())

	// 注册自定义采集
	register.MustRegister(collector.NewDemoCollector())


	//	基于注册表的Http Handler
	http.Handle("/metrics/", promhttp.HandlerFor(register, promhttp.HandlerOpts{Registry: register}))
	//	HTTP 接口暴露
	http.ListenAndServe(":8050", nil)
}
