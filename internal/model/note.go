package model

// Note represents the note model
type Note struct {
	NoteID    string   `json:"NoteID"`
	CreatedAt string   `json:"CreatedAt"`
	Content   string   `json:"Content"`
	Tags      []string `json:"Tags"`
	Title     string   `json:"Title"`
	UpdatedAt string   `json:"UpdatedAt"`
	UserID    string   `json:"UserID"`
}
