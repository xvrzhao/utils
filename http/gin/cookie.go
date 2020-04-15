package gin

import (
	"github.com/gin-gonic/gin"
	"math"
	"time"
)

// SetCookieGenerally set a cookie to gin response.
// The default attributes of the cookie:
//   path: "/"
//   domain: ""
//   secure: false
//   httpOnly: true
func SetCookieGenerally(c *gin.Context, name, value string, duration time.Duration) {
	c.SetCookie(name, value, int(math.Ceil(duration.Seconds())), "/", "", false, true)
}
