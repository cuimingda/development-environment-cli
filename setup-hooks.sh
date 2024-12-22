#!/bin/sh
# 创建符号链接，将钩子脚本链接到 .git/hooks 目录
ln -s $(pwd)/scripts/hooks/post-commit .git/hooks/post-commit
