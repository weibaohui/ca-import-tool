#!/bin/bash

# 生成自签名证书用于测试

# 创建SSL证书目录
mkdir -p ssl

# 生成私钥
openssl genrsa -out ssl/test.example.com.key 2048

# 生成证书签名请求(CSR)
openssl req -new -key ssl/test.example.com.key -out ssl/test.example.com.csr -subj "/C=CN/ST=Beijing/L=Beijing/O=Example Corp/OU=IT Department/CN=test.example.com"

# 生成自签名证书
openssl x509 -req -days 365 -in ssl/test.example.com.csr -signkey ssl/test.example.com.key -out ssl/test.example.com.crt

# 显示证书信息
echo "证书已生成:"
echo "私钥: ssl/test.example.com.key"
echo "证书: ssl/test.example.com.crt"
echo ""
echo "证书信息:"
openssl x509 -in ssl/test.example.com.crt -text -noout

echo ""
echo "使用说明:"
echo "1. 将 ssl/test.example.com.crt 作为测试证书导入系统"
echo "2. 使用Docker运行测试服务器:"
echo "   docker build -t ca-test-server ."
echo "   docker run -d -p 80:80 -p 443:443 --name ca-test ca-test-server"
echo "3. 在主机上添加域名解析:"
echo "   echo '127.0.0.1 test.example.com' | sudo tee -a /etc/hosts"
echo "4. 访问 https://test.example.com 测试证书信任状态"