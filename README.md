# 📝 Go To-Do CLI App

This is a simple **Command-Line To-Do App** built with Go. It allows you to:

- ✅ Add tasks
- 📋 List tasks
- 🟢 Mark tasks as done
- 💾 Save tasks to a JSON file for persistence

---

## 🚀 **Setup and Installation**

1. **Clone the Repository:**
```bash
git clone https://github.com/zbinx/todo-cli-go.git
cd todo-cli-go
```

2. **Ensure Go is Installed:**
```bash
go version
```
If Go is not installed, download it from [golang.org](https://golang.org/dl/).

3. **Initialize Go Module:**
```bash
go mod init todo-cli
```

4. **Build the App (Optional):**
```bash
go build -o todo
```

---

## 📚 **Usage**

Run the app using `go run main.go` followed by a command.

### **1. Add a Task:**
```bash
go run main.go add "Buy groceries"
```
**Output:**
```
Task added: Buy groceries
```

### **2. List All Tasks:**
```bash
go run main.go list
```
**Output:**
```
Task List:
1. [❌] Buy groceries
```

### **3. Mark a Task as Done:**
```bash
go run main.go done 1
```
**Output:**
```
Task 1 marked as done: Buy groceries
```

### **4. View Updated Task List:**
```bash
go run main.go list
```
**Output:**
```
Task List:
1. [✅] Buy groceries
```

---

## 💾 **Data Storage**
- Tasks are stored in a `taskList.json` file in the project folder.
- The app loads tasks from this file when it starts and saves updates automatically.

**Example `taskList.json` file:**
```json
[
  {
    "id": 1,
    "content": "Buy groceries",
    "done": true
  }
]
```

---

## ⚙️ **Project Structure**
```
├── main.go        // Main program file
├── tasks.json     // JSON file storing tasks
├── go.mod         // Go module file
└── README.md      // Project documentation
```

---

## 🛠️ **Future Improvements**
- 🗑️ Delete tasks by ID
- 🖊️ Edit task descriptions
- 📅 Set task deadlines
- 🔍 Filter tasks by status

---

## 🤝 **Contributing**
1. Fork the repository.
2. Create a new branch: `git checkout -b feature-branch`.
3. Make your changes and commit: `git commit -m "Add new feature"`.
4. Push to your fork: `git push origin feature-branch`.
5. Open a pull request!

---

## 📜 **License**
This project is licensed under the [MIT License](LICENSE).

---

## 🌟 **Support**
If you find this project helpful, please ⭐ star the repo and share it with others!

---

## 📧 **Contact**
For questions or suggestions, reach out.

---

🎉 *Happy coding!* 🎉

