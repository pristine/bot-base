package profile

import "github.com/iancoleman/orderedmap"

type Profile struct {
	Name           string `json:"name"`
	ProfileGroup   string `json:"profileGroup"`
	BillingAddress struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Line1    string `json:"line1"`
		Line2    string `json:"line2"`
		Line3    string `json:"line3"`
		Postcode string `json:"postCode"`
		City     string `json:"city"`
		Country  string `json:"country"`
		State    string `json:"state"`
	} `json:"billingAddress"`
	ShippingAddress struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Line1    string `json:"line1"`
		Line2    string `json:"line2"`
		Line3    string `json:"line3"`
		Postcode string `json:"postCode"`
		City     string `json:"city"`
		Country  string `json:"country"`
		State    string `json:"state"`
	} `json:"shippingAddress"`
	PaymentDetails struct {
		NameOnCard   string `json:"nameOnCard"`
		CardType     string `json:"cardType"`
		CardNumber   string `json:"cardNumber"`
		CardExpMonth string `json:"cardExpMonth"`
		CardExpYear  string `json:"cardExpYear"`
		CardCvv      string `json:"cardCvv"`
	} `json:"paymentDetails"`
	SameBillingAndShipping bool `json:"sameBillingAndShippingAddress"`
}

type ProfileGroup struct {
	ID       string                 `json:"id"`
	Name     string                 `json:"name"`
	Profiles *orderedmap.OrderedMap `json:"profiles"` // ordered map to make sure our profile selection works
}