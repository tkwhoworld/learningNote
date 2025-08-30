#!/bin/bash

BASE_DIR=`pwd`
if [ ! -d "${BASE_DIR}/logs" ]; then
  mkdir ${BASE_DIR}/logs
fi

logfile="${BASE_DIR}/logs/startup.log"
if [ ! -f "$logfile" ]; then
  touch "$logfile"
fi

nohup java -Dserver.port=8999 -Dcsp.sentinel.dashboard.server=localhost:8999 -Dproject.name=sentinel-dashboard -jar sentinel-dashboard-1.8.8.jar >> "$logfile" 2>&1 &

echo "nacos is starting. you can check the ${logfile}"