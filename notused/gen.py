MODE  = 'clash'
MODE = 'quan x'
urls_proxy = []
urls_direct = []

DIRECT = ""
# DIRECT
## Microsoft
DIRECT += "microsoft.com microsoftonline.com sharepoint.com office.com office.net "
## Crusaders Quest
DIRECT += "cq.hangame.com cq-pvp.hangame.com cq-cha.hangame.com gslb-gamebase.nhncloudservice.com " # ios
DIRECT += ""
urls_direct.extend(DIRECT.split(' '))

PROXY = "amp-api-edge.apps.apple.com push.apple.com inappcheck.itunes.apple.com app-measurement.com nexoncdn.co.kr nexon.com nexon.io "
# PROXY
urls_proxy.extend(PROXY.split(' '))

urls_proxy = set(urls_proxy)
urls_direct = set(urls_direct)




















# if mode == 'clash':
#     res = requests.get("https://ghelper.me/clash/4ff905a0af778432e7dd12dfe21788cb")
#     proxy = res.text.split('rules:')[0]
#     print(proxy)
#     print('rules:')
#     for token in urls_proxy:
#         print('- DOMAIN-SUFFIX,{},Ghelper'.format(token))
#     for token in urls_direct:
#         print('- DOMAIN-SUFFIX,{},DIRECT'.format(token))
#     s = '# default\n- DOMAIN-SUFFIX,ip-api.com,DIRECT\n- DOMAIN-SUFFIX,ipip.net,DIRECT\n- DOMAIN-SUFFIX,ip138.com,DIRECT\n- DOMAIN-SUFFIX,stunnel.vip,DIRECT\n- DOMAIN-SUFFIX,gotochinatown.net,DIRECT\n- DOMAIN-SUFFIX,ghelper.net,DIRECT\n- DOMAIN-SUFFIX,ghelper.me,DIRECT\n- DOMAIN-SUFFIX,ghelper.xyz,DIRECT\n- DOMAIN-SUFFIX,ghelper.org,DIRECT\n- DOMAIN-SUFFIX,fastapi.me,DIRECT\n- DOMAIN-SUFFIX,vps315.com,DIRECT\n- DOMAIN-SUFFIX,copyplay.net,DIRECT\n- DOMAIN-SUFFIX,pickdown.net,DIRECT\n- DOMAIN-KEYWORD,google,Ghelper\n- DOMAIN-KEYWORD,youtube,Ghelper\n- DOMAIN-SUFFIX,ggpht.com,Ghelper\n- DOMAIN-SUFFIX,gmail.com,Ghelper\n- DOMAIN-SUFFIX,gvt2.com,Ghelper\n- DOMAIN-SUFFIX,gvt3.com,Ghelper\n- DOMAIN-SUFFIX,chrome.com,Ghelper\n- DOMAIN-SUFFIX,wikipedia.org,Ghelper\n- DOMAIN-SUFFIX,wikimedia.org,Ghelper\n- DOMAIN-SUFFIX,appspot.com,Ghelper\n- DOMAIN-SUFFIX,android.com,Ghelper\n- DOMAIN-SUFFIX,github.com,Ghelper\n- DOMAIN-SUFFIX,gitbook.com,Ghelper\n- DOMAIN-SUFFIX,kaggle.com,Ghelper\n- DOMAIN-SUFFIX,arxiv.org,Ghelper\n- DOMAIN-SUFFIX,wiktionary.org,Ghelper\n- DOMAIN-SUFFIX,blogger.com,Ghelper\n- DOMAIN-SUFFIX,youtu.be,Ghelper\n- DOMAIN-SUFFIX,ytimg.com,Ghelper\n- DOMAIN-SUFFIX,youtube.com,Ghelper\n- DOMAIN-SUFFIX,instagram.com,Ghelper\n- DOMAIN-SUFFIX,twitter.com,Ghelper\n- DOMAIN-SUFFIX,t.co,Ghelper\n- DOMAIN-SUFFIX,facebook.com,Ghelper\n- DOMAIN-SUFFIX,telegram.org,Ghelper\n- IP-CIDR,192.168.0.0/16,DIRECT\n- IP-CIDR,10.0.0.0/8,DIRECT\n- IP-CIDR,172.16.0.0/12,DIRECT\n- IP-CIDR,172.23.0.0/12,DIRECT\n- IP-CIDR,127.0.0.0/8,DIRECT\n- IP-CIDR,100.64.0.0/10,DIRECT\n- IP-CIDR6,::1/128,DIRECT\n- IP-CIDR6,fc00::/7,DIRECT\n- IP-CIDR6,fe80::/10,DIRECT\n- IP-CIDR6,fd00::/8,DIRECT\n- GEOIP,CN,DIRECT\n- MATCH,DIRECT'
#     print(s)

if MODE == "quan x":
    for token in urls_proxy:
        print('host-suffix,{},proxy'.format(token))
    for token in urls_direct:
        print('host-suffix,{},direct'.format(token))
    # s = 'ip-cidr, 10.0.0.0/8, direct\nip-cidr, 127.0.0.0/8, direct\nip-cidr, 172.16.0.0/12, direct\nip-cidr, 192.168.0.0/16, direct\nip-cidr, 224.0.0.0/24, direct\ngeoip, cn, direct\nfinal, direct'
if MODE == "nekoray":
    pass