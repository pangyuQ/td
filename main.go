package main

import (
	"fmt"
	"os"
	"strings"
)

func printHelp() {
	fmt.Println("个人待办 CLI 工具 (td)")
	fmt.Println("用法:")
	fmt.Println("  td add, a <内容>   - 添加新待办")
	fmt.Println("  td list, ls        - 列出未完成待办")
	fmt.Println("  td list -a, ls -a  - 列出所有待办 (包含已完成)")
	fmt.Println("  td do, d [id]      - 完成待办 (不加 ID 将进入交互式选择)")
	fmt.Println("  td rm [id]         - 删除待办 (不加 ID 将进入交互式选择)")
	fmt.Println("  td rm -a           - 删除所有已完成的待办事项")
	fmt.Println("  td help, h         - 显示此帮助信息")
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	command := strings.ToLower(os.Args[1])

	switch command {
	case "add", "a":
		content := strings.Join(os.Args[2:], " ")
		CmdAdd(content)
	case "list", "ls":
		showAll := false
		if len(os.Args) > 2 && os.Args[2] == "-a" {
			showAll = true
		}
		CmdList(showAll)
	case "do", "d":
		idStr := ""
		if len(os.Args) > 2 {
			idStr = os.Args[2]
		}
		CmdDo(idStr)
	case "rm":
		idStr := ""
		if len(os.Args) > 2 {
			idStr = os.Args[2]
		}
		CmdRm(idStr)
	case "help", "h":
		printHelp()
	default:
		fmt.Printf("未知命令: %s\n\n", command)
		printHelp()
	}
}
