package crossref

type response struct {
	Status  string          `json:"status"`
	Message ArticleMetadata `json:"message"`
}

type ArticleMetadata struct {
	Title          []string   `json:"title"`
	ContainerTitle []string   `json:"container-title"`
	Authors        []Author   `json:"author"`
	Volume         string     `json:"volume"`
	Issue          string     `json:"issue"`
	Page           string     `json:"page"`
	ArticleNumber  string     `json:"article-number"`
	Issued         IssuedDate `json:"issued"`
	DOI            string     `json:"doi"`
}

type Author struct {
	Family string `json:"family"`
	Given  string `json:"given"`
}

type IssuedDate struct {
	DateParts [][]int `json:"date-parts"`
}
