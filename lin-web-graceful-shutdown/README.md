

### 使用信号处理正常关闭
我们的一个服务在 Kubernetes 滚动更新期间崩溃，导致数据丢失和事务不完整。在与 Gopher Slack 上的其他开发人员讨论后，我意识到没有正确处理信号是根本原因。

### 问题：
当 Kubernetes 发送SIGTERM来终止 pod 时，我们的 Go 服务会立即关闭，而不会完成正在进行的任务。

### 解决方法：使用 Context 优雅关闭
实现信号处理可确保服务在关闭之前完成当前工作。

```
package main 

import ( 
    "context" 
    "fmt" 
    "os" 
    "os/signal" 
    "syscall" 
    "time"
 ) 

func  runServer (ctx context.Context) { 
    for { 
        select { 
        case <-ctx.Done(): 
            fmt.Println( "正在正常关闭..." ) 
            return 
        default : 
            fmt.Println( "正在处理..." ) 
            time.Sleep( 500 * time.Millisecond) 
        } 
    } 
} 

func  main () { 
    ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM) 
    defer stop() 
    go runServer(ctx) 
    <-ctx.Done() 
    fmt.Println( "服务已停止。" ) 
}
```