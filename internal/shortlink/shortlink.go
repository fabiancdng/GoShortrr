package shortlink

type Shortlink struct {
	Id       int    `json:"id"`
	Link     string `json:"link"`
	Short    string `json:"short"`
	User     string `json:"user"`
	Password bool   `json:"password"`
	Created  int    `json:"created"`
}
