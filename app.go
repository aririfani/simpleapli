package main

import (
	"database/sql"
	"simpleapi/handlers"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()
	db := initDB("storage.db")
	migrate(db)

	e.GET("/tasks", handlers.GetTasks(db))
	e.POST("/tasks", handlers.PutTask(db))
	e.PUT("/task", handlers.EditTask(db))
	e.DELETE("/task/:id", handlers.DeleteTask(db))

	e.Logger.Fatal(e.Start(":9000"))

}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}

	if db == nil {
		panic(db)
	}

	return db
}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS tasks(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		status INTEGER
    );
	`

	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
}
