package gin

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetPagePerFromQs retrieves the page number and the number of items per page from query string in request.
// The prerequisite for using this function is that the parameters' name of the page number and the number of items
// per page in query string must be `page` and `per`, like `example.com?page=2&per=10`.
//
// If the page parameter in query string is not a number or less than 1, the returned page will be 1. If the per
// parameter in query string is not a number, the returned per will be defaultPer, and if the requested per is greater
// than maxPer, the returned per will be maxPer.
func GetPagePerFromQs(c *gin.Context, defaultPer, maxPer int) (page, per int) {
	pageStr := c.Query("page")
	perStr := c.Query("per")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	per, err = strconv.Atoi(perStr)
	if err != nil {
		per = defaultPer
	}
	if per > maxPer {
		per = maxPer
	}

	return
}
