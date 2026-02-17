package domain

type Payload struct {
	FullName string
	Date     int
}

func NewPayload(fullName string, datetime int) Payload {
	return Payload{FullName: fullName, Date: datetime}
}
