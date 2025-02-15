package meets_types

type Meet struct {
	Meet_id        string `json:"meet_id"`
	Creator_id     string `json:"creator_id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Location       string `json:"location"`
	Meet_date      string `json:"meet_date"`
	Theme          string `json:"theme"`
	Meet_banner    string `json:"meet_banner"`
	Meet_thumbnail string `json:"meet_thumbnail"`
}

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type Location struct {
}
