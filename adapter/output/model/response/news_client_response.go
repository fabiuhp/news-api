package response

type NewsClientResponse struct {
	Status       string            `json:"status"`
	TotalResults int               `json:"totalResults"`
	Articles     []ArticleResponse `json:"articles"`
}

type ArticleResponse struct {
	Source      SourceResponse `json:"source"`
	Id          string         `json:"id"`
	Name        string         `json:"name"`
	Author      string         `json:"author"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Url         string         `json:"url"`
	PublishedAt string         `json:"publishedAt"`
	UrlToImage  string         `json:"urlToImage"`
	Content     string         `json:"content"`
}

type SourceResponse struct {
	Id   *string `json:"id"`
	Name string  `json:"name"`
}
