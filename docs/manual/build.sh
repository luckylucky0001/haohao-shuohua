#!/usr/bin/env bash
# 生成三页使用手册 PDF。需要 WeasyPrint（pip install weasyprint）和 Noto Sans/Serif CJK SC 字体。
set -euo pipefail
cd "$(dirname "$0")"
weasyprint quick.html haohao-shuohua-manual.pdf
echo "已生成 $(pwd)/haohao-shuohua-manual.pdf"
