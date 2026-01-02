# Project: hw4

## Overview
`hw4` is a Go-based REST API designed to manage and retrieve information regarding students and class schedules. It utilizes the **Echo** web framework for handling HTTP requests and **pgx** for high-performance PostgreSQL database interactions. The project follows a standard clean architecture layout.

## Key Technologies
- **Language:** Go (1.24.0)
- **Web Framework:** [Echo v4](https://echo.labstack.com/)
- **Database Driver:** [pgx v5](https://github.com/jackc/pgx) (PostgreSQL)
- **Configuration:** [godotenv](https://github.com/joho/godotenv)

## Project Structure
```text
.
├── cmd/
│   └── app/
│       └── main.go       # Application entry point
├── internal/
│   ├── config/           # Configuration management
│   ├── handler/          # HTTP request handlers
│   ├── model/            # Data models and structs
│   ├── server/           # Server setup and route mapping
│   ├── service/          # Business logic layer
│   └── storage/          # Database interaction layer
│       └── postgres/     # PostgreSQL specific repository implementation
├── .env                  # Environment variables file
├── go.mod                # Go module definition
└── GEMINI.md             # Project documentation (this file)
```

## Setup & Configuration

### Environment Variables
The application uses a `.env` file for configuration. Ensure the following variable is set:

- `DATABASE_URL`: The connection string for your PostgreSQL database.
  - Example: `postgres://user:password@localhost:5432/dbname`

## Building and Running

1.  **Install Dependencies:**
    ```bash
    go mod tidy
    ```

2.  **Run the Application:**
    Execute the entry point located in `cmd/app/`:
    ```bash
    go run cmd/app/main.go
    ```
    The server will start on port `8080`.

3.  **Build Binary:**
    ```bash
    go build -o hw4 cmd/app/main.go
    ./hw4
    ```

## API Endpoints

The application exposes the following GET endpoints:

| Method | Endpoint                 | Handler Function     | Description                                      |
| :----- | :----------------------- | :------------------- | :----------------------------------------------- |
| `GET`  | `/student/:id`           | `GetStudent`         | Retrieves details for a specific student by ID.  |
| `GET`  | `/all_class_schedule`    | `GetAllSchedules`    | Retrieves the complete schedule of all classes.  |
| `GET`  | `/schedule/group/:id`    | `GetGroupSchedule`   | Retrieves the class schedule for a specific group.|

## Database Schema (Inferred)

The application interacts with a PostgreSQL database. Based on `internal/storage/postgres/repository.go` and `internal/model/model.go`, the expected schema includes:

### `students` Table
| Column         | Type      | Description |
| :------------- | :-------- | :---------- |
| `student_id`   | int       | Primary Key |
| `student_name` | text      |             |
| `birth_date`   | timestamp |             |
| `gender`       | text      |             |
| `group_id`     | int       | Foreign Key |

### `groups` Table
| Column       | Type | Description |
| :----------- | :--- | :---------- |
| `group_id`   | int  | Primary Key |
| `group_name` | text |             |

### `class_schedule` Table
| Column       | Type      | Description |
| :----------- | :-------- | :---------- |
| `class_id`   | int       | Primary Key |
| `group_id`   | int       |             |
| `class_name` | text      |             |
| `class_date` | timestamp |             |
| `start_time` | timestamp |             |
| `end_time`   | timestamp |             |
| `room`       | text      |             |
| `teacher_id` | int       |             |
