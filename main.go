package main

import (
	"encoding/json"
	"flag"
	"fmt"

	// "log"
	"os"
)

type Task struct {
	Name   string `json:"name"`
	Id     int    `json:"id"`
	Status int    `json:"status"`
}

/*

    Add tasks - [x]
	Update tasks - [x]
	Delete tasks - [x]
    Mark a task as in progress or done - [x]
    List all tasks -[x]
    List all tasks that are done - [x]
    List all tasks that are not done - [x]
    List all tasks that are in progress - [x]

*/

func main() {
	add := flag.String("add", "", "add a new task")
	status := flag.Int("status", 0, "in-pogress = 0 ,done = 1 ")
	update := flag.String("update", "", "update a task name")
	id := flag.Int("id", 0, "task id")
	delete := flag.String("delete", "", "delete a task -> pass the id flag")
	in_progress := flag.Bool("mark-in-progress", true, "mark a task as in-progress")
	done := flag.Bool("mark-done", false, "mark a task as complete")
	list_all := flag.Bool("list-all", false, "list all tasks")
	list_done := flag.Bool("list-done", false, "list all done tasks")
	list_inprogress := flag.Bool("list-in-progress", false, "list all tasks in progress")

	flag.Parse()
	tasks_file, _ := os.ReadFile("tasks.json")
	file, _ := os.Create("tasks.json")

	defer file.Close()
	var tasks []Task

	json.Unmarshal(tasks_file, &tasks)

	if *add != "" {
		task := Task{Name: *add, Id: len(tasks) + 1, Status: *status}
		tasks = append(tasks, task)
	}
	if *update != "" && *id > 0 {
		tasks[*id-1].Name = *update
	}
	if *delete == "" && *id > 0 {
		for i, task := range tasks {
			if task.Id == *id {
				tasks = append(tasks[:i], tasks[i+1:]...)
			}
		}
	}
	if *in_progress && *id > 0 {
		tasks[*id-1].Status = 0
	}
	if *done && *id > 0 {
		tasks[*id-1].Status = 1
	}

	if *list_all {
		fmt.Println(tasks)
	}
	if *list_done {
		for _, task := range tasks {
			if task.Status == 1 {
				fmt.Println(task)
			}
		}
	}
	if *list_inprogress {
		for _, task := range tasks {
			if task.Status == 0 {
				fmt.Println(task)
			}
		}
	}
	as_json, _ := json.MarshalIndent(tasks, "", "\t")
	file.Write(as_json)

}
