#!/bin/zsh
echo "开始编译文件"
gf build
echo "编译完成"
echo "开始压缩文件"
upx -9 ./bin/sms_backend
echo "压缩完成"