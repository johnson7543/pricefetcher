package types

type PriceResponse struct {
	Ticker string  `json:"ticker"`
	Price  float64 `json:"price"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type TemplateResponse struct {
	UUID       string      `json:"uuid"`
	Timestamp  string      `json:"timestamp"`
	StatusCode int         `json:"status"`
	Data       interface{} `json:"data,omitempty"`
	Message    string      `json:"message,omitempty"`
}

type ContextKey string
