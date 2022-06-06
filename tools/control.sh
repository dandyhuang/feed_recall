#! /bin/bash

ROOT_DIR=$(cd "$(dirname "$0")"; cd ../; pwd)
pwd
cd ${ROOT_DIR}

config="./configs"
envflag="prd"
app_name=`awk -F'/' '{print $(NF-1)}' <<< ${ROOT_DIR}`


# 参数要大于2个
if [ $# -lt 3 ]; then
    echo "Usage: $0 {start|stop|restart|status} procname env{test|pre|prd}"
    exit
fi

if [ $# -ge 3 ];then
    envflag=$3
fi

procname=bin/$2
start_proc="${ROOT_DIR}/${procname}"
consul_service_name=""
service_port=""
register_type=""
status_dir="${ROOT_DIR}/supervise_status"
mk_status_dir=${status_dir}/bin
local_ip=`hostname -i`



_supervise_start() {
    if [ ! -d "$mk_status_dir" ]; then
        mkdir -p ${mk_status_dir}
    fi

    if test $( ps -ef | grep ${start_proc} | grep  supervisor | wc -l ) -eq 0
    then
        echo "supervisor进程未启动, 启动之前先杀掉已经存在的进程"
        PID=`ps -ef | grep -w ${start_proc} | grep -v grep |  grep -v $0 | awk '{print $2}'`
        if [ "${PID}x" != "x" ]; then
            for x in $PID
            do
                echo "${APP_NAME} process is alive, PID is:${PID} send 9 to kill"
                kill -9 ${PID}
            done
        fi

        chmod 755 ${ROOT_DIR}/tools/supervisor
        ${ROOT_DIR}/tools/supervisor -p $status_dir/${procname} -f "${ROOT_DIR}/${procname} \
             -conf ${config} -env ${envflag}" > ./super.log 2>&1
        sleep 3
        echo "=== ${ROOT_DIR}/tools/supervisor -p $status_dir/${procname} -f ${ROOT_DIR}/${procname}"
        if test $( ps -ef | grep ${procname} | grep  supervisor | wc -l ) -eq 0
        then
            echo "supervisor进程启动失败"
            exit 1
        else
            echo "supervisor进程启动成功"
        fi
    else
        echo "supervisor进程已启动"
    fi
}

_supervise_stop() {
    if test $( ps -ef | grep ${start_proc} | grep  supervisor | wc -l ) -eq 0
    then
        echo "supervisor进程不存在, 不需要stop"
    else
        echo "supervisor进程将停止"
        echo "d" > $status_dir/${procname}/control
        echo "x" > $status_dir/${procname}/control
        sleep 1
    fi
}

_crontab_stop() {
    if test $( crontab -l | grep -w ${procname} | wc -l ) -eq 0
    then
        echo "crontab监控进程不存在, 不需要卸载"
    else
        echo "crontab监控进程将卸载"
        crontab -l > ./tmp.cron
        sed -i  "/${procname}/d" ./tmp.cron
        crontab ./tmp.cron
        sleep 1
    fi
}

start()
{
    _supervise_start;
    if [ "${register_type}" == "consul" ];then
        sh tools/consul.sh register "${consul_service_name}" ${service_port}
        sleep 3
    fi

    if [ $? -gt 0 ];then
        echo "${start_proc} consul register failed"
        exit 1
    fi
}

stop()
{
    _supervise_stop

    PID=`ps -ef | grep -w ${start_proc} | grep -v grep |  grep -v $0 | awk '{print $2}'`
    if [ "${PID}x" != "x" ]; then
        for x in $PID
        do
            echo "${app_name} process is alive, PID is:${PID} send 9 to kill"
            kill -9 ${PID}
        done
    fi
}

status()
{
    count=`ps -ef | grep -w ${start_proc} | grep -v grep |  grep -v $0 | awk '{print $2}' | wc -l`
    if [ ${count} -eq 0 ];then
        return 1
    fi

    return 0
}
monitor()
{
    cn=`ps -ef | grep -w ${start_proc} | grep -v grep |  grep -v $0 | awk '{print $2}' | wc -l`
    if [ ${cn} -eq 0 ];then
        start
        echo "${start_proc}  no alive"
    else
        echo "${start_proc} alive"
    fi
}
# See how we were called.
case "$1" in
    start)
        start
        ;;
    stop)
        stop
        ;;
    restart)
        stop
        start
        ;;
    status)
        status
        ;;
    monitor)
      monitor
        ;;
       *)
        echo $"Usage: $0 {start|stop|restart|status} procname [config]"
        exit 2
esac

exit $?
