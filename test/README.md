# CA证书导入工具测试环境

这个测试环境用于验证CA证书导入工具的功能，包括证书导入前后的对比效果。

## 目录结构

```
test/
├── Dockerfile          # Nginx测试服务器的Docker配置
├── nginx.conf          # Nginx配置文件
├── index.html          # 测试页面
├── generate-cert.sh    # 生成自签名证书的脚本
├── run-local.sh        # 本地运行测试服务器的脚本(不使用Docker)
├── README.md           # 本说明文件
└── ssl/                # 生成的证书文件存放目录
```

## 使用方法

### 1. 生成自签名证书

```bash
cd test
./generate-cert.sh
```

这将生成以下文件：
- `ssl/test.example.com.key` - 私钥
- `ssl/test.example.com.crt` - 自签名证书

### 2. 使用Podman/Docker运行测试服务器（推荐）

```bash
# 构建Docker镜像
podman build -t ca-test-server test/
# 或者如果使用Docker:
# docker build -t ca-test-server test/

# 运行容器
podman run -d -p 80:80 -p 443:443 --name ca-test ca-test-server
# 或者如果使用Docker:
# docker run -d -p 80:80 -p 443:443 --name ca-test ca-test-server
```

### 3. 配置本地域名解析

```bash
echo "127.0.0.1 test.example.com" | sudo tee -a /etc/hosts
```

### 4. 测试证书导入前的状态

1. 打开浏览器访问: https://test.example.com
2. 浏览器应显示证书不受信任的警告

### 5. 使用CA证书导入工具

```bash
# 回到项目根目录
cd ..

# 使用工具导入证书
./ca-import-tool test/ssl/test.example.com.crt
```

### 6. 测试证书导入后的状态

1. 重启浏览器
2. 再次访问: https://test.example.com
3. 浏览器应显示受信任的连接，无安全警告

## 管理测试环境

### 停止测试服务器

```bash
podman stop ca-test
# 或者如果使用Docker:
# docker stop ca-test
```

### 删除测试容器

```bash
podman rm ca-test
# 或者如果使用Docker:
# docker rm ca-test
```

### 删除测试镜像

```bash
podman rmi ca-test-server
# 或者如果使用Docker:
# docker rmi ca-test-server
```

## 注意事项

1. 在Linux和macOS系统上可能需要sudo权限来修改/etc/hosts文件
2. 测试完成后，建议从系统信任库中移除测试证书
3. 如果需要重新测试，可以删除并重新生成证书
4. 在生产环境中不要使用自签名证书

## 故障排除

### 证书仍然不被信任

1. 确认证书已正确导入系统信任库
2. 检查工具执行时是否有错误信息
3. 确认操作系统已更新证书库
4. 重启浏览器和系统后再试

### 无法访问测试站点

1. 检查测试服务器是否正在运行: `podman ps`
2. 检查端口是否被占用: `sudo lsof -i :80,443`
3. 确认域名解析是否正确: `ping test.example.com`
4. 检查防火墙设置是否阻止了端口访问