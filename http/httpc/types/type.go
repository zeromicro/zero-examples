package types

type Request struct {
	Node   string `path:"node"`
	ID     int    `form:"id"`
	Header string `header:"X-Header"`
	Body   string `json:"body"`
}
