package models

// response model info
// @Description Response data model
type Response struct {
	Status   string       `json:"status" example:"OK"`                               //request HTTP status
	Code     int          `json:"code" example:"200"`                                //request HTTP Code
	Messages string       `json:"messages,omitempty" example:"request successfully"` //request HTTP Message
	Data     *interface{} `json:"data,omitempty" `                                   //request interface data response
}
