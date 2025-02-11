package article_api

import (
	"web/models/res"
	"web/service/es_ser"
	"github.com/gin-gonic/gin"
)

type ESIDRequest struct {
  ID string `json:"id" form:"id" uri:"id"`
}

func (ArticleApi) ArticleDetailView(c *gin.Context) {
  var cr ESIDRequest
  err := c.ShouldBindUri(&cr)
  if err != nil {
    res.FailWithCode(res.ArgumentError, c)
    return
  }
  model, err := es_ser.CommDetail(cr.ID)
  if err != nil {
    res.FailWithMessage(err.Error(), c)
    return
  }
  res.OkWithData(model, c)
}
