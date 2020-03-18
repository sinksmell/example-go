# !/bin/bash
docker run -v /etc/localtime:/etc/localtime:ro -e NOTICE_TYPE=mail \ 
-e NOTICE_MAIL_TO=2448016708@qq.com,sinksmell@163.com \ 
-e NOTICE_MAIL_PWD=liicblmxribhdiab \ 
-e NOTICE_MAIL_PORT=25 \  
-e NOTICE_MAIL_HOST=smtp.qq.com \ 
-e NOTICE_MAIL_EMAIL=2448016708@qq.com \ 
-e NOTICE_MAIL_CC=2448016708@qq.com,sinksmell@163.com 