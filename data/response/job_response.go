package response

type JobResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Detail      string `json:"description"`
	Requirement string `json:"requirement"`
	PostedBy    int    `json:"posted_by"`
}
