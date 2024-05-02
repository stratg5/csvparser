package entities

type Address struct {
	Street       string `json:"street,omitempty"`
	City         string `json:"city,omitempty"`
	ZipCode      string `json:"zipcode,omitempty"`
	OriginString string
	Valid        bool
}
