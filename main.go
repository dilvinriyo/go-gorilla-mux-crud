package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Post struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author string `json:"author"`
}

var posts []Post

func main() {

	// add few default values to the list
	posts = append(posts, Post{Id: 1, Title: "Alchemist", Body: "alchemist body here", Author: "Dilvin Riyo"})
	posts = append(posts, Post{Id: 2, Title: "Ikigai", Body: "ikigai body here", Author: "Nikhil jackson"})

	// create router using gorilla mux
	router := mux.NewRouter()

	// list all routes
	router.HandleFunc("/posts", getAllPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", getPostById).Methods("GET")
	router.HandleFunc("/posts", createPost).Methods("POST")
	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	// run server
	http.ListenAndServe(":9000", router)
}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(posts)
}

func getPostById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)
	requestedPostId, _ := strconv.Atoi(params["id"])

	for _, item := range posts {
		if requestedPostId == item.Id {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/josn")
	var post Post
	json.NewDecoder(r.Body).Decode(&post)
	post.Id = rand.Intn(1000000) // create random integer number to use it as id
	posts = append(posts, post)
	json.NewEncoder(w).Encode(post)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)
	requestedPostId, _ := strconv.Atoi(params["id"])

	for index, item := range posts {
		if requestedPostId == item.Id {
			posts = append(posts[:index], posts[index+1:]...) // exclude current item and create new list
			var post Post
			json.NewDecoder(r.Body).Decode(&post)
			post.Id = rand.Intn(1000000)
			posts = append(posts, post) // add new item to the end of the list
			json.NewEncoder(w).Encode(post)
		}
	}
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Appplication/json")
	params := mux.Vars(r)
	requestedPostId, _ := strconv.Atoi(params["id"])

	for index, item := range posts {
		if item.Id == requestedPostId {
			posts = append(posts[:index], posts[index+1:]...) // exclude current item and create new list
			break
		}
	}
}
