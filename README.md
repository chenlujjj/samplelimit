# samplelimit

当Prometheus提示 "sample limit exceed" 时，通过本程序可快速找出采集对象的哪个指标的sample数最多。

使用示例：
```go
go run main.go -url http://localhost:9090/metrics

Total samples: 643
==============================
...
prometheus_http_requests_total 14
prometheus_target_sync_length_seconds 14
prometheus_tsdb_compaction_chunk_samples 15
prometheus_tsdb_compaction_chunk_size_bytes 15
prometheus_tsdb_compaction_duration_seconds 17
prometheus_sd_kubernetes_events_total 18
prometheus_engine_query_duration_seconds 20
prometheus_http_response_size_bytes 143
prometheus_http_request_duration_seconds 156
```
