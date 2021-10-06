package news

import (
	"context"
	"encoding/json"
	"errors"
	"finalproject-BE/business/news"
	"net/http"
)

type NewsApi struct {
	httpClient http.Client
}

func NewNewsApi() news.Repository {
	return &NewsApi{
		httpClient: http.Client{},
	}
}

func (new *NewsApi) GetByCategory(ctx context.Context, category string) (news.Domain, error) {
	url := "https://newsapi.org/v2/top-headlines?country=id&category=" + category 
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Api-Key", "5a507ea673d6400baa60f12cb228e35a")
	resp, err := new.httpClient.Do(req)
	
	if err != nil {
		return news.Domain{}, err
	}
	
	data := Response{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if data.TotalResults == 0 {
		return news.Domain{}, errors.New("Category is not found")
	}
	return data.ToDomain(), err
}