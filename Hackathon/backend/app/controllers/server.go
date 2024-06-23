package controllers

import (
	"backend/app/models"
	"backend/config"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"text/template"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("../frontend/app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.ChaeckSession(); !ok {
			err = fmt.Errorf("invalid session")
		}
	}
	return sess, err
}

var validPath = regexp.MustCompile("^/mypage/(heightedit|heightupdate|heightdelete|weightedit|weightupdate|weightdelete)/([0-9]+)$")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, qi)
	}
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/mypage", index)
	http.HandleFunc("/mypage/eat", eat)
	http.HandleFunc("/mypage/sleep", sleep)
	http.HandleFunc("/mypage/training", train)
	http.HandleFunc("/mypage/explanation", privateexplanation)
	http.HandleFunc("/mypage/setting", setting)
	http.HandleFunc("/mypage/heightnew", heightNew)
	http.HandleFunc("/mypage/heightsave", heightSave)
	http.HandleFunc("/mypage/heightedit/", parseURL(heightedit))
	http.HandleFunc("/mypage/heightupdate/", parseURL(heightupdate))
	http.HandleFunc("/mypage/heightdelete/", parseURL(heightdelete))
	http.HandleFunc("/mypage/weightnew", weightNew)
	http.HandleFunc("/mypage/weightsave", weightSave)
	http.HandleFunc("/mypage/weightedit/", parseURL(weightedit))
	http.HandleFunc("/mypage/weightupdate/", parseURL(weightupdate))
	http.HandleFunc("/mypage/weightdelete/", parseURL(weightdelete))
	http.HandleFunc("/explanation", publicexplanation)
	http.HandleFunc("/mypage/game", game)
	http.HandleFunc("/mypage/calendar", calendar)
	http.HandleFunc("/mypage/item", itemcheck)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
