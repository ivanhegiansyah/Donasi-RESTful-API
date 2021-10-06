package news

import (
	"context"
	"encoding/json"
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

func (new *NewsApi) GetByName(ctx context.Context, name string) (news.Domain, error) {
	req, _ := http.NewRequest("GET","https://newsapi.org/v2/top-headlines?country=id&category=health&apiKey=5a507ea673d6400baa60f12cb228e35a", nil)
	req.Header.Set("X-Api-Key", "5a507ea673d6400baa60f12cb228e35a")
	resp, err := new.httpClient.Do(req)
	
	if err != nil {
		return news.Domain{}, err
	}
	
	data := Response{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return news.Domain{}, err
	}
	return data.ToDomain(), err
}