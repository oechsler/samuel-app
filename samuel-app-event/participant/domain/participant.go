package domain

type Participant struct {
	Id string `bson:"_id" json:"id"`
	EventId string `bson:"eventId" json:"eventId"`
	Address *Address `bson:"address" json:"address"`
	CheckedIn bool `bson:"active" json:"checkedIn"`
}

func NewParticipant(id string, eventId string, address *Address) *Participant {
	return &Participant{Id: id, EventId: eventId, Address: address, CheckedIn: true}
}

func (c Participant) CheckOut() *Participant {
	c.CheckedIn = false

	return &c
}