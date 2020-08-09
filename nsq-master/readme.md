第一步、nsqlookupd.exe
第二步、nsqd.exe --lookupd-tcp-address=127.0.0.1:4160
第三步、nsqadmin.exe --lookupd-http-address=127.0.0.1:4161
第四步、浏览器打开127.0.0.1:4171进入nsq管理界面
