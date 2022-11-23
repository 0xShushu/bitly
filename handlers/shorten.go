package handlers

import (
	"html/template"
	"net/http"
	"github.com/0xshushu/bitly/database"
	"github.com/0xshushu/bitly/utils"
)

//struct for template
type Page struct {
	Text string
	URL string
}

//handler which makes the shorten url
func Shorten(w http.ResponseWriter, r *http.Request) {
	//allowed protocols
	allowed := [3]string{"https://", "http://", "mailto:"}

	//get url param
	url := r.URL.Query().Get("url")
	
	//check if url is empty or using a now allowed protocol
	if url == "" {
		http.Error(w, "GET param url is empty", http.StatusInternalServerError)
		return
	} else if !utils.IsProtocolAllowed(url, allowed[:]) {
		http.Error(w, "Protocol not allowed, URLs have to start with http://, https:// or mailto:", http.StatusInternalServerError)
		return
	}
	
	//generate random id
	id := utils.RandString()

	//save url and id into database
	err := database.Insert(id, url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//making template struct
	page := Page {
		Text: "Shorten URL: ",
		URL: "http://127.0.0.1:3000/s/"+id,
	}

	//execute template
	t, err := template.ParseFiles("templates/index.html")	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, page); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}