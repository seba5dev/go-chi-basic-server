package models

type Song struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Artist_id string `json:"artist"`
	Album     string `json:"album"`
	Year      int    `json:"year"`
	Genre     string `json:"genre"`
}

type Artist struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Bio       string `json:"bio"`
	Albums_id string `json:"albums"`
}

type Album struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Artist_id string `json:"artist"`
	Year      int    `json:"year"`
	Genre     string `json:"genre"`
}

type SongsDoc struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
	Year   int    `json:"year"`
	Genre  string `json:"genre"`
}
