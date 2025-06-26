package restserver

import (
	"fmt"
	"gochat/src/server/groupmanager"
	"gochat/src/server/inmemorygroupmanager"
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

func createUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	r.FormValue("username")
	r.FormValue("fullname")
}

func login(w http.ResponseWriter, r *http.Request) {

}

func register(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/register.html"))
	tmpl.Execute(w, nil)
}

func Start(listeningPort uint) {
	groupManager = inmemorygroupmanager.NewInMemoryGroupManager()

	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/processLogin", login)
	mux.HandleFunc("/register", register)
	mux.HandleFunc("/processRegister", register)
	mux.HandleFunc("/createGroup", createGroup)
	mux.HandleFunc("/createUser", createUser)

	// Serve static files (CSS, JS, images)
	// 1. Serve static files from the "static" folder
	fs := http.FileServer(http.Dir("./static"))
	// 2. Handle URL path "/static/*"
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	port := ":" + strconv.Itoa(int(listeningPort))
	http.ListenAndServe(port, mux)
}
