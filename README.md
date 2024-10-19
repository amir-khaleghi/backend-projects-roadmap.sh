# Go Backend Projects Collection

This repository is a comprehensive collection of backend projects built using **Go (Golang)**, inspired by the [Backend Developer Roadmap](https://roadmap.sh/backend/projects). The purpose of this repository is to guide developers through essential backend concepts by solving real-world challenges. From building APIs to handling microservices, each project is designed to improve proficiency in Go and introduce you to various aspects of backend development, including:

- Designing and implementing RESTful APIs
- Managing databases and data flow
- Securing applications with authentication and authorization
- Working with caching systems like Redis
- Building and scaling microservices

## 1. Project: [Task Tracker](https://github.com/amir-khaleghi/backend-projects-roadmap.sh/tree/main/1-task-tracker)

[Project URL](https://roadmap.sh/projects/task-tracker)

The **Task Tracker** project is a command-line interface (CLI) application used to track and manage tasks in your to-do list. This project helps you practice essential programming skills like handling user input, working with the filesystem, and building a robust CLI app.

### Key Features:

- **Track Tasks**: Add, update, and delete tasks.
- **Status Management**: Keep track of what you need to do, what you're working on, and what you've completed.
- **Task Timer**: Option to start and stop a timer to measure how long you've worked on a specific task.

### Technologies:

- **Go (Golang)**: The main programming language used to build the CLI.
- **Filesystem Management**: For storing and retrieving task data.
- **User Input Handling**: Command-line arguments and user prompts for managing tasks.

### How to Get Started with the Task Tracker:

1. **Clone the repository**:

   ```bash
   git clone https://github.com/your-username/go-backend-projects.git
   ```

2. **Navigate to the Task Tracker project**:

   ```bash
   cd go-backend-projects/task-tracker
   ```

3. **Install dependencies**:

   ```bash
   go mod tidy
   ```

4. **Run the Task Tracker CLI**:

   ```bash
   go run main.go
   ```

5. **Testing** (if applicable):
   You can run the tests using:
   ```bash
   go test ./...
   ```

## Technologies Used

Each project is implemented using **Go (Golang)** as the primary programming language. For the **Task Tracker** project, key technologies include:

- **Filesystem Operations**: For task storage.
- **Go’s Flag Package**: To handle CLI arguments.
- **Time Management**: To track how long tasks are worked on.

## Contributing

Contributions are welcome! If you'd like to add more projects, improve documentation, or fix bugs, feel free to submit a pull request. Please follow the contribution guidelines outlined in the `CONTRIBUTING.md` file.

## License

This repository is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.
