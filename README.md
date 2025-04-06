# 🛠️ taskcli

**taskcli** is a command-line application for managing projects and tasks, written in Go using the cobra-cli tool
It supports full CRUD operations for projects and tasks which are assigned to any project.  
The app uses PostgreSQL or MySQL for persistence. I build this aplication using design patterns as much as i can.

---

## Requirements

- Go 1.20+
- Database (PostgreSQL or MySQL)

---

## Installation

1. Clone the repository:

```bash
git clone https://github.com/ajezierski/taskcli.git
cd taskcli
```
2. Create a .env file from the example:
```
cp .env.example .env
```
3. Set up your database (example for PostgreSQL):
```
createdb taskcli
```
4. Run migrations
```
postgres://postgres:password@localhost:5432/taskcli?sslmode=disable
```
5. Build the application
go build -o taskcli .


## Features
* CRUD for projects

* CRUD for tasks

* Filter tasks by project

* Marking task in table ( DONE / TODO )

* Support for PostgreSQL and MySQL


## Example usage

* Projects
```
./taskcli project create "Backend API"
./taskcli project list
./taskcli project get <project_id>
./taskcli project update <project_id> "Renamed project"
./taskcli project delete <project_id>
```

* Tasks
```
./taskcli task add <project_id> "Implement login"
./taskcli task list
./taskcli task list --project <project_id>
./taskcli task done <task_id>
./taskcli task delete <task_id>
```
