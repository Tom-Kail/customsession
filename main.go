package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

var globalSessions *session.Manager

func init() {
	globalSessions, _ = session.NewManager("file", `{
			"cookieName":"tomId",
			"enableSetCookie,omitempty": true,
			"gclifetime":600,
			"maxLifetime": 3600,
			"secure": false,
			"sessionIDHashFunc": "aes",
			"sessionIDHashKey": "1312312",
			"cookieLifeTime": 3600,
			"providerConfig": "./session"
		}
		`)
	go globalSessions.GC()
}

func login(w http.ResponseWriter, r *http.Request) {
	sess, _ := globalSessions.SessionStart(w, r)
	//defer sess.SessionRelease(w)
	username := sess.Get("username")
	fmt.Println("sessionId:", sess.SessionID())
	if username == nil {
		sess.Set("username", r.Form["username"])
	}
	beego.Info("cookie:", r.Cookies())
	beego.Info(username)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.tpl")
		t.Execute(w, nil)
	} else {
		sess.Set("username", r.Form["username"])
	}
}

func main() {
	http.HandleFunc("/", login)
	http.ListenAndServe(":8089", nil)
}
