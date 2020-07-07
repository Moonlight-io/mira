package main

import (
	"github.com/Moonlight-io/mira"
)

func main() {
	r, _ := mira.Init(mira.ReadCredsFromFile("login.conf"))
	c, _ := r.Subreddit("all").StreamSubmissions()
	for {
		post := <-c
		r.Submission(post.GetId()).Save("hello there")
	}
}
