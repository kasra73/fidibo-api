package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kasra73/fidibo-api/pkg/app"
	"github.com/kasra73/fidibo-api/pkg/e"
	"github.com/kasra73/fidibo-api/pkg/logging"
	"github.com/kasra73/fidibo-api/service/search_service"
)

// @Summary		SearchBooks	search among books
// @Security	ApiKeyAuth
// @Accept		mpfd
// @Produce 	json
// @Param		keyword query string true "keyword to search books"
// @Success 	200 {object} app.Response
// @Success 	400 {object} app.Response
// @Failure 	500 {object} app.Response
// @Router 		/search/book [get]
func SearchBooks(c *gin.Context) {
	appG := app.Gin{C: c}

	keyword := c.Query("keyword")
	if keyword == "" {
		appG.Response(http.StatusBadRequest, e.ERROR_KEYWORD_EMPTY, nil)
		return
	}

	bookSearcher := search_service.GetBookSearcher()
	books, err := bookSearcher.Search(keyword)
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_KEYWORD_SEARCH_FAILED, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, books)
}
