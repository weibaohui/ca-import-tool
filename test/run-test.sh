#!/bin/bash

# 测试CA证书导入工具的完整流程脚本

echo "=== CA证书导入工具测试流程 ==="
echo

# 检查Podman是否可用
if ! command -v podman &> /dev/null
then
    echo "错误: 未找到Podman，请先安装Podman"
    exit 1
fi

echo "1. 检查测试证书是否存在..."
if [ ! -f "ssl/test.example.com.crt" ] || [ ! -f "ssl/test.example.com.key" ]; then
    echo "   证书不存在，正在生成测试证书..."
    ./generate-cert.sh
else
    echo "   证书已存在"
fi

echo
echo "2. 构建测试环境Docker镜像..."
# 回到项目根目录构建镜像
cd ..
podman build -t ca-test-server test/
cd test

echo
echo "3. 启动测试环境..."
# 停止并删除已存在的容器
podman stop ca-test >/dev/null 2>&1
podman rm ca-test >/dev/null 2>&1
# 启动新容器
podman run -d -p 80:80 -p 443:443 --name ca-test ca-test-server >/dev/null

echo
echo "4. 配置本地域名解析..."
# 检查是否已存在条目
if ! grep -q "127.0.0.1 test.example.com" /etc/hosts; then
    echo "   添加域名解析到 /etc/hosts..."
    echo "127.0.0.1 test.example.com" | sudo tee -a /etc/hosts >/dev/null
else
    echo "   域名解析已存在"
fi

echo
echo "5. 验证测试环境..."
sleep 3  # 等待容器完全启动
if curl -k -s https://test.example.com | grep -q "CA证书测试站点"; then
    echo "   测试环境已成功启动"
else
    echo "   警告: 无法访问测试环境"
fi

echo
echo "=== 测试环境已准备就绪 ==="
echo
echo "接下来的测试步骤："
echo "1. 打开浏览器访问 https://test.example.com"
echo "2. 观察浏览器显示证书不受信任的警告"
echo "3. 回到项目根目录: cd .."
echo "4. 运行CA证书导入工具: ./ca-import-tool test/ssl/test.example.com.crt"
echo "5. 重启浏览器并重新访问 https://test.example.com"
echo "6. 观察浏览器显示受信任的连接"

echo
echo "管理命令："
echo "停止测试环境: make test-docker-stop"
echo "查看容器日志: podman logs ca-test"