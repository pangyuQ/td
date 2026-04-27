#!/bin/bash
echo "=== 1. 初始状态 (清空数据) ==="
rm -f ~/.td.json
./td ls

echo -e "\n=== 2. 帮助命令 ==="
./td
./td h

echo -e "\n=== 3. 添加待办 ==="
./td a 买牛奶
./td add 跑步5公里
./td a 看书一小时

echo -e "\n=== 4. 查看待办列表 ==="
./td ls

echo -e "\n=== 5. 完成待办 (直接指定 ID) ==="
./td d 1

echo -e "\n=== 6. 完成待办 (重复完成) ==="
./td d 1

echo -e "\n=== 7. 完成待办 (无效 ID) ==="
./td d 999

echo -e "\n=== 8. 完成待办 (交互式输入 ID=2) ==="
echo "2" | ./td do

echo -e "\n=== 9. 查看待办列表 (1和2应为完成状态) ==="
./td ls

echo -e "\n=== 10. 删除待办 (直接指定 ID) ==="
./td rm 1

echo -e "\n=== 11. 删除待办 (交互式输入 ID=3) ==="
echo "3" | ./td rm

echo -e "\n=== 12. 查看待办列表 (应只剩已完成的'跑步5公里') ==="
./td ls

echo -e "\n=== 13. 异常测试 (空内容添加) ==="
./td a
