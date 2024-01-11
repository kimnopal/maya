package model

type WebResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// type Error struct {
// 	Field    string   `json:"field"`
// 	Messages []string `json:"messages"`
// }
