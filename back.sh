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

filename="$path/some.$date.sql"
echo $filename

docker exec mysql sh -c 'exec mysqldump -uroot -proot some' > $filename

# TODO 对备份文件进行压缩打包
ls -l $path
