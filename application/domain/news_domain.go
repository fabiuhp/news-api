package domain

type NewsReqDomain struct {
	Query string
	From  string
}

type NewsDomain struct {
	Status       string
	TotalResults string
	Articles     []Article
}

type Article struct {
	Source      string
	Id          string
	Name        string
	Author      string
	Title       string
	Description string
	Url         string
	PublishedAt string
	UrlToImage  string
	Content     string
}
