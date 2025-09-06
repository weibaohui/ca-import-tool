#!/bin/bash

# CA证书自动导入工具使用示例

echo "CA证书自动导入工具使用示例"
echo "========================"

# 显示帮助信息
echo "1. 显示帮助信息:"
./bin/ca-import-tool -h

echo -e "\n2. 查看版本信息:"
./bin/ca-import-tool --version

echo -e "\n3. 导入证书到系统信任库:"
echo "注意：此操作需要管理员权限"
# ./bin/ca-import-tool example.crt

echo -e "\n4. 同时为Docker配置证书:"
echo "注意：此操作需要管理员权限"
# ./bin/ca-import-tool -d registry.example.com example.crt

echo -e "\n注意事项:"
echo "- 所有系统级操作都需要管理员权限"
echo "- 在Windows上以管理员身份运行此工具"
echo "- 在macOS和Linux上可能需要输入sudo密码"
echo "- 请使用真实的CA证书文件替换example.crt"