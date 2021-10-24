#!/usr/bin/bash
# 每月定时，备份整个数据库， 文件格式: some.2109.sql
# 1 23 15 * * /usr/bin/bash /home/gelu/go/src/some/back.sh

year=`date +%Y`
date=`date +%y%m`
echo $date

path="/data/back/db/$year"

if [ ! -d $path ]; then
    mkdir -p $path
fi

ddl="$path/some.$date.sql"
dml="$path/some.$date.data.sql"
echo $filename

# 导出表结构
docker exec mysql mysqldump -uroot -proot -d some > $ddl

# 导出数据
docker exec mysql mysqldump -uroot -proot -t some > $dml

# TODO 对备份文件进行压缩打包
ls -l $path
