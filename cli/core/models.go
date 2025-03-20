package core

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type Task struct {
	DateCreated time.Time `json:"date_created"`
	Description string    `json:"description"`
}

type TaskCollection struct {
	Tasks map[int]Task `json:"tasks"`
}

func LoadTaskCollectionFromFile(ctx context.Context, tasksFile string) (*TaskCollection, error) {
	file, err := os.Open(tasksFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var tasks TaskCollection
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return &tasks, nil
}

func (tc *TaskCollection) Add(taskDescription string) error {
	taskId := len(tc.Tasks) + 1
	for {
		if _, ok := tc.Tasks[taskId]; ok {
			break
		}
		taskId++
	}
	newTask := Task{
		DateCreated: time.Now(),
		Description: taskDescription,
	}
	tc.Tasks[taskId] = newTask
	return nil
}

func (tc *TaskCollection) Save(tasksFile string) error {
	data, err := json.MarshalIndent(tc, "", "\t")
	if err != nil {
		return fmt.Errorf("Failed to write: %w", err)
	}
	file, err := os.Create(tasksFile)
	if err != nil {
		return fmt.Errorf("Failed to write: %w", err)
	}
	defer file.Close()
	_, err = file.Write(data)
	return err
}

func (tc *TaskCollection) List() {
	fmt.Println("JOCASTA TASKS")
	fmt.Println("-------------")
	if len(tc.Tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for taskId, task := range tc.Tasks {
		fmt.Printf("%d. %s (%s)\n", taskId, task.Description, task.DateCreated.Format(time.DateTime))
	}
}
