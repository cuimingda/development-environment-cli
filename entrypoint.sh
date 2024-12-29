#!/bin/sh

# 显示 Go 版本
echo
echo "[curl]"
curl --version

echo
echo "[jq]"
jq --version

echo
echo "[go]"
go version

echo
echo "[Node.js]"
node --version

echo
echo "[npm]"
npm --version

echo
echo "[PHP]"
php --version

echo
exec "$@"
