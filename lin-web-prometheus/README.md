
### 使用Prometheus和Grafana进行监控
部署后，我们遇到了间歇性的延迟峰值。盲目猜测原因是行不通的，所以我们决定设置指标和监控。

### 问题：
- 缺乏可观察性：无法实时洞察应用程序的性能。
- 调试困难：无法将问题追溯到特定事件。

### 解决方案：集成 Prometheus 以获取指标

- go get github.com/prometheus/client_golang/prometheus


### 测试
- http://localhost:8080
- http://localhost:8080/metrics


### 在 Grafana 中跟踪性能指标

1. docker compose
- docker-compose down -v  #停止并删除现有容器
- docker-compose up -d  # 重新启动服务

2. 