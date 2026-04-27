package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CmdAdd 添加新的待办事项
func CmdAdd(content string) {
	if strings.TrimSpace(content) == "" {
		fmt.Println("待办内容不能为空。使用方法: td add <内容>")
		return
	}

	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("加载待办事项失败:", err)
		return
	}

	newID := 1
	for _, t := range tasks {
		if t.ID >= newID {
			newID = t.ID + 1
		}
	}

	task := Task{
		ID:          newID,
		Content:     content,
		IsCompleted: false,
	}

	tasks = append(tasks, task)
	if err := SaveTasks(tasks); err != nil {
		fmt.Println("保存待办事项失败:", err)
		return
	}

	fmt.Printf("成功添加待办: [%d] %s\n", task.ID, task.Content)
}

// CmdList 列出待办事项
func CmdList(showAll bool) {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("加载待办事项失败:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("当前没有任何待办事项。")
		return
	}

	count := 0
	for _, t := range tasks {
		if !showAll && t.IsCompleted {
			continue
		}
		count++
		status := " "
		if t.IsCompleted {
			status = "x"
		}
		fmt.Printf("[%s] %d. %s\n", status, t.ID, t.Content)
	}

	if count == 0 && !showAll {
		fmt.Println("当前没有任何未完成的待办事项。使用 'td ls -a' 查看所有事项。")
	}
}

// getTaskIDFromInput 交互式获取用户输入的任务 ID
func getTaskIDFromInput(prompt string) int {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	id, err := strconv.Atoi(input)
	if err != nil {
		return -1
	}
	return id
}

// CmdDo 完成待办事项
func CmdDo(idStr string) {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("加载待办事项失败:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("当前没有任何待办事项。")
		return
	}

	var id int
	if idStr == "" {
		// 打印未完成的列表供用户参考
		fmt.Println("未完成的待办事项:")
		hasUncompleted := false
		for _, t := range tasks {
			if !t.IsCompleted {
				fmt.Printf("[ ] %d. %s\n", t.ID, t.Content)
				hasUncompleted = true
			}
		}
		if !hasUncompleted {
			fmt.Println("所有待办均已完成！")
			return
		}
		
		id = getTaskIDFromInput("请输入要完成的待办 ID: ")
		if id == -1 {
			fmt.Println("无效的 ID。")
			return
		}
	} else {
		parsedID, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("无效的 ID，ID 必须是数字。")
			return
		}
		id = parsedID
	}

	found := false
	for i, t := range tasks {
		if t.ID == id {
			if t.IsCompleted {
				fmt.Printf("待办事项 '%s' 已经处于完成状态。\n", t.Content)
				return
			}
			tasks[i].IsCompleted = true
			found = true
			fmt.Printf("已完成待办: %s\n", t.Content)
			break
		}
	}

	if !found {
		fmt.Printf("未找到 ID 为 %d 的待办事项。\n", id)
		return
	}

	if err := SaveTasks(tasks); err != nil {
		fmt.Println("保存待办事项失败:", err)
	}
}

// CmdRm 删除待办事项
func CmdRm(idStr string) {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("加载待办事项失败:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("当前没有任何待办事项。")
		return
	}

	var id int
	if idStr == "" {
		// 打印所有列表供用户参考
		CmdList(true)
		id = getTaskIDFromInput("请输入要删除的待办 ID: ")
		if id == -1 {
			fmt.Println("无效的 ID。")
			return
		}
	} else {
		parsedID, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("无效的 ID，ID 必须是数字。")
			return
		}
		id = parsedID
	}

	found := false
	newTasks := make([]Task, 0, len(tasks)-1)
	for _, t := range tasks {
		if t.ID == id {
			found = true
			fmt.Printf("已删除待办: %s\n", t.Content)
		} else {
			newTasks = append(newTasks, t)
		}
	}

	if !found {
		fmt.Printf("未找到 ID 为 %d 的待办事项。\n", id)
		return
	}

	if err := SaveTasks(newTasks); err != nil {
		fmt.Println("保存待办事项失败:", err)
	}
}
