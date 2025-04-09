当然可以！下面是上面关于使用 Docker 运行 **Prometheus + Grafana** 的中文翻译。

---

## 🚀 使用 Docker Compose 运行 Prometheus + Grafana

### 🔧 第一步：创建 `docker-compose.yml`

创建一个名为 `docker-compose.yml` 的文件：

```yaml
version: '3.7'

services:
  prometheus:
    image: prom/prometheus
    container_name: prometheus
    volumes:
      - ./prometheus.yml:/etc/prometheus/prometheus.yml
    command:
      - --config.file=/etc/prometheus/prometheus.yml
    ports:
      - "9090:9090"

  grafana:
    image: grafana/grafana-oss
    container_name: grafana
    environment:
      - GF_SECURITY_ADMIN_USER=admin
      - GF_SECURITY_ADMIN_PASSWORD=admin12346
    ports:
      - "3000:3000"
    depends_on:
      - prometheus
    volumes:
      - grafana-storage:/var/lib/grafana

volumes:
  grafana-storage:
```

---

### 📝 第二步：创建 `prometheus.yml` 配置文件

在同一目录下创建一个名为 `prometheus.yml` 的文件：

```yaml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'prometheus'
    static_configs:
      - targets: ['localhost:9090']
```

这个配置的意思是 Prometheus 每 15 秒抓取一次自己的指标数据（你可以后续添加更多目标）。

---

### ▶️ 第三步：运行服务

在终端中运行以下命令：

```bash
docker-compose up -d
```

---

### 🌐 访问 Web 界面

- **Prometheus**：访问 [http://localhost:9090](http://localhost:9090)
- **Grafana**：访问 [http://localhost:3000](http://localhost:3000)
  - 默认用户名：`admin`
  - 默认密码：`admin`

---

### 📊 在 Grafana 中添加 Prometheus 数据源

1. 登录 Grafana
2. 点击左边侧边栏的 ⚙️ **“配置”** → **“数据源”**
3. 点击“添加数据源”
4. 选择 **Prometheus**
5. 把 URL 填写为 `http://prometheus:9090`
6. 点击 **“保存并测试”**

现在你可以在 Grafana 中用 Prometheus 的指标创建可视化仪表盘了！



