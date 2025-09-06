#!/bin/bash

# 本地运行测试服务器的脚本 (不使用Docker)

# 检查是否安装了nginx
if ! command -v nginx &> /dev/null
then
    echo "未找到nginx，请先安装nginx:"
    echo "  macOS: brew install nginx"
    echo "  Ubuntu: sudo apt install nginx"
    echo "  CentOS: sudo yum install nginx"
    exit 1
fi

# 检查是否安装了openssl
if ! command -v openssl &> /dev/null
then
    echo "未找到openssl，请先安装openssl"
    exit 1
fi

# 创建临时目录
TEMP_DIR="/tmp/ca-test"
mkdir -p $TEMP_DIR

# 复制配置文件
cp nginx.conf $TEMP_DIR/nginx.conf
cp index.html $TEMP_DIR/index.html

# 创建SSL目录
mkdir -p $TEMP_DIR/ssl

# 复制证书文件（如果存在）
if [ -f "ssl/test.example.com.crt" ] && [ -f "ssl/test.example.com.key" ]; then
    cp ssl/test.example.com.crt $TEMP_DIR/ssl/
    cp ssl/test.example.com.key $TEMP_DIR/ssl/
    echo "使用现有的证书文件"
else
    echo "生成自签名证书..."
    openssl genrsa -out $TEMP_DIR/ssl/test.example.com.key 2048
    openssl req -new -key $TEMP_DIR/ssl/test.example.com.key -out $TEMP_DIR/ssl/test.example.com.csr -subj "/C=CN/ST=Beijing/L=Beijing/O=Example Corp/OU=IT Department/CN=test.example.com"
    openssl x509 -req -days 365 -in $TEMP_DIR/ssl/test.example.com.csr -signkey $TEMP_DIR/ssl/test.example.com.key -out $TEMP_DIR/ssl/test.example.com.crt
fi

# 更新nginx配置中的路径
sed -i '' "s|/etc/nginx/ssl|$TEMP_DIR/ssl|g" $TEMP_DIR/nginx.conf
sed -i '' "s|/usr/share/nginx/html|$TEMP_DIR|g" $TEMP_DIR/nginx.conf

# 添加域名解析
if ! grep -q "127.0.0.1 test.example.com" /etc/hosts; then
    echo "请将以下行添加到 /etc/hosts 文件中:"
    echo "127.0.0.1 test.example.com"
    echo ""
    echo "在macOS上可以使用以下命令:"
    echo "echo '127.0.0.1 test.example.com' | sudo tee -a /etc/hosts"
fi

echo ""
echo "启动nginx测试服务器:"
echo "sudo nginx -c $TEMP_DIR/nginx.conf"
echo ""
echo "然后访问: https://test.example.com"
echo ""
echo "停止nginx服务器:"
echo "sudo nginx -c $TEMP_DIR/nginx.conf -s stop"