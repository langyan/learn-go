
读取环境变量和配置文件

### 创建工程
mkdir lin-go-config-viper
 cd .\lin-go-config-viper\
go mod init github.com/langyan/lin-go-config-viper

### 管理依赖项 加载viper
go get github.com/spf13/viper


### 示例：使用 Viper 的配置
```
package main 

import ( 
    "fmt" 
    "github.com/spf13/viper" 
    "log"
 ) 

func  loadConfig () { 
    viper.SetConfigName( "config" ) 
    viper.SetConfigType( "yaml" ) 
    viper.AddConfigPath( "." ) 
    viper.AutomaticEnv() 
    if err := viper.ReadInConfig(); err != nil { 
        log.Fatalf( "加载配置错误：%v" , err) 
    } 
} 

func  main () { 
    loadConfig() 
    port := viper.GetInt( "server.port" ) 
    env := viper.GetString( "app.env" ) 
    fmt.Printf( "服务器在 %s 环境中的 %d 端口上运行\n" , port, env) 
}

```

### 示例配置（config.yaml）

```
server:
  port: 8080
app:
  env: production
```

