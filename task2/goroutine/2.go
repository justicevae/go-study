package main

import (
	"fmt"
	"sync"
	"time"
)

type Task func()

type TaskResult struct {
	TaskID   int
	Duration time.Duration
}

type TaskScheduler struct {
	tasks   []Task
	results []TaskResult
	wg      sync.WaitGroup
}

func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{
		tasks:   make([]Task, 0),
		results: make([]TaskResult, 0),
	}
}

func (ts *TaskScheduler) AddTask(task Task) {
	ts.tasks = append(ts.tasks, task)
}

func (ts *TaskScheduler) Run() {
	ts.results = make([]TaskResult, len(ts.tasks))
	ts.wg.Add(len(ts.tasks))

	for i, task := range ts.tasks {
		go func(taskID int, t Task) {
			defer ts.wg.Done()

			start := time.Now()
			t()
			end := time.Now()

			ts.results[taskID] = TaskResult{
				TaskID:   taskID,
				Duration: end.Sub(start),
			}
		}(i, task)
	}

	ts.wg.Wait()
}

func (ts *TaskScheduler) Print() {
	for _, result := range ts.results {
		fmt.Printf("任务 %d 耗时: %v\n", result.TaskID+1, result.Duration)
	}
}

func main() {
	scheduler := NewTaskScheduler()
	scheduler.AddTask(func() {})
	scheduler.AddTask(func() {})
	scheduler.Run()
	scheduler.Print()
}
