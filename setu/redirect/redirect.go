package redirect

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Redirect(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<html>success</html>"))
}
