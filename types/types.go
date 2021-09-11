package types

// Struct for Teams data details
type Teams struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Division string `json:"division"`
}

// Struct for Results response API
type Results struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// Struct for list endpoint Response API
type Endpoints struct {
	GetAll    string `json:"get_all"`
	GetById   string `json:"get_by_id"`
	SaveTeams string `json:"save_teams"`
}