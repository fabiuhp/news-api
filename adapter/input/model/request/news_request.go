package request

type NewsRequest struct {
	Query string `form:"q" binding:"required,min=3"`
	From  string `form:"from" binding:"required,datetime" time_format:"2006-01-02"`
}
