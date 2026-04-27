#!/bin/bash
go build -o td

echo "=== 1. 初始化数据 ==="
rm -f ~/.td.json
./td a 买牛奶
./td a 跑步
./td a 看书
./td do 2

echo -e "\n=== 2. 测试 td list (只显示未完成) ==="
./td ls

echo -e "\n=== 3. 测试 td list -a (显示所有) ==="
./td ls -a

echo -e "\n=== 4. 测试全部完成后 td list 的表现 ==="
./td d 1
./td d 3
./td ls

echo -e "\n=== 5. 测试全部完成后 td list -a 的表现 ==="
./td ls -a
