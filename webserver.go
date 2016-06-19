package main

import(
    "fmt"
    "time"
    "strconv"
    "encoding/json"
    "net/http"
)

func getHomepage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "files/index.html")
}

func getWellKnown(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "files" + r.URL.Path)
}

func getFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
        fmt.Println(time.Now().String() + " " + r.URL.Path[1:])
}
func getTop10(w http.ResponseWriter, r *http.Request) {
    p := getTop10Posts()
    encoder := json.NewEncoder(w)
    err := encoder.Encode(&p)
    checkErr(err)
}

func getPostJson(w http.ResponseWriter, r *http.Request) {
    id, err_conv := strconv.Atoi(r.URL.Path[9:])
    checkErr(err_conv)
    p := getPost(id)

    encoder := json.NewEncoder(w)
    err := encoder.Encode(&p)
    checkErr(err)

}

func newPost(w http.ResponseWriter, r *http.Request) {
    var p Post
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&p)
    checkErr(err) 

    insertPost(p)    
    fmt.Println(p)
}

func redirect(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "https://nwaggoner.com"+r.RequestURI, http.StatusMovedPermanently)
}

func main() {
    http.HandleFunc("/", getHomepage)
    http.HandleFunc("/.well-known/", getWellKnown)
    http.HandleFunc("/top", getTop10)
    http.HandleFunc("/newPost", newPost)
    http.HandleFunc("/files/", getFile)
    http.HandleFunc("/getPost/", getPostJson)

    go func() {
        fmt.Println("Starting TLS Listner")
        err := http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/nwaggoner.com/cert.pem", "/etc/letsencrypt/keys/0000_key-certbot.pem", nil)
        fmt.Println("TLS:", err)
    }()
    fmt.Println("Starting standard Listner")
    err:= http.ListenAndServe(":80", http.HandlerFunc(redirect))
    fmt.Println("HTTP:", err)
}
