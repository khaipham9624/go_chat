package restserver

import (
	"fmt"
	"gochat/src/server/groupmanager"
	"gochat/src/server/inmemorygroupmanager"
	"gochat/src/server/user"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

var groupManager groupmanager.GroupManager
var sessions = make(map[string]string)

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err != nil || sessions[cookie.Value] == "" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		return
	}

	if cookie, err := r.Cookie("session_id"); err == nil {
		if sessions[cookie.Value] != "" {
			http.Redirect(w, r, "/chat", http.StatusSeeOther)
			return
		}
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func createGroupHandler(w http.ResponseWriter, r *http.Request) {
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

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	userLogin := user.UserLogin{
		UserName: username,
		Password: password,
	}
	fmt.Println(userLogin)
	if userLogin.Login() {
		sessionId := fmt.Sprintf("%d", time.Now().UnixNano())
		sessions[sessionId] = username
		http.SetCookie(w, &http.Cookie{
			Name:  "session_id",
			Value: sessionId,
			Path:  "/",
		})
		http.Redirect(w, r, "/chat", http.StatusSeeOther)
	} else {
		w.Write([]byte("Login failed!"))
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
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

func chatHandler(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session_id")
	username := sessions[cookie.Value]
	w.Write([]byte("Welcome to chat " + username))
}

func Start(listeningPort uint) {
	groupManager = inmemorygroupmanager.NewInMemoryGroupManager()

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/register", registerHandler)
	mux.HandleFunc("/chat", chatHandler)
	mux.HandleFunc("/createGroup", createGroupHandler)

	// Serve static files (CSS, JS, images)
	// 1. Serve static files from the "static" folder
	fs := http.FileServer(http.Dir("./static"))
	// 2. Handle URL path "/static/*"
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	port := ":" + strconv.Itoa(int(listeningPort))
	http.ListenAndServe(port, mux)
}
