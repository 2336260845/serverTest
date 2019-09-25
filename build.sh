#!/bin/bash

function getcommit() {
    branch=`git rev-parse --abbrev-ref HEAD`
    commitid=`git rev-parse HEAD`
    echo $branch-$commitid
}

TARGET=myServerTest
version=`git describe --tags --always --dirty | tr '-' '.'`
commit=`getcommit`

BuildType=$1

echo "build" ${myServerTest}

go build -ldflags "-X main.versionStr=$version -X main.commitStr=$commit"  -o ${TARGET}

# 清空bin目录下所有文件,会将日志保留
mkdir -p ./log
mkdir -p ./bin

current=`date "+%Y-%m-%d-%H:%M:%S"`

if [[ -f "./bin/stdout.log" ]]; then
mv bin/stdout.log  ./log/stdout${current}.log
else
echo "没有日志文件,不需要保留"
fi

mv ${TARGET} ./bin

chmod +x start.sh
cp start.sh bin
cp config.toml bin
