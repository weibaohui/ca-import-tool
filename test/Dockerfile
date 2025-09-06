# 用于测试CA证书导入工具的Nginx服务器
FROM nginx:alpine

# 复制nginx配置文件
COPY nginx.conf /etc/nginx/nginx.conf

# 创建SSL证书目录
RUN mkdir -p /etc/nginx/ssl

# 复制SSL证书和私钥
COPY ssl/test.example.com.crt /etc/nginx/ssl/test.example.com.crt
COPY ssl/test.example.com.key /etc/nginx/ssl/test.example.com.key

# 创建网站根目录
RUN mkdir -p /usr/share/nginx/html

# 复制测试页面
COPY index.html /usr/share/nginx/html/index.html

# 设置权限
RUN chmod 644 /usr/share/nginx/html/index.html
RUN chmod 644 /etc/nginx/ssl/test.example.com.crt
RUN chmod 644 /etc/nginx/ssl/test.example.com.key

# 暴露端口
EXPOSE 80 443

# 启动nginx
CMD ["nginx", "-g", "daemon off;"]