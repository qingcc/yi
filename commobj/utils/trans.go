package trans_service

type Args struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Query string `json:"query"`
	Ssl   bool   `json:"ssl"`
}

type Reply struct {
	Query string `json:"query"`
	BaseServiceResult
}

type BaseServiceResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}
