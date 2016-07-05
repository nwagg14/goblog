package main 

import (
    "time"
)

type Blog struct {
    Name string  `json:"name"`
    About string `json:"about"`
}

type Submission struct {
    Form Post `json:"post"`
    Password string `json:"pass"`
}

type Post struct {
    Id    int    `json:"id"`
    Title string `json:"title"`
    Content string `json:"content"`
    Author string `json:"author"`
    Date time.Time  `json:"date"`
}
