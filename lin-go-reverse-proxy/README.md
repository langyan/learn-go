循环负载均衡：使用 sync/atomic 进行线程安全索引循环。
反向代理：用于 HTTP 转发。httputil.NewSingleHostReverseProxy()
自定义错误处理：防止客户端与信息丰富的错误日志混淆。