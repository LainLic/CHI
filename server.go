package main

import (
	"html/template"
	"log"
	"net/http"

	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/neerajbg/chi-htmx/model"
)

var DBConn *sql.DB

func init() {
	dsn := "host=localhost port5432 user=postgres password=root dbname=Chi-Demo sslmode=disable"
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Println("Error in DB", err)
	}

	DBConn = db
	log.Println("Database Connection successful")
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", homeHandler)
	r.Get("/user-info", userInfoHandler)

	r.Get("/posts", postHandler)
	http.ListenAndServe(":3009", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	ctx := make(map[string]string)

	ctx["Name"] = "Test"

	t, _ := template.ParseFiles("templates/index.html")

	err := t.Execute(w, ctx)

	if err != nil {
		log.Println("Error in execution")

	}

}

func userInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User info from API server"))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("OK")

	var posts []model.Post

	sql := "select * from posts"

	rows, err := DBConn.Query(sql)

	defer DBConn.Close()

	if err != nil {
		log.Println("error in DB execution")
	}

	for rows.Next() {
		data := model.Post{}

		err := rows.Scan(&data.Id, &data.Title)

		if err != nil {
			log.Println(err)
		}
		posts = append(posts, data)
	}
	log.Println(posts)

	w.Write([]byte("User info from API server"))
}
