package gin

import (
	"strconv"
)

//ServiceCheckHandle ..
func ServiceCheckHandle() HandlerFunc {
	return func(c *Context) {
		if true == c.IsInternalURL() {
			u := c.Query("user_id")
			if "" != u {
				if i, err := strconv.ParseUint(u, 10, 64); nil == err {
					c.UserID = i
				} else {
					return
				}
			}
		} else {
			token := c.Query("token")
			if "" != token {

			}
		}
	}
}
