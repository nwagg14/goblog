package main

import(
    "fmt"
    "strconv"
    "encoding/json"
    "net/http"
)

func getHomepage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "files/index.html")
}

func getFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
        fmt.Println(r.URL.Path[1:])
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

func main() {
    http.HandleFunc("/", getHomepage)
    http.HandleFunc("/top", getTop10)
    http.HandleFunc("/newPost", newPost)
    http.HandleFunc("/files/", getFile)
    http.HandleFunc("/getPost/", getPostJson)
    http.ListenAndServe(":80", nil)
}
