package models

type Payload struct {
	FullName string
	Date     string
}

func NewPayload(FullName, Date string) Payload {
	return Payload{FullName: FullName, Date: Date}
}