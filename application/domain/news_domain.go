package domain

type NewsReqDomain struct {
	Query string
	From  string
}

type NewsDomain struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

type Article struct {
	Source      string `json:"source"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PublishedAt string `json:"publishedAt"`
	UrlToImage  string `json:"urlToImage"`
	Content     string `json:"content"`
}
