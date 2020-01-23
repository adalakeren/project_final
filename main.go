package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	//"text/template"
	

	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/go-sessions"
	// "os"
)

//var db *sql.DB
var err error

type user struct {
	ID        int
	Username  string
	FirstName string
	LastName  string
	Password  string
}

// func connect_db() {
// 	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1)/go_db")

// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

func dbConn() (db *sql.DB) {

	dbDriver := "mysql"   // Database driver
	dbUser := "root"      // Mysql username
	dbPass := "" // Mysql password
	dbName := "go_db"   // Mysql schema

	// Realize the connection with mysql driver
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	// If error stop the application
	if err != nil {
		panic(err.Error())
	}

	// Return db object to be used by other functions
	return db
}

func routes() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homepublic)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.HandleFunc("/about", about)
	http.HandleFunc("/kontak", kontak)

	http.HandleFunc("/berhasil", berhasil)

	http.HandleFunc("/home", home)

	http.HandleFunc("/kontakadmin", kontakadmin)
	http.HandleFunc("/deletekontak", DeleteKontakAdmin)

	http.HandleFunc("/artikeladmin", artikeladmin)
	http.HandleFunc("/deleteartikel", DeleteArtikelAdmin)
	http.HandleFunc("/tambahartikel", tambahartikel)
	http.HandleFunc("/editartikeldata", updateartikeladmin)
	http.HandleFunc("/update", Updatedataartikel)
}

/*====================Modul kontak Admin=======================*/
type KontakAdmin struct {
    IDKontak    int
    EmailKontak  string
    NamaKontak string
    KeteranganKontak string
}

func kontakadmin(w http.ResponseWriter, r *http.Request) {
	// Open database connection
	db := dbConn()

	session := sessions.Start(w, r)
	if len(session.GetString("username")) == 0 {
		http.Redirect(w, r, "/login", 301)
	}

	var t, err = template.ParseFiles("views/kontak.html")
	
	// Prepare a SQL query to select all data from database and threat errors
	selDB, err := db.Query("SELECT id_kontak,email,nama,keterangan FROM kontak")
	if err != nil {
		panic(err.Error())
	}

	// Call the struct to be rendered on template
	n := KontakAdmin{}

	// Create a slice to store all data from struct
	res := []KontakAdmin{}

	//n.UsernameSsion = session.GetString("username")

	// Read all rows from database
	for selDB.Next() {
		// Must create this variables to store temporary query
		var id_kontak int
		var email, nama,keterangan string

		// Scan each row storing values from the variables above and check for errors
		err = selDB.Scan(&id_kontak,&email, &nama, &keterangan)
		if err != nil {
			panic(err.Error())
		}

		// Get the Scan into the Struct
		n.IDKontak = id_kontak
		n.NamaKontak = nama
		n.KeteranganKontak = keterangan
		n.EmailKontak = email

		//fmt.Println(username)

		// Join each row on struct inside the Slice
		res = append(res, n)

	}

	/*var data = map[string]string{
		"username": session.GetString("username"),
		"message":  "Welcome to the Go !",
	}*/

	

	// Execute template `Index` from `tmpl/*` folder and send the struct
	// (View the file: `tmpl/Index`
	t.Execute(w, res)

	// Close database connection
	defer db.Close()
}

func DeleteKontakAdmin(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	session := sessions.Start(w, r)
	if len(session.GetString("username")) == 0 {
		http.Redirect(w, r, "/login", 301)
	}
	// Get the URL `?id=X` parameter
	nId := r.URL.Query().Get("id")

	// Prepare the SQL Delete
	delForm, err := db.Prepare("DELETE FROM kontak WHERE id_kontak=?")
	if err != nil {
		panic(err.Error())
	}

	// Execute the Delete SQL
	delForm.Exec(nId)

	// Show on console the action
	log.Println("DELETE")

	defer db.Close()

	// Redirect a HOME
	http.Redirect(w, r, "/kontakadmin", 301)
}
/*=================END Kontak Admin=================================================================*/

/*====================Modul Artikel Admin=======================*/
type ArtikelAdmin struct {
    IDArtikel    int
    NamaArtikel  string
    KeteranganArtikel  string
}

