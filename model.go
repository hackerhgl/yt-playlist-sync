package main

type SyncPlaylistItem struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Downloaded bool   `json:"downloaded"`
}
