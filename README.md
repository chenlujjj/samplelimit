# samplelimit

当Prometheus提示 "sample limit exceed" 时，通过本程序可快速找出采集对象的哪个指标的sample数最多。

使用示例：
```go
go run main.go -url http://localhost:9100/metrics
```
