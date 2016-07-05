package main 

import (
    "fmt"
    "time"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func createTables() {
    db, err := sql.Open("sqlite3", "./blog.db")
    checkErr(err)

    stmt_posts, err_posts := db.Prepare(
        `CREATE TABLE POSTS(
            ID INT PRIMARY KEY     NOT NULL,
            TITLE          TEXT    NOT NULL,
            AUTHOR         TEXT    NOT NULL,
            DATE           TEXT    NOT NULL,
            CONTENT        TEXT    NOT NULL
     );`)
    
    checkErr(err_posts)
    _, err_exec := stmt_posts.Exec()
    checkErr(err_exec)
    db.Close()
}

func getNextPostId(db *sql.DB) (id int) {
    result, err_query := db.Query(`
        SELECT MAX(ID) FROM POSTS;
   `)

    checkErr(err_query)

    defer result.Close()
    if(result.Next()) {
        result.Scan(&id)
        id  = id + 1
    } else {
        id = -1
    }  
    return
}

func getPost(id int) (post Post){
    fmt.Println("getPost:", id)
    db, err := sql.Open("sqlite3", "./blog.db")
    checkErr(err)

    rows, err_query := db.Query(`
        SELECT * FROM POSTS WHERE ID = ?;
   `, id)

    checkErr(err_query)
   
    defer rows.Close()
    if(rows.Next()) {
        var timeStr string        
        err = rows.Scan(&post.Id, &post.Title, &post.Author, &timeStr, &post.Content)
        checkErr(err)

        form := "2006-01-02 15:04:05-07:00"
	t, err2 := time.Parse(form, timeStr)
        checkErr(err2)
        post.Date = t

    }
    db.Close()
    return
}
func getTop10Posts() (posts []Post) {
    db, err := sql.Open("sqlite3", "./blog.db")
    checkErr(err)

    rows, err_query := db.Query(
        `SELECT * FROM POSTS ORDER BY ID DESC LIMIT 10;`)

    checkErr(err_query)
   
    defer rows.Close()
    for rows.Next() {
        var timeStr string
        var post = Post{}        
        err = rows.Scan(&post.Id, &post.Title, &post.Author, &timeStr, &post.Content)
        checkErr(err)

        form := "2006-01-02 15:04:05-07:00"
	t, err2 := time.Parse(form, timeStr)
        checkErr(err2)
        post.Date = t
        posts = append(posts, post)

    }
    db.Close()
    return

}

func insertPost(post Post) {
    db, err := sql.Open("sqlite3", "./blog.db")
    checkErr(err)

    stmt_posts, err_posts := db.Prepare(
        `INSERT INTO POSTS (ID,TITLE,CONTENT,DATE,AUTHOR)
        VALUES (?, ?, ?, ?, ?);`)
    checkErr(err_posts)
    nextPostId := getNextPostId(db) 
    fmt.Println(nextPostId)
    _, err_exec1 := stmt_posts.Exec(nextPostId, post.Title, post.Content, post.Date, post.Author)
    checkErr(err_exec1)
    db.Close()
}

