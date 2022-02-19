package main

import (
  "github.com/gin-gonic/gin"
  //"github.com/zsais/go-gin-prometheus"
  "github.com/penglongli/gin-metrics/ginmetrics"
)

func main() {
  r := gin.Default()

  // Prometheus exporter
  m := ginmetrics.GetMonitor()
  m.SetMetricPath("/metrics")
  m.Use(r)

  // Routes
  r.GET("/", getRoot)
  r.GET("/hello", getHello)
  r.GET("/joke", getJoke)

  // Fire !
  r.Run()
}
