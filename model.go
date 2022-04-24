package main

type Playlist struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Downloaded bool `json:"downloaded"`
}