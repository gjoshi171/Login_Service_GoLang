package models

type Response struct {
	Results *Result `json:"result,omitempty"`
	Message string  `json:"message,omitempty"`
}

func (r Response) GetResult() *Result {
	return r.Results
}

func (r *Response) SetResult(Result *Result) {
	r.Results = Result
}

func (r Response) GetMessage() string {
	return r.Message
}

func (r *Response) SetMessage(Message string) {
	r.Message = Message
}
