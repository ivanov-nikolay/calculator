package model

// Request входящее сообщение, содержащее арифметическое вырежение
type Request struct {
	Expression string `json:"expression"`
}

// Response исходящее сообщение, содержащее результат вычисления
type Response struct {
	Result string `json:"result"`
}
