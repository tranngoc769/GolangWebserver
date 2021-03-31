package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/prometheus/common/log"
	"github.com/tranngoc769/golang/model"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

type Test struct {
	LoggingPageID         string `json:"logging_page_id"`
	ShowSuggestedProfiles bool   `json:"show_suggested_profiles"`
	ShowFollowDialog      bool   `json:"show_follow_dialog"`
}

func AllBooks(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/main.html"))

	// vars := mux.Vars(r)
	// title := vars["title"]
	// page := vars["page"]
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	var instaModel model.Insta
	instaModel, err := get_content("ngt.lngoc")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		w.Write([]byte("Some error"))
	}

	fmt.Printf("%v\n", instaModel.Graphql.User.EdgeOwnerToTimelineMedia.Edges[0].Node)
	// for i, v := range instaModel.Graphql.User.EdgeOwnerToTimelineMedia.Edges {
	// 	fmt.Printf("%v - %v\n", i, v)
	// }
	tmpl.Execute(w, data)
}
func main() {
	r := NewRouter()
	// fs := http.FileServer(http.Dir("statics/"))
	// http.Handle("statics/", http.StripPrefix("statics/", fs))
	bookrouter := r.PathPrefix("/").Subrouter()
	bookrouter.HandleFunc("/", AllBooks)
	bookrouter.HandleFunc("/xx", AllBooks)
	// bookrouter.HandleFunc("/{title}/page/{page}", AllBooks)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Print("Running :" + port + "\n")
	http.ListenAndServe(":"+port, r)
}
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	staticDir := "/statics/"
	router.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))
	return router
}
func get_content(username string) (model.Insta, error) {
	url := "https://www.instagram.com/" + username + "/?__a=1"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return model.Insta{}, err
	}
	req.AddCookie(&http.Cookie{Name: "ig_did", Value: "31801E8F-FBD6-4C76-B3A8-B8770E2DD549"})
	req.AddCookie(&http.Cookie{Name: "sessionid", Value: "mid=YFod6gALAAHmYewmH1DQn9jVrIJs"})
	req.AddCookie(&http.Cookie{Name: "ig_nrcb", Value: "1"})
	req.AddCookie(&http.Cookie{Name: "shbid", Value: "9515"})
	req.AddCookie(&http.Cookie{Name: "fbm_124024574287414", Value: "base_domain=.instagram.com"})
	req.AddCookie(&http.Cookie{Name: "shbts", Value: "1616942700.529952"})
	req.AddCookie(&http.Cookie{Name: "ds_user_id", Value: "6894997460"})
	req.AddCookie(&http.Cookie{Name: "csrftoken", Value: "Pqt42EQ7tWpdR4qn4pdHFZJjopfx4bjo"})
	req.AddCookie(&http.Cookie{Name: "sessionid", Value: "6894997460%3AazACKNMvI3ol5N%3A11"})
	req.AddCookie(&http.Cookie{Name: "rur", Value: "VLL"})
	client := &http.Client{}
	resp, err2 := client.Do(req)
	if err2 != nil {
		return model.Insta{}, err2
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			return model.Insta{}, err
		}
		defer resp.Body.Close()
		var msg model.Insta
		err = json.Unmarshal(bodyBytes, &msg)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return model.Insta{}, err
		}
		return msg, nil
		// for i, _ := range msg.Graphql.User.EdgeOwnerToTimelineMedia.Edges {
		// 	fmt.Printf("Results: %v\n", i)
		// }
		// fmt.Printf("Results: %v\n", msg.Graphql.User.EdgeOwnerToTimelineMedia.Edges)
	}
	return model.Insta{}, errors.New((string)(resp.StatusCode))
}
