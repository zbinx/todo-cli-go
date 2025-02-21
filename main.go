package main

import (
    "fmt"
    "os"
    "strconv"
    "encoding/json"
)

type Task struct {
    ID int `json: "id"`
    Content string `json:"content"`
    Done bool `json:"done"`
}

const taskFile = "taskFile.json"

var taskList []Task

func main() {

    loadTaskFile()

    if len(os.Args) < 2 {
        fmt.Println("Usage: todo <command> [arguments]")
        return
    }

    command := os.Args[1]

    switch command {
    case "add":
        addTask()

    case "list":
        listTask()

    case "done":
        markTaskAsDone()

    default:
        fmt.Println("Unknown command. Available commands: add, list, done")

    }

}

func loadTaskFile() {
    file, err := os.Open(taskFile)
    if err != nil {
        if os.IsNotExist(err) {
            return
        }
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    if err := decoder.Decode(&taskList); err != nil {
        fmt.Println("Error decoding JSON:", err)
    }
}

func saveTasks() {
    file, err := os.Create(taskFile)
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    if err := encoder.Encode(taskList); err != nil {
        fmt.Println("Error writing to file:", err)
    } else {
        fmt.Println("Tasks successfully saved to", taskFile)
    }
}

func addTask() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: todo <command> [arguments]")
        return
    }

    taskContent := os.Args[2]

    newTask := Task{
        ID: len(taskList) + 1,
        Content: taskContent,
        Done: false, 	
    }

    taskList = append(taskList, newTask)
    saveTasks()

    fmt.Printf("Task added: %s\n", newTask.Content)
}

func listTask() {
    if len(taskList) == 0 {
        fmt.Println("No tasks available.")
        return
    }

    fmt.Println("Task list:")
    for _, task := range taskList {
        status := "❌"
        if task.Done {
            status = "✅"
        }
        fmt.Printf("%d, [%s], %s\n", task.ID, status, task.Content)
    }
}

func markTaskAsDone() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: todo done <task id>")
        return
    }

    taskID, err := strconv.Atoi(os.Args[2])
    if err != nil {
        fmt.Println("Invalid task ID. Pleaes enter a valid number")
        return
    }

    updated := false
    for i, task := range taskList {
        if task.ID == taskID {
            if taskList[i].Done {
                fmt.Printf("Task %d marked as done: %s\n", taskID, task.Content)
                return
            }
            taskList[i].Done = true
            updated = true
            break
        }
    }

    if updated {
        saveTasks()
        fmt.Printf("Task %d marked as done.\n", taskID)
    } else {
        fmt.Printf("Task with ID %d not found.\n", taskID)
    }
}	
