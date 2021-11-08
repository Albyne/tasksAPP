package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	
	"github.com/Albyne/tasksApp/models"
)

//go:embed templates/*
var tmpl embed.FS

//go:embed static/img*
//go:embed static/styles*
var files embed.FS

var templates = template.Must(template.New("").ParseFS(tmpl, "templates/*.html"))

func main() {
	fileStyle := http.FileServer(http.FS(files))

	models.CreateConection()

	router := http.NewServeMux()

	router.Handle("/static/", fileStyle)

	router.HandleFunc("/", Home)
	router.HandleFunc("/create", Create)
	router.HandleFunc("/insert", Insert)
	router.HandleFunc("/edit", Edit)
	router.HandleFunc("/update", Update)
	router.HandleFunc("/delete", Delete)

	http.ListenAndServe("localhost:8080", router)
}

//*************************************************************************************************************

func Home(w http.ResponseWriter, r *http.Request) {
	conexionEstablecida := models.CreateConection()

	w.WriteHeader(200)

	rows, err := conexionEstablecida.Query("SELECT * FROM tareas")
	if err != nil {
		log.Fatal(err)
	}

	task := models.Task{}
	tasks := []models.Task{}

	for rows.Next() {
        var id int
		var name string
		var description string

		err := rows.Scan(&id, &name, &description)
		if err != nil {
			log.Fatal(err)
		}
        
		task.Id = id
		task.Name = name
		task.Description = description

		tasks = append(tasks, task)
		fmt.Println(tasks)

	}

	templates.ExecuteTemplate(w, "home", tasks)

}

func Create(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "create", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")

		conexionEstablecida := models.CreateConection()

		insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO tareas(name, description )VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}

		insertarRegistros.Exec(name, description)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}

}

func Edit(w http.ResponseWriter, r *http.Request){

	idTarea := r.URL.Query().Get("id")
	fmt.Println(idTarea)

	conexionEstablecida := models.CreateConection()

	row, err := conexionEstablecida.Query("SELECT * FROM tareas WHERE id=?", idTarea)
	if err != nil {
		log.Fatal(err)
	}

	task := models.Task{}

	for row.Next() {
        var id int
		var name string
		var description string

		err := row.Scan(&id, &name, &description)
		if err != nil {
			log.Fatal(err)
		}
        
		task.Id = id
		task.Name = name
		task.Description = description

		
		fmt.Println(task)

	}

	templates.ExecuteTemplate(w, "update", task)
}

func Update(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")

		conexionEstablecida := models.CreateConection()

		modificarRegistros, err := conexionEstablecida.Prepare("UPDATE tareas SET name=?, description=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		modificarRegistros.Exec(name, description, id)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func Delete(w http.ResponseWriter, r *http.Request){
	idTarea := r.URL.Query().Get("id")
	fmt.Println(idTarea)

	conexionEstablecida := models.CreateConection()

	deleteTask, err := conexionEstablecida.Prepare("DELETE FROM tareas WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	deleteTask.Exec(idTarea)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}
