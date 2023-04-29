#!/bin/bash

# 参数解释： 
# -listenAddr: 本地 OpenAI API http 代理监听地址（非 https），后续使用 OpenAI API 时将 URL https://api.openai.com 改写为 http://{listenAddr} 即可。
# -httpProxy: 本地 OpenAI API http 代理转发地址，即 gost 本地 HTTP Client 监听的端口，即 1080
gost_ai -listenAddr '127.0.0.1:1090' -httpProxy 'http://127.0.0.1:1080' -L :1080?bypaas=~/.ssh/gost-bypass.txt -F 'http2://USER:NAME@example.com:443' 
