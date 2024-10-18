package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

/* ■■■■■■■■■■■■■■■■■■■■■■■■ Types ■■■■■■■■■■■■■■■■■■■■■■■ */

// Task represents a single task in the tracker
type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TaskTracker manages the task collection and file operations
type TaskTracker struct {
	Tasks    []Task `json:"tasks"`
	Filename string
	Scanner  *bufio.Scanner
}

/* ■■■■■■■■■■■■■■■■■■■■■■■■ Main ■■■■■■■■■■■■■■■■■■■■■■■ */
func main() {
	tracker := NewTaskTracker()
	if err := tracker.LoadTasks(); err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}

	fmt.Println("Welcome to Task Tracker!")

	for {
		showMenu()
		tracker.Scanner.Scan()
		choice := tracker.Scanner.Text()

		var err error
		switch choice {
		case "1":
			err = tracker.AddTask()
		case "2":
			err = tracker.UpdateTask()
		case "3":
			err = tracker.DeleteTask()
		case "4":
			err = tracker.MarkTask()
		case "5":
			tracker.ListTasks()
		case "6":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
			continue
		}

		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

/* New Task --------------------------------------------- */
// NewTaskTracker creates a new instance of TaskTracker
func NewTaskTracker() *TaskTracker {
	return &TaskTracker{
		Tasks:    make([]Task, 0),
		Filename: "tasks.json",
		Scanner:  bufio.NewScanner(os.Stdin),
	}
}

/* Load TASKS ------------------------------------------- */
// LoadTasks reads tasks from the JSON file
func (t *TaskTracker) LoadTasks() error {
	if _, err := os.Stat(t.Filename); os.IsNotExist(err) {
		return nil
	}

	data, err := os.ReadFile(t.Filename)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	if len(data) == 0 {
		return nil
	}

	if err := json.Unmarshal(data, &t.Tasks); err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}

	return nil
}

/* SAVE TASK -------------------------------------------- */
// SaveTasks writes tasks to the JSON file
func (t *TaskTracker) SaveTasks() error {
	data, err := json.MarshalIndent(t.Tasks, "", "    ")
	if err != nil {
		return fmt.Errorf("error creating JSON: %v", err)
	}

	err = os.WriteFile(t.Filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	return nil
}

/* GET NEXT ID ------------------------------------------ */

// GetNextID generates the next available task ID
func (t *TaskTracker) GetNextID() int {
	maxID := 0
	for _, task := range t.Tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1
}

/* ADD TASK --------------------------------------------- */
// AddTask creates a new task
func (t *TaskTracker) AddTask() error {
	fmt.Print("Enter task description: ")
	t.Scanner.Scan()
	description := t.Scanner.Text()

	if description == "" {
		return fmt.Errorf("task description cannot be empty")
	}

	task := Task{
		ID:          t.GetNextID(),
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.Tasks = append(t.Tasks, task)
	if err := t.SaveTasks(); err != nil {
		return err
	}

	fmt.Printf("Task added successfully (ID: %d)\n", task.ID)
	return nil
}

/* FIND TASK -------------------------------------------- */
// FindTask locates a task by ID
func (t *TaskTracker) FindTask(id int) (*Task, int) {
	for i, task := range t.Tasks {
		if task.ID == id {
			return &t.Tasks[i], i
		}
	}
	return nil, -1
}

/* UPDATE TASK ------------------------------------------ */
// UpdateTask modifies an existing task's description
func (t *TaskTracker) UpdateTask() error {
	fmt.Print("Enter task ID to update: ")
	t.Scanner.Scan()
	idStr := t.Scanner.Text()

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("invalid task ID: %v", err)
	}

	task, _ := t.FindTask(id)
	if task == nil {
		return fmt.Errorf("task %d not found", id)
	}

	fmt.Printf("Current description: %s\n", task.Description)
	fmt.Print("Enter new description: ")
	t.Scanner.Scan()
	description := t.Scanner.Text()

	if description == "" {
		return fmt.Errorf("task description cannot be empty")
	}

	task.Description = description
	task.UpdatedAt = time.Now()

	if err := t.SaveTasks(); err != nil {
		return err
	}

	fmt.Printf("Task %d updated successfully\n", id)
	return nil
}

/* DELETE TASK ------------------------------------------ */

// DeleteTask removes a task
func (t *TaskTracker) DeleteTask() error {
	fmt.Print("Enter task ID to delete: ")
	t.Scanner.Scan()
	idStr := t.Scanner.Text()

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("invalid task ID: %v", err)
	}

	task, index := t.FindTask(id)
	if task == nil {
		return fmt.Errorf("task %d not found", id)
	}

	fmt.Printf("Are you sure you want to delete task '%s'? (y/n): ", task.Description)
	t.Scanner.Scan()
	confirm := strings.ToLower(t.Scanner.Text())

	if confirm != "y" && confirm != "yes" {
		return fmt.Errorf("deletion cancelled")
	}

	t.Tasks = append(t.Tasks[:index], t.Tasks[index+1:]...)

	if err := t.SaveTasks(); err != nil {
		return err
	}

	fmt.Printf("Task %d deleted successfully\n", id)
	return nil
}

