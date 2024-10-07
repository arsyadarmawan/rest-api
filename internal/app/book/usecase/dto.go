package usecase

type BookResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type BookRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
