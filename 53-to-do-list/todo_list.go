package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Default Redis port
	})

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nTodo List Menu:")
		fmt.Println("1. Add a task")
		fmt.Println("2. View tasks")
		fmt.Println("3. Complete (remove) a task")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			for {
				fmt.Print("Enter a task (blank to stop): ")
				task, _ := reader.ReadString('\n')
				task = strings.TrimSpace(task)
				if task == "" {
					break
				}
				rdb.RPush(ctx, "todo:tasks", task)
			}
		case "2":
			tasks, err := rdb.LRange(ctx, "todo:tasks", 0, -1).Result()
			if err != nil {
				fmt.Println("Error fetching tasks:", err)
				continue
			}
			if len(tasks) == 0 {
				fmt.Println("No tasks found.")
			} else {
				fmt.Println("\nYour Tasks:")
				for i, task := range tasks {
					fmt.Printf("%d. %s\n", i+1, task)
				}
			}
		case "3":
			tasks, _ := rdb.LRange(ctx, "todo:tasks", 0, -1).Result()
			if len(tasks) == 0 {
				fmt.Println("No tasks to remove.")
				continue
			}
			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i+1, task)
			}
			fmt.Print("Enter task number to mark as completed: ")
			indexStr, _ := reader.ReadString('\n')
			indexStr = strings.TrimSpace(indexStr)
			index, err := strconv.Atoi(indexStr)
			if err != nil || index < 1 || index > len(tasks) {
				fmt.Println("Invalid task number.")
				continue
			}
			taskToRemove := tasks[index-1]
			rdb.LRem(ctx, "todo:tasks", 1, taskToRemove)
			fmt.Println("Task removed.")
		case "4":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice.")
		}
		time.Sleep(1 * time.Second)
	}
}
