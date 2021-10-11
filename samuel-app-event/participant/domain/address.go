package domain

type Address struct {
	Street string `bson:"street" json:"street"`
	Number int `bson:"number" json:"number"`
	ZipCode string `bson:"zipCode" json:"zipCode"`
	City    string `bson:"city" json:"city"`
}

func NewAddress(street string, number int, zipCode string, city string) *Address {
	return &Address{Street: street, Number: number, ZipCode: zipCode, City: city}
}