package trans_service

type Args struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Query string `json:"query"`
	Ssl   bool   `json:"ssl"`
}

type Reply struct {
	Query string `json:"query"`
}