type EditArtikel struct {
    IDArtikelEdit    int
    NamaArtikelEdit  string
    KeteranganEdit  string
}

func artikeladmin(w http.ResponseWriter, r *http.Request) {
	// Open database connection
	db := dbConn()

	session := sessions.Start(w, r)
	if len(session.GetString("username")) == 0 {
		http.Redirect(w, r, "/login", 301)
	}

	var t, err = template.ParseFiles("views/artikel.html")
	
	// Prepare a SQL query to select all data from database and threat errors
	selDB, err := db.Query("SELECT id_artikel,nama_artikel FROM artikel")
	if err != nil {
		panic(err.Error())
	}

	// Call the struct to be rendered on template
	n := ArtikelAdmin{}

	// Create a slice to store all data from struct
	res := []ArtikelAdmin{}

	//n.UsernameSsion = session.GetString("username")

	// Read all rows from database
	for selDB.Next() {
		// Must create this variables to store temporary query
		var id_artikel int
		var nama_artikel string

		// Scan each row storing values from the variables above and check for errors
		err = selDB.Scan(&id_artikel,&nama_artikel)
		if err != nil {
			panic(err.Error())
		}

		// Get the Scan into the Struct
		n.IDArtikel = id_artikel
		n.NamaArtikel = nama_artikel

		//fmt.Println(username)

		// Join each row on struct inside the Slice
		res = append(res, n)

	}

	/*var data = map[string]string{
		"username": session.GetString("username"),
		"message":  "Welcome to the Go !",
	}*/

	

	// Execute template `Index` from `tmpl/*` folder and send the struct
	// (View the file: `tmpl/Index`
	t.Execute(w, res)

	// Close database connection
	defer db.Close()
}

func DeleteArtikelAdmin(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	session := sessions.Start(w, r)
	if len(session.GetString("username")) == 0 {
		http.Redirect(w, r, "/login", 301)
	}
	// Get the URL `?id=X` parameter
	nId := r.URL.Query().Get("id")

	// Prepare the SQL Delete
	delForm, err := db.Prepare("DELETE FROM artikel WHERE id_artikel=?")
	if err != nil {
		panic(err.Error())
	}

	// Execute the Delete SQL
	delForm.Exec(nId)

	// Show on console the action
	log.Println("DELETE")

	defer db.Close()

	// Redirect a HOME
	http.Redirect(w, r, "/artikeladmin", 301)
}

func tambahartikel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	if r.Method != "POST" {
		http.ServeFile(w, r, "views/addartikel.html")
		return
	}

	nama := r.FormValue("nama")
	artikel := r.FormValue("artikel")

	stmt, err := db.Prepare("INSERT INTO artikel SET nama_artikel=?, keterangan=?")
	if err == nil {
		_, err := stmt.Exec(&nama, &artikel)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/artikeladmin", http.StatusSeeOther)
		return
	}
}

func updateartikeladmin(w http.ResponseWriter, r *http.Request) {
	// Open database connection
	db := dbConn()

	if r.Method != "POST" {
		//http.ServeFile(w, r, "views/editartikeldata.html")
		var t, err = template.ParseFiles("views/editartikeldata.html")

		nId := r.URL.Query().Get("id")

		// Prepare a SQL query to select all data from database and threat errors
		selDB, err := db.Query("SELECT id_artikel,nama_artikel,keterangan FROM artikel WHERE id_artikel=?", nId)
		if err != nil {
			panic(err.Error())
		}

		// Call the struct to be rendered on template
		n := ArtikelAdmin{}

		// Create a slice to store all data from struct
		res := []ArtikelAdmin{}

		//n.UsernameSsion = session.GetString("username")

		// Read all rows from database
		for selDB.Next() {
			// Must create this variables to store temporary query
			var id_artikel int
			var nama_artikel,keterangan string

			// Scan each row storing values from the variables above and check for errors
			err = selDB.Scan(&id_artikel,&nama_artikel,&keterangan)
			if err != nil {
				panic(err.Error())
			}

			// Get the Scan into the Struct
			n.IDArtikel = id_artikel
			n.NamaArtikel = nama_artikel
			n.KeteranganArtikel=keterangan

			//fmt.Println(username)

			// Join each row on struct inside the Slice
			res = append(res, n)

		}

		t.Execute(w, res)
		defer db.Close()
		return
	}

	
}

