package main

import(
    "fmt"
    "time"
    "strconv"
    "encoding/json"
    "net/http"
)

func log(url string) {
    fmt.Println(time.Now().String() + " " + url)
}

func getHomepage(w http.ResponseWriter, r *http.Request) {
    log(r.URL.Path[1:])
    http.ServeFile(w, r, "files/index.html")
}

func getPostHtml(w http.ResponseWriter, r *http.Request) {
    log(r.URL.Path[1:])
    http.ServeFile(w, r, "files/post.html")
}

func getFile(w http.ResponseWriter, r *http.Request) {
    log(r.URL.Path[1:])
    http.ServeFile(w, r, r.URL.Path[1:])
}
func getTopJson(w http.ResponseWriter, r *http.Request) {
    log(r.URL.Path[1:])
    p := getTop10Posts()
    encoder := json.NewEncoder(w)
    err := encoder.Encode(&p)
    checkErr(err)
}

func getPostJson(w http.ResponseWriter, r *http.Request) {
    log(r.URL.Path[1:])
    id, err_conv := strconv.Atoi(r.URL.Path[len("/data/post/"):])
    checkErr(err_conv)
    fmt.Println(id)
    p := getPost(id)

    encoder := json.NewEncoder(w)
    err := encoder.Encode(&p)
    checkErr(err)

}

func submitPost(w http.ResponseWriter, r *http.Request) {
    log(r.URL.Path[1:])
    var s Submission
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&s)
    checkErr(err)

    if(s.Password == "") {
        insertPost(s.Form)
        fmt.Println(s.Form)
        fmt.Fprintf(w, "true")
    } else {
        fmt.Fprintf(w, "false")
    }
}

func getWellKnown(w http.ResponseWriter, r *http.Request) {
    log(r.URL.Path[1:])
    http.ServeFile(w, r, "files" + r.URL.Path)
}

func redirect(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "https://nwaggoner.com"+r.RequestURI, http.StatusMovedPermanently)
}

func main() {
    http.HandleFunc("/.well-known/", getWellKnown)

    http.HandleFunc("/", getHomepage)
    http.HandleFunc("/top", getHomepage)
    http.HandleFunc("/post/", getPostHtml)
    http.HandleFunc("/files/", getFile)

    http.HandleFunc("/data/top", getTopJson)
    http.HandleFunc("/data/post/", getPostJson)
    http.HandleFunc("/data/submit", submitPost)

    go func() {
        fmt.Println("Starting TLS Listner")
        err := http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/nwaggoner.com/cert.pem", "/etc/letsencrypt/keys/0000_key-certbot.pem", nil)
        fmt.Println("TLS:", err)
    }()
    fmt.Println("Starting standard Listner")
    err:= http.ListenAndServe(":80", http.HandlerFunc(redirect))
    fmt.Println("HTTP:", err)
}
