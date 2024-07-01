package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	post := fetchPost()
	likes := fetchPostLikes(post)
	comments := fetchPostComments(post)

	fmt.Println("likes:", likes)
	fmt.Println("comments:", comments)
	fmt.Println("took:", time.Since(start))
}

func fetchPost() string {
	time.Sleep(time.Millisecond * 50)
	return "What programming languages do you prefer?"
}

func fetchPostLikes(_ string) int {
	time.Sleep(time.Millisecond * 50)
	return 10
}

func fetchPostComments(_ string) []string {
	time.Sleep(time.Millisecond * 100)
	return []string{"Golang", "Java", "Rust"}
}