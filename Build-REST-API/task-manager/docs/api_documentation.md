package docs


````
# 📘 Task Manager API Documentation



This API allows basic CRUD operations for managing tasks. Built using **Go** and the **Gin Framework**.

---

## 📌 Base URL

````

[http://localhost:8080](http://localhost:8080)

````

---

## 🔍 Endpoints

### 🟢 GET /tasks

**Description:** Fetch all tasks

**Response:**
```json
{
  "tasks": [
    {
      "id": "1",
      "title": "Task 1",
      "description": "First task",
      "due_date": "2025-07-20T00:00:00Z",
      "status": "Pending"
    }
  ]
}
````

---

### 🟢 GET /tasks/\:id

**Description:** Fetch a task by ID

**Example:**

```
GET /tasks/1
```

**Response:**

```json
{
  "id": "1",
  "title": "Task 1",
  "description": "First task",
  "due_date": "2025-07-20T00:00:00Z",
  "status": "Pending"
}
```

**If not found:**

```json
{
  "error": "Task not found"
}
```

---

### 🟡 POST /tasks

**Description:** Create a new task

**Request Body:**

```json
{
  "id": "4",
  "title": "New Task",
  "description": "This is a new task",
  "due_date": "2025-07-22T00:00:00Z",
  "status": "Pending"
}
```

**Response:**

```json
{
  "message": "Task created"
}
```

---

### 🟠 PUT /tasks/\:id

**Description:** Update an existing task

**Request Example:**

```
PUT /tasks/1
```

**Request Body:**

```json
{
  "title": "Updated Task",
  "description": "Updated description",
  "due_date": "2025-07-25T00:00:00Z",
  "status": "Completed"
}
```

**Response:**

```json
{
  "message": "Task updated"
}
```

**If not found:**

```json
{
  "message": "Task not found"
}
```

---

### 🔴 DELETE /tasks/\:id

**Description:** Delete a task by ID

**Example:**

```
DELETE /tasks/1
```

**Response:**

```json
{
  "message": "Task removed"
}
```

**If not found:**

```json
{
  "message": "Task not found"
}
```

---

## 📬 Notes

* All dates should be in **ISO 8601 format**, e.g., `"2025-07-22T00:00:00Z"`
* All responses are in JSON
* Status codes:

  * `200 OK` for success
  * `201 Created` for new resources
  * `400 Bad Request` for bad input
  * `404 Not Found` when a task doesn’t exist

---

## ✅ Testing

Use Postman or curl to test the API:

```bash
curl http://localhost:8080/tasks
```

---

🧑‍💻 Built with Go + Gin 💙

```