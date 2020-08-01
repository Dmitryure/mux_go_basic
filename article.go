package main

// Article data
type Article struct {
	ID      string `json:"id,omitempty"`
	Author  string `json:"author,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}
