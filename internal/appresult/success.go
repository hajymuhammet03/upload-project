package appresult

import "encoding/json"

var (
	Success = NewAppSuccess("Maladess!!!", "SS-10000", nil)
)

type AppSuccess struct {
	Data interface{} `json:"data"`
}

func (s *AppSuccess) Error() string {
	return ""
}

func (s *AppSuccess) Marshal() []byte {
	marshal, err := json.Marshal(s)
	if err != nil {
		return nil
	}

	return marshal
}

func NewAppSuccess(message, code string, data interface{}) *AppSuccess {
	return &AppSuccess{
		Data: data,
	}
}
