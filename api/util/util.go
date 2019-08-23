package util

type Article struct {
	ID         int         `json:"id"`
	UUID       string      `json:"uuid"`
	TITLE      string      `json:"title"`
	CONTENT    string      `json:"content"`
	IMAGENAMES []ImageName `json:"imageNames"`
}

type ImageName struct {
	NAME string `json:"name"`
}

type ImageData struct {
	ARTICLEUUID string      `json:"articleUUID"`
	IMAGENAMES  []ImageName `json:"imageNames"`
}
