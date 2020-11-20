package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv01/service"
	"net/http"
	"strconv"
)

type ArticleController struct{}

func NewArticleController() ArticleController {
	return ArticleController{}
}
//得到一篇文章的详情
func (a ArticleController) GetOne(c *gin.Context) {

	id := c.Params.ByName("id")
	fmt.Println("id:"+id);

	articleId,err := strconv.ParseInt(id, 10, 64);
    if (err != nil) {
		c.AbortWithStatus(400)
		fmt.Println(err.Error())
	}

	articleOne,err := service.GetOneArticle(articleId);
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err.Error())
	} else {
		c.JSON(http.StatusOK, &articleOne)
	}
	return
}

//得到多篇文章，按分页返回
func (a ArticleController) GetList(c *gin.Context) {

	page := c.DefaultQuery("page", "0")
	pageInt, err := strconv.Atoi(page)
	if (err != nil) {
		c.AbortWithStatus(400)
		fmt.Println(err.Error())
	}
	pageSize := 2;
	pageOffset := pageInt * pageSize

	articles,err := service.GetArticleList(pageOffset,pageSize)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err.Error())
	} else {
		c.JSON(http.StatusOK, &articles)
	}
	return
}

