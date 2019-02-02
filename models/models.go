package models

type Items struct{
	ResourceURI string `json:"resourceURI"`
	Name string `json:"name"`
	Type string `json:"type"` 
}
type Generic struct {
	Avaible int `json:"available"`
	Returned int `json:"returned"`
	CollectionURI string `json:"collectionURI"`
	Items []Items `json:"items"`
}
type Series struct {
	Generic
}
type Events struct {
	Generic
}

type Stories struct {
	Generic 
}
type Comics struct {
	Generic 
}

type Thumbnail struct {
	Path string `json:"path"`
	Extension string `json:"extension"`
}
type Urls struct {
	Type string `json:"type"`
	Url string  `json:"url"`
}
type Results struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Modified string `json:"modified"` 
	ResourceURI string `json:"resourceURI"`
	Urls 	[]Urls `json:"urls"`
	Thumbnail Thumbnail `json:"thumbnail"`
	Comics Comics `json:"comics"`
	Stories Stories `json:"stories"`
	Events Events `json:"events"`
	Series Series `json:"series"`
}
type Data struct {
	Offset int `json:"offset"`
	Limit int `json:"limit"`
	Total int `json:"total"`
	Count int `json:"count"`
	Results []Results `json:"results"`
}

type Feed struct {
	Data Data  `json:"data"`
}