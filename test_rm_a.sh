#!/bin/bash
go build -o td

echo "=== 1. 初始化数据 ==="
rm -f ~/.td.json
./td a 任务1
./td a 任务2
./td a 任务3
./td a 任务4

echo -e "\n=== 2. 完成任务 1 和 3 ==="
./td d 1
./td d 3

echo -e "\n=== 3. 删除所有已完成任务 ==="
./td rm -a

echo -e "\n=== 4. 再次删除已完成任务 (期望提示没有可删) ==="
./td rm -a

echo -e "\n=== 5. 检查列表 (期望只剩下任务 2 和 4) ==="
./td ls -a

