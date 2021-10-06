package news

import (
	"finalproject-BE/business/news"
)

type NewsResponse struct {
	// Name        string `json:"name"`
	// Author      string `json:"author"`
	// Title       string `json:"title"`
	// Description string `json:"description"`
	// URL         string `json:"url"`
	Article interface{} `json:"news"`
}

func FromDomain(domain news.Domain) NewsResponse {
	return NewsResponse{
		// Name:        domain.Name,
		// Author:      domain.Author,
		// Title:       domain.Title,
		// Description: domain.Description,
		// URL:         domain.URL,
		Article: domain.Article,
	}
}

func FromListDomain(domain []news.Domain) []NewsResponse {
	var response []NewsResponse
	for _, value := range domain {
		response = append(response, FromDomain(value))
	}
	return response
}