/* TASK STATUS ------------------------------------------ */

// MarkTask updates a task's status
func (t *TaskTracker) MarkTask() error {
	fmt.Print("Enter task ID to mark: ")
	t.Scanner.Scan()
	idStr := t.Scanner.Text()

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("invalid task ID: %v", err)
	}

	task, _ := t.FindTask(id)
	if task == nil {
		return fmt.Errorf("task %d not found", id)
	}

	fmt.Println("\nAvailable statuses:")
	fmt.Println("1. todo")
	fmt.Println("2. in-progress")
	fmt.Println("3. done")
	fmt.Print("Enter status number: ")

	t.Scanner.Scan()
	statusNum := t.Scanner.Text()

	var status string
	switch statusNum {
	case "1":
		status = "todo"
	case "2":
		status = "in-progress"
	case "3":
		status = "done"
	default:
		return fmt.Errorf("invalid status number")
	}

	task.Status = status
	task.UpdatedAt = time.Now()

	if err := t.SaveTasks(); err != nil {
		return err
	}

	fmt.Printf("Task %d marked as %s\n", id, status)
	return nil
}

/* LIST TASKS ------------------------------------------- */

// ListTasks displays tasks, optionally filtered by status
func (t *TaskTracker) ListTasks() {
	fmt.Println("\nList options:")
	fmt.Println("1. All tasks")
	fmt.Println("2. Todo tasks")
	fmt.Println("3. In-progress tasks")
	fmt.Println("4. Done tasks")
	fmt.Print("Enter option number: ")

	t.Scanner.Scan()
	option := t.Scanner.Text()

	var status string
	switch option {
	case "1":
		status = ""
	case "2":
		status = "todo"
	case "3":
		status = "in-progress"
	case "4":
		status = "done"
	default:
		fmt.Println("Invalid option. Showing all tasks.")
		status = ""
	}

	if len(t.Tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	fmt.Println("\nID | Status | Description | Created At | Updated At")
	fmt.Println(strings.Repeat("-", 80))

	for _, task := range t.Tasks {
		if status == "" || task.Status == status {
			fmt.Printf("%d | %s | %s | %s | %s\n",
				task.ID,
				task.Status,
				task.Description,
				task.CreatedAt.Format("2006-01-02 15:04"),
				task.UpdatedAt.Format("2006-01-02 15:04"))
		}
	}
}

func showMenu() {
	fmt.Println("\nTask Tracker Menu:")
	fmt.Println("1. Add new task")
	fmt.Println("2. Update task")
	fmt.Println("3. Delete task")
	fmt.Println("4. Mark task status")
	fmt.Println("5. List tasks")
	fmt.Println("6. Exit")
	fmt.Print("Enter your choice (1-6): ")
}
