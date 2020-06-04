package model

type Note struct {
	Model
	Type string `json:"type"`
	Content string `json:"content"`
}