func Updatedataartikel(w http.ResponseWriter, r *http.Request) {

	db := dbConn()

	if r.Method == "POST" {

		// Get the values from form
		name := r.FormValue("nama")
		artikel := r.FormValue("artikel")
		id := r.FormValue("uid") // This line is a hidden field on form (View the file: `tmpl/Edit`)

		// Prepare the SQL Update
		insForm, err := db.Prepare("UPDATE artikel SET nama_artikel=?, keterangan=? WHERE id_artikel=?")
		if err != nil {
			panic(err.Error())
		}

		// Update row based on hidden form field ID
		insForm.Exec(name, artikel, id)

		// Show on console the action
		log.Println("UPDATE: Name: " + name + " | Artikel: " + artikel)
	}

	defer db.Close()

	// Redirect to Home
	http.Redirect(w, r, "/artikeladmin", 301)
}
/*=================END Artikel Admin=================================================================*/

func main() {
	//connect_db()
	routes()

	//defer db.Close()

	fmt.Println("Server running on port :6969")
	http.ListenAndServe(":6969", nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	var t, err = template.ParseFiles("public/aboutpublic.html")

	if err != nil {
		panic(err.Error())
	}

	t.Execute(w, "")
}

func berhasil(w http.ResponseWriter, r *http.Request) {
	var t, err = template.ParseFiles("public/berhasilpublic.html")

	if err != nil {
		panic(err.Error())
	}

	t.Execute(w, "")
}

/*==============Modul Kontak Us Halaman Publik====================*/
func kontak(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	if r.Method != "POST" {
		http.ServeFile(w, r, "public/kontakpublic.html")
		return
	}

	nama := r.FormValue("nama")
	email := r.FormValue("email")
	keterangan := r.FormValue("keterangan")

	stmt, err := db.Prepare("INSERT INTO kontak SET email=?, nama=?, keterangan=?")
	if err == nil {
		_, err := stmt.Exec(&nama, &email, &keterangan)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/berhasil", http.StatusSeeOther)
		return
	}
}
/*==============End Modul Kontak Us Halaman Publik====================*/

func checkErr(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {

		fmt.Println(r.Host + r.URL.Path)

		http.Redirect(w, r, r.Host+r.URL.Path, 301)
		return false
	}

	return true
}

func QueryUser(username string) user {
	db := dbConn()

	var users = user{}
	err = db.QueryRow(`
		SELECT id, 
		username, 
		first_name, 
		last_name, 
		password 
		FROM users WHERE username=?
		`, username).
		Scan(
			&users.ID,
			&users.Username,
			&users.FirstName,
			&users.LastName,
			&users.Password,
		)
	return users
}

type Userinfo struct {
    Id    int
    UserName  string
    FirstName string
}

//var tmpl = template.Must(template.ParseGlob("views/*"))

// Function Index shows all values on home
func home(w http.ResponseWriter, r *http.Request) {
	// Open database connection
	db := dbConn()

	session := sessions.Start(w, r)
	if len(session.GetString("username")) == 0 {
		http.Redirect(w, r, "/login", 301)
	}

	var t, err = template.ParseFiles("views/home.html")
	
	// Prepare a SQL query to select all data from database and threat errors
	selDB, err := db.Query("SELECT id,username,first_name FROM users")
	if err != nil {
		panic(err.Error())
	}

	// Call the struct to be rendered on template
	n := Userinfo{}

	// Create a slice to store all data from struct
	res := []Userinfo{}

	//n.UsernameSsion = session.GetString("username")

	// Read all rows from database
	for selDB.Next() {
		// Must create this variables to store temporary query
		var id int
		var username, first_name string

		// Scan each row storing values from the variables above and check for errors
		err = selDB.Scan(&id, &username, &first_name)
		if err != nil {
			panic(err.Error())
		}

		// Get the Scan into the Struct
		n.Id = id
		n.UserName = username
		n.FirstName = first_name

		//fmt.Println(username)

		// Join each row on struct inside the Slice
		res = append(res, n)

	}

	/*var data = map[string]string{
		"username": session.GetString("username"),
		"message":  "Welcome to the Go !",
	}*/

	

	// Execute template `Index` from `tmpl/*` folder and send the struct
	// (View the file: `tmpl/Index`
	t.Execute(w, res)

	// Close database connection
	defer db.Close()
}


type Artikel struct {
    NamaArtikel  string
    KeteranganArtikel string
}

func homepublic(w http.ResponseWriter, r *http.Request) {
	// Open database connection
	db := dbConn()

	var t, err = template.ParseFiles("public/homepublic.html")
	
	// Prepare a SQL query to select all data from database and threat errors
	selDB, err := db.Query("SELECT nama_artikel,keterangan FROM artikel")
	if err != nil {
		panic(err.Error())
	}

	// Call the struct to be rendered on template
	n := Artikel{}

	// Create a slice to store all data from struct
	res := []Artikel{}

	//n.UsernameSsion = session.GetString("username")

	// Read all rows from database
	for selDB.Next() {
		// Must create this variables to store temporary query
		var nama_artikel, keterangan string

		// Scan each row storing values from the variables above and check for errors
		err = selDB.Scan(&nama_artikel, &keterangan)
		if err != nil {
			panic(err.Error())
		}

		// Get the Scan into the Struct
		n.NamaArtikel = nama_artikel
		n.KeteranganArtikel = keterangan

		//fmt.Println(username)

		// Join each row on struct inside the Slice
		res = append(res, n)

	}

	/*var data = map[string]string{
		"username": session.GetString("username"),
		"message":  "Welcome to the Go !",
	}*/

	

	// Execute template `Index` from `tmpl/*` folder and send the struct
	// (View the file: `tmpl/Index`
	t.Execute(w, res)

	// Close database connection
	defer db.Close()
}
/*func home(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	if len(session.GetString("username")) == 0 {
		http.Redirect(w, r, "/login", 301)
	}

	var data = map[string]string{
		"username": session.GetString("username"),
		"message":  "Welcome to the Go !",
	}
	var t, err = template.ParseFiles("views/home.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	t.Execute(w, data)
	return

}*/

func register(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	if r.Method != "POST" {
		http.ServeFile(w, r, "views/register.html")
		return
	}

	username := r.FormValue("email")
	first_name := r.FormValue("first_name")
	last_name := r.FormValue("last_name")
	password := r.FormValue("password")

	users := QueryUser(username)

	if (user{}) == users {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		if len(hashedPassword) != 0 && checkErr(w, r, err) {
			stmt, err := db.Prepare("INSERT INTO users SET username=?, password=?, first_name=?, last_name=?")
			if err == nil {
				_, err := stmt.Exec(&username, &hashedPassword, &first_name, &last_name)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
		}
	} else {
		http.Redirect(w, r, "/register", 302)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	//db := dbConn()

	session := sessions.Start(w, r)
	if len(session.GetString("username")) != 0 && checkErr(w, r, err) {
		http.Redirect(w, r, "/", 302)
	}
	if r.Method != "POST" {
		http.ServeFile(w, r, "views/login.html")
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	users := QueryUser(username)

	//deskripsi dan compare password
	var password_tes = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))

	if password_tes == nil {
		//login success
		session := sessions.Start(w, r)
		session.Set("username", users.Username)
		session.Set("name", users.FirstName)
		http.Redirect(w, r, "/home", 302)
	} else {
		//login failed
		http.Redirect(w, r, "/login", 302)
	}

}

func logout(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	session.Clear()
	sessions.Destroy(w, r)
	http.Redirect(w, r, "/", 302)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	session := sessions.Start(w, r)
	if len(session.GetString("username")) == 0 {
		http.Redirect(w, r, "/login", 301)
	}
	// Get the URL `?id=X` parameter
	nId := r.URL.Query().Get("id")

	// Prepare the SQL Delete
	delForm, err := db.Prepare("DELETE FROM users WHERE id=?")
	if err != nil {
		panic(err.Error())
	}

	// Execute the Delete SQL
	delForm.Exec(nId)

	// Show on console the action
	log.Println("DELETE")

	defer db.Close()

	// Redirect a HOME
	http.Redirect(w, r, "/", 301)
}