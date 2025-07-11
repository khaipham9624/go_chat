package restserver

import (
	"fmt"
	"gochat/src/server/groupmanager"
	"gochat/src/server/inmemorygroupmanager"
	"gochat/src/server/user"
	"html/template"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

var groupManager groupmanager.GroupManager

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		return
	}
	fmt.Println(r.URL.Path)
	// fmt.Println(os.Getwd())
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)

	// http.ServeFile(w, r, "frontend/html/index.html")
}

func createGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Group")
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte("error"))
	}
	groupName := r.FormValue("groupname")
	fmt.Println("Group:" + groupName)

	userIds := []uuid.UUID{uuid.New(), uuid.New()}
	groupId := groupManager.CreateGroup(groupName, uuid.New(), userIds)
	fmt.Println(groupId)

	w.Write([]byte("Created Group " + groupName + "\n"))
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	userLogin := user.UserLogin{
		UserName: username,
		Password: password,
	}
	if userLogin.Login() {
		w.Write([]byte("Login successufully!"))
	} else {
		w.Write([]byte("Login failed!"))
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tmpl := template.Must(template.ParseFiles("templates/register.html"))
		tmpl.Execute(w, nil)
	case http.MethodPost:
		r.ParseForm()
		username := r.FormValue("username")
		fullname := r.FormValue("fullname")
		password := r.FormValue("password")

		userRegister := user.UserRegister{
			UserName: username,
			Password: password,
			FullName: fullname,
			Email:    "",
		}
		fmt.Println(userRegister)
		if userRegister.Register() {
			w.Write([]byte("Register successufully!"))
		} else {
			w.Write([]byte("Register failed!"))
		}
		// create user
	default:

	}

}

func Start(listeningPort uint) {
	groupManager = inmemorygroupmanager.NewInMemoryGroupManager()

	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/register", register)
	mux.HandleFunc("/createGroup", createGroup)

	// Serve static files (CSS, JS, images)
	// 1. Serve static files from the "static" folder
	fs := http.FileServer(http.Dir("./static"))
	// 2. Handle URL path "/static/*"
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	port := ":" + strconv.Itoa(int(listeningPort))
	http.ListenAndServe(port, mux)
}
