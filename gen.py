# -*- coding: utf-8 -*-
import requests

mode  = 'clash'
mode = 'quan x'
urls_proxy = []
urls_direct = []

# Custom
s = 'googleapis.com kamihq.com bugsnag.com spotify.com sentry.io scdn.co akamaized.net pximg.net pixiv.net pixiv.org sndcdn.com soundcloud.com google.com steampowered.com unsplash.com steamcommunity.com arxiv.org live.com office.com v2ex.com semanticscholar.org tensorflow.org unsplash.com twitter.com linkedin.com reddit.com leetcode.com google.com mathworks.com rule34.xxx tumblr.com twimg.com python.org discord.com discordapp.com discord.gg discordapp.net facebook.net mypikpak.com notion.so lencr.org crl.comodoca.com duosecurity.com overleaf.com loggly.com githubassets.com googleoptimize.com spotifycdn.com cookielaw.org googletagmanager.com google-analytics.com ads-twitter.com sc-static.net qualaroo.com fastly-insights.com scorecardresearch.com doubleclick.net rlcdn.com demdex.net go.dev microsoft.com github.dev visualstudio.com vscode-cdn.net exp-tas.com msecnd.net azureedge.net vsassets.io laowang.vip msauth.net windows.net github.com apache.org confluent.io golang.org herokuapp.com heroku.com'
urls_proxy.extend(s.split(' '))

# DIRECT
s = '163.com bilibili.com bilivideo.cn bilivideo.com biliapi.net baidu.com 189.cn aliyundrive.com zhihu.com amap.com huya.com jd.com taobao.com miobt.com'
urls_direct.extend(s.split(' '))

## do not need to change
# Games
s = 'amp-api-edge.apps.apple.com push.apple.com inappcheck.itunes.apple.com app-measurement.com nexoncdn.co.kr nexon.com nexon.io'
urls_proxy.extend(s.split(' '))

# Ghelper Default
s = 'youtube.com googlevideo.com ytimg.com youtu.be yimg.com behance.net instagram.com t.co facebook.com instagram.com t.co ebay.co yahoo.co.jp google.com google.de google.fr google.com.hk google.com.sg google.com.tw google.co.jp google.com.za gstatic.com googleusercontent.com gvt2.com gvt3.com ggpht.com googleapis.com chrome.com googleadservices.com googleusercontent.com googlesyndication.com googlesource.com wikipedia.org pbskids.org dropbox.com googleblog.com appspot.com android.com github.com gmail.com wikimedia.org googlegroups.com githubusercontent.com arxiv.org mediawiki.org instagram.com t.co'
urls_proxy.extend(s.split(' '))

urls_proxy = set(urls_proxy)
urls_direct = set(urls_direct)
if mode == 'clash':
    res = requests.get("https://ghelper.me/clash/4ff905a0af778432e7dd12dfe21788cb")
    proxy = res.text.split('rules:')[0]
    print(proxy)
    print('rules:')
    for token in urls_proxy:
        print('- DOMAIN-SUFFIX,{},Ghelper'.format(token))
    for token in urls_direct:
        print('- DOMAIN-SUFFIX,{},DIRECT'.format(token))
    s = '# default\n- DOMAIN-SUFFIX,ip-api.com,DIRECT\n- DOMAIN-SUFFIX,ipip.net,DIRECT\n- DOMAIN-SUFFIX,ip138.com,DIRECT\n- DOMAIN-SUFFIX,stunnel.vip,DIRECT\n- DOMAIN-SUFFIX,gotochinatown.net,DIRECT\n- DOMAIN-SUFFIX,ghelper.net,DIRECT\n- DOMAIN-SUFFIX,ghelper.me,DIRECT\n- DOMAIN-SUFFIX,ghelper.xyz,DIRECT\n- DOMAIN-SUFFIX,ghelper.org,DIRECT\n- DOMAIN-SUFFIX,fastapi.me,DIRECT\n- DOMAIN-SUFFIX,vps315.com,DIRECT\n- DOMAIN-SUFFIX,copyplay.net,DIRECT\n- DOMAIN-SUFFIX,pickdown.net,DIRECT\n- DOMAIN-KEYWORD,google,Ghelper\n- DOMAIN-KEYWORD,youtube,Ghelper\n- DOMAIN-SUFFIX,ggpht.com,Ghelper\n- DOMAIN-SUFFIX,gmail.com,Ghelper\n- DOMAIN-SUFFIX,gvt2.com,Ghelper\n- DOMAIN-SUFFIX,gvt3.com,Ghelper\n- DOMAIN-SUFFIX,chrome.com,Ghelper\n- DOMAIN-SUFFIX,wikipedia.org,Ghelper\n- DOMAIN-SUFFIX,wikimedia.org,Ghelper\n- DOMAIN-SUFFIX,appspot.com,Ghelper\n- DOMAIN-SUFFIX,android.com,Ghelper\n- DOMAIN-SUFFIX,github.com,Ghelper\n- DOMAIN-SUFFIX,gitbook.com,Ghelper\n- DOMAIN-SUFFIX,kaggle.com,Ghelper\n- DOMAIN-SUFFIX,arxiv.org,Ghelper\n- DOMAIN-SUFFIX,wiktionary.org,Ghelper\n- DOMAIN-SUFFIX,blogger.com,Ghelper\n- DOMAIN-SUFFIX,youtu.be,Ghelper\n- DOMAIN-SUFFIX,ytimg.com,Ghelper\n- DOMAIN-SUFFIX,youtube.com,Ghelper\n- DOMAIN-SUFFIX,instagram.com,Ghelper\n- DOMAIN-SUFFIX,twitter.com,Ghelper\n- DOMAIN-SUFFIX,t.co,Ghelper\n- DOMAIN-SUFFIX,facebook.com,Ghelper\n- DOMAIN-SUFFIX,telegram.org,Ghelper\n- IP-CIDR,192.168.0.0/16,DIRECT\n- IP-CIDR,10.0.0.0/8,DIRECT\n- IP-CIDR,172.16.0.0/12,DIRECT\n- IP-CIDR,172.23.0.0/12,DIRECT\n- IP-CIDR,127.0.0.0/8,DIRECT\n- IP-CIDR,100.64.0.0/10,DIRECT\n- IP-CIDR6,::1/128,DIRECT\n- IP-CIDR6,fc00::/7,DIRECT\n- IP-CIDR6,fe80::/10,DIRECT\n- IP-CIDR6,fd00::/8,DIRECT\n- GEOIP,CN,DIRECT\n- MATCH,DIRECT'
    print(s)
if mode == 'quan x':
    for token in urls_proxy:
        print('host-suffix,{},proxy'.format(token))
    for token in urls_direct:
        print('host-suffix,{},direct'.format(token))
    s = 'ip-cidr, 10.0.0.0/8, direct\nip-cidr, 127.0.0.0/8, direct\nip-cidr, 172.16.0.0/12, direct\nip-cidr, 192.168.0.0/16, direct\nip-cidr, 224.0.0.0/24, direct\ngeoip, cn, direct\nfinal, direct'
    print(s)