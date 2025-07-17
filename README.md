# To-Do List REST API

A simple CRUD (Create, Read, Update, Delete) API for managing to-do items, built with Go, Gin, GORM and a pure-Go SQLite driver. No external database required—your data lives in a local `todos.db` file.

---

## 🚀 Features

- **CRUD endpoints** for to-do items  
- **Gin** for HTTP routing  
- **GORM** ORM with [glebarez/sqlite] driver (no cgo)  
- Auto-migrations: creates the `todos` table on first run  
- Environment-driven configuration via `.env`  

---

## 📦 Prerequisites

- Go ≥ 1.20  
- Git  

---

## 🔧 Installation & Setup

1. **Clone the repo**  
   ```bash
   git clone https://github.com/yourusername/todo-api.git
   cd todo-api

2. Download dependencies

    ```bash
    go mod tidy

## 🔗 API Endpoints

| Method | Path         | Description                  | Body (JSON)                                           |
|--------|--------------|------------------------------|-------------------------------------------------------|
| POST   | `/todos`     | Create a new to-do           | `{ "title": "Buy groceries", "completed": false }`    |
| GET    | `/todos`     | List all to-dos              | —                                                     |
| GET    | `/todos/:id` | Get a single to-do by ID     | —                                                     |
| PUT    | `/todos/:id` | Update an existing to-do     | `{ "title": "New title", "completed": true }`         |
| DELETE | `/todos/:id` | Delete a to-do               | —                                                     |

## 🧪 Testing with Postman

1. **Set up Environment**  
   - Create a Postman Environment named e.g. “Local”  
   - Add variable `baseUrl` with value `http://localhost:8080`

2. **Create Todo**  
   - Method: `POST`  
   - URL: `{{baseUrl}}/todos`  
   - Headers: `Content-Type: application/json`  
   - Body (raw JSON):
     ```json
     {
       "title": "Buy groceries",
       "completed": false
     }
     ```
   - Expected: `201 Created` and the created object

3. **List All Todos**  
   - Method: `GET`  
   - URL: `{{baseUrl}}/todos`  
   - Expected: `200 OK` and an array of todos

4. **Get Todo by ID**  
   - Method: `GET`  
   - URL: `{{baseUrl}}/todos/1`  
   - Expected: `200 OK` and the todo with `id=1`

5. **Update a Todo**  
   - Method: `PUT`  
   - URL: `{{baseUrl}}/todos/1`  
   - Headers: `Content-Type: application/json`  
   - Body (raw JSON):
     ```json
     {
       "title": "Buy almond milk",
       "completed": true
     }
     ```
   - Expected: `200 OK` and the updated object

6. **Delete a Todo**  
   - Method: `DELETE`  
   - URL: `{{baseUrl}}/todos/1`  
   - Expected: `204 No Content`