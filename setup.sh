#!/bin/sh

mkdir -p bin
go build -o bin/dev ./cmd

if [ "$(id -u)" -eq 0 ]; then
    ln -sf $(pwd)/bin/dev /usr/local/bin/dev
else
    sudo ln -sf $(pwd)/bin/dev /usr/local/bin/dev

    # 创建符号链接，将钩子脚本链接到 .git/hooks 目录
    find .git/hooks -type l -exec rm {} +
    ln -s $(pwd)/scripts/hooks/pre-commit .git/hooks/pre-commit
    ls -l .git/hooks | grep '^l'
fi
