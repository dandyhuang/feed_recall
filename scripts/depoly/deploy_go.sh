#!/bin/bash
rm -rf data_proxy.tar.gz
tar zcvf data_proxy.tar.gz configs tools bin/${1} Makefile
echo "应用名：$1"
echo "服务器：$2"
# ssh ${2} "cd /home/www/server/${1} && mv bin.tar.gz binbak.tar.gz"
#scp bin.tar.gz ${2}:/home/www/server/${1}/

#ssh ${2} "cd /home/www/server/${1} && \
#	    tar xvf bin.tar.gz && \
#	    chown www:www -R /home/www/server/${1} && \
#	    supervisorctl restart ${1}"
