package entity

type GetNews struct {
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
}
