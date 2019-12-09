package models

const (
	Region      = "ap-southeast-1"
	CredProfile = "default"
	Url         = "https://sqs.ap-southeast-1.amazonaws.com/296950926313/testaja.fifo"
)

type Messages struct {
	Data int `json:"data"`
}

func GetData() []Messages {
	data := []Messages{
		Messages{
			Data: 1,
		},
		Messages{
			Data: 2,
		},
		Messages{
			Data: 3,
		},
		Messages{
			Data: 4,
		},
		Messages{
			Data: 5,
		},
		Messages{
			Data: 6,
		},
		Messages{
			Data: 7,
		},
		Messages{
			Data: 8,
		},
		Messages{
			Data: 9,
		},
		Messages{
			Data: 10,
		},
	}
	return data
}
