package request

import "time"

type NewsRequest struct {
	Query string    `form:"q" binding:"required,min=3"`
	From  time.Time `form:"from" binding:"required" time_format:"2006-01-02"`
}
