package api

type URIBookID struct {
	ID *int64 `uri:"id" binding:"required"`
}
