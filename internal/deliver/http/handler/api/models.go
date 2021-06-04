package api

type UriBookID struct {
	ID *int64 `uri:"id" binding:"required"`
}

type JSONBook struct {
	Title string `json:"title"`
	Desc  string `json:"description"`
}
