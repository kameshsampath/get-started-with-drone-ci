package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type Post struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Text string `json:"text"`
}

var (
	db  *sql.DB
	err error
)

func init() {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s`, os.Getenv("MARIADB_USER"), os.Getenv("MARIADB_PASSWORD"), os.Getenv("MARIADB_HOST"), os.Getenv("MARIADB_DATABASE"))
	log.Printf("Using DSN %s", dsn)
	for {
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}
		err = db.Ping()
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
		continue
	}
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	e.POST("/add", addPost)
	e.Logger.Fatal(e.Start(":8080"))
}

// e.POST("/add", add)
func addPost(c echo.Context) error {
	p := new(Post)
	if err := c.Bind(p); err != nil {
		return err
	}
	stmt, err := db.Prepare("insert into posts(name,text) values(?,?)")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(p.Name, p.Text)
	if err != nil {
		return err
	}
	p.ID, err = res.LastInsertId()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, p)
}
