package news

import (
	"finalproject-BE/business/news"
	"finalproject-BE/controllers"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type NewsController struct {
	NewsRepo news.Repository
	
}

func NewNewsController(newsRepo news.Repository) *NewsController {
	return &NewsController{
		NewsRepo: newsRepo,
	}
}

func (newsController NewsController) GetByCategory(c echo.Context) error {
	fmt.Println("GetNews")
	category := c.QueryParam("category")
	ctx := c.Request().Context()
	data, error := newsController.NewsRepo.GetByCategory(ctx, category)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, FromDomain(data))
}
