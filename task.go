package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Task 表示单个待办事项
type Task struct {
	ID          int    `json:"id"`
	Content     string `json:"content"`
	IsCompleted bool   `json:"is_completed"`
}

// getStoragePath 返回 JSON 存储文件的路径
func getStoragePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		// 如果无法获取用户主目录，退回到当前目录
		return ".td.json"
	}
	return filepath.Join(homeDir, ".td.json")
}

// LoadTasks 从文件中加载待办事项列表
func LoadTasks() ([]Task, error) {
	path := getStoragePath()
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

// SaveTasks 将待办事项列表保存到文件
func SaveTasks(tasks []Task) error {
	path := getStoragePath()
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
