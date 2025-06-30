# learn-go

# 创建工程
mkdir lin-go-config-viper
 cd .\lin-go-config-viper\
go mod init github.com/langyan/lin-go-config-viper

E:\source_files\learn-go\lin-go-config-viper> go run main.go

# docker 启动
docker-compose up -d


# 在 Go 项目中添加依赖包

在 Go 中，你可以使用 `go get` 命令或直接编辑 `go.mod` 文件来添加依赖包。以下是几种常见方法：

## 1. 使用 `go get` 命令

```bash
# 添加最新版本
go get example.com/package

# 添加特定版本
go get example.com/package@v1.2.3

# 添加 master 分支最新代码
go get example.com/package@master
```

## 2. 直接编辑 go.mod 文件

1. 打开 `go.mod` 文件
2. 在 `require` 部分添加你需要的包和版本：
   ```
   require example.com/package v1.2.3
   ```
3. 运行 `go mod tidy` 下载依赖并更新 `go.sum` 文件

## 3. 使用 `go mod tidy`

这个命令会根据你代码中实际导入的包来自动添加依赖：

```bash
go mod tidy
```

## 4. 添加本地模块

如果你要添加本地开发的模块：

```bash
go mod edit -replace example.com/package=../local/package
```

## 注意事项

1. 确保你的项目已经初始化了 Go 模块：
   ```bash
   go mod init yourmodule
   ```

2. 添加依赖后，Go 会自动更新 `go.mod` 和 `go.sum` 文件

3. 你可以使用 `go list -m all` 查看所有依赖

4. 要移除未使用的依赖，可以运行 `go mod tidy`

5. 对于私有仓库，你可能需要设置 `GOPRIVATE` 环境变量：
   ```bash
   export GOPRIVATE=example.com/private/*
   ```

希望这些信息对你有帮助！如果你有特定的依赖包需要添加，可以提供更多细节，我可以给出更具体的指导。