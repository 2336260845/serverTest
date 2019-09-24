#!/bin/bash

ServerName=myServerTest
ServerPort=3333
MaxCount=10

check()
{
    if (($(ps -ef | grep -w ${ServerName}| grep -v grep|wc -l) == 0));then
        # stopped
        return 1;
    else
        # running
        return 0;
    fi
}

start()
{
    check
    if (($? == 1));then
        echo  "standard server port:${ServerPort} to start......"
        # 启动服务
        nohup ./${ServerName} > stdout.log 2>&1 &
        count=1
        while true
        do
            if [[ ${count} -gt ${MaxCount} ]]; then
                echo "最大尝试次数为" ${MaxCount}
                echo "服务启动失败"
                break
            fi
            check
            if (($? == 1));then
                echo '启动失败,等待3s后重试...'
                sleep 3
            else
                echo -e '\033[32mstarted\033[1m\033[0m'
                break
            fi
            count=${count}+1
        done
    else
        echo "standard server port:${ServerPort} has been running!"
    fi
}

fstop()
{
    check
    if (($? == 1));then
        echo "standard server ${ServerPort} has been stopped!"
    else
        echo "standard server ${ServerPort} force to stop....."
        ps -ef | grep -w ${ServerName} |  grep -v grep | awk '{print $2}'| xargs kill -9
        count=1
        while true
        do
            if [[ ${count} -gt ${MaxCount} ]]; then
                echo "最大尝试停止服务次数" ${MaxCount}
                echo "服务停止失败"
                break
            fi
            check
            if (($? == 1));then
                echo -e '\033[32mstopped\033[1m\033[0m'
                break
            else
                echo '服务停止失败,等待3s后重试...'
                sleep 3
            fi
            count=${count}+1
        done
    fi
}

case ${1} in
    start)
        start
        ;;
    stop)
        fstop
        ;;
    *)
          echo "${0} <start|stop>"
          exit 1
          ;;
esac

