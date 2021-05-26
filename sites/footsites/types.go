package footsites

import "github.com/EdwinJ0124/bot-base/internal/profile"

type SessionResponse struct {
	Data struct {
		CSRFToken string `json:"csrfToken"`
	} `json:"data"`
}

type AddToCartRequest struct {
	ProductQuantity int    `json:"productQuantity"`
	ProductID       string `json:"productId"`
}

type ShippingCountry struct {
	Isocode string `json:"isocode"`
	Name    string `json:"name"`
}

type ShippingRegion struct {
	Countryiso   string `json:"countryIso"`
	Isocode      string `json:"isocode"`
	Isocodeshort string `json:"isocodeShort"`
	Name         string `json:"name"`
}

type ShippingAddress struct {
	Setasdefaultbilling  bool   `json:"setAsDefaultBilling"`
	Setasdefaultshipping bool   `json:"setAsDefaultShipping"`
	Firstname            string `json:"firstName"`
	Lastname             string `json:"lastName"`
	Email                bool   `json:"email"`
	Phone                string `json:"phone"`
	Country              ShippingCountry `json:"country"`
	ID                interface{} `json:"id"`
	Setasbilling      bool        `json:"setAsBilling"`
	Saveinaddressbook bool        `json:"saveInAddressBook"`
	Region           ShippingRegion `json:"region"`
	Type            string      `json:"type"`
	Loqatesearch    string      `json:"LoqateSearch"`
	Line1           string      `json:"line1"`
	Line2           string      `json:"line2"`
	Postalcode      string      `json:"postalCode"`
	Town            string      `json:"town"`
	Regionfpo       interface{} `json:"regionFPO"`
	Shippingaddress bool        `json:"shippingAddress"`
	Recordtype      string      `json:"recordType"`
}

type ShippingRequest struct {
	Shippingaddress ShippingAddress `json:"shippingAddress"`
}

type footsites struct {
	// internal things
	Host string

	CSRFToken string
	VariantID string

	ProfileRetrieved bool
	Profile *profile.Profile
}
