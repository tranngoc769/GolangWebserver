package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/common/log"
	"github.com/tranngoc769/golang/model"
)

type Config struct {
	Domain string `json:"host"`
	Path   string `json:"path"`
}

func dateFormat(d int) string {
	intTime := int64(d)
	t := time.Unix(intTime, 0)
	layout := " 15:04:05 2006-01-02"
	return t.Format(layout)
}
func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
func Index(w http.ResponseWriter, r *http.Request) {
	cfg := LoadConfiguration("config.json")
	tmpl := template.Must(template.ParseFiles("templates/main.html"))

	vars := mux.Vars(r)
	username := vars["username"]
	if username == "" || username == "favicon.ico" {
		username = "tnquang769"
	}
	var instaModel model.Insta
	instaModel, err := get_content(username, cfg)
	if err != nil {
		w.Write([]byte("Not found " + username))
		return
	}
	for i, node := range instaModel.Graphql.User.EdgeOwnerToTimelineMedia.Edges {
		date_format := dateFormat(node.Node.TakenAtTimestamp)
		instaModel.Graphql.User.EdgeOwnerToTimelineMedia.Edges[i].Node.TakenAt = date_format
	}
	tmpl.Execute(w, instaModel)
}
func main() {
	r := NewRouter()
	router := r.PathPrefix("/").Subrouter()
	router.HandleFunc("/{username}", Index)
	router.HandleFunc("/", Index)
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
func get_content(username string, cfg Config) (model.Insta, error) {
	url := cfg.Domain + username + cfg.Path
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
	req.AddCookie(&http.Cookie{Name: "csrftoken", Value: "MTufVyiaobNx9Yp3StkTiak8flVMFpbU"})
	req.AddCookie(&http.Cookie{Name: "sessionid", Value: "6894997460%3A3GFaUUXquhlDdN%3A23"})
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
	}
	return model.Insta{}, errors.New((string)(resp.StatusCode))
}
