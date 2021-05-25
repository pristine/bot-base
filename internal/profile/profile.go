package profile

import (
	"errors"
	"github.com/lithammer/shortuuid"
)

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

var (
	ProfileDoesNotExistErr = errors.New("profile does not exist")
	ProfileNotAssignedErr = errors.New("profile not assigned")
	profiles = make(map[string]*Profile)
)

// DoesProfileExist checks if a profile exists
func DoesProfileExist(id string) bool {
	_, ok := profiles[id]
	return ok
}

// CreateProfile creates a new profile
func CreateProfile(profile *Profile) string {
	id := shortuuid.New()

	profiles[id] = profile

	return id
}

// RemoveProfile removes a profile
func RemoveProfile(id string) error {
	if !DoesProfileExist(id) {
		return ProfileGroupDoesNotExistErr
	}

	delete(profiles, id)

	return nil
}

// GetProfileById gets a profile by id
func GetProfileById(id string) (*Profile, error) {
	if !DoesProfileExist(id) {
		return &Profile{}, ProfileGroupDoesNotExistErr
	}

	return profiles[id], nil
}

// GetAllProfileIDs gets all profile ids
func GetAllProfileIDs() []string {
	ids := []string{}

	for id := range profiles {
		ids = append(ids, id)
	}

	return ids
}

// AssignProfileToProfileGroup assigns a profile to a profile group
func AssignProfileToProfileGroup(profileId, profileGroupId string) error {
	if !DoesProfileExist(profileId) {
		return ProfileDoesNotExistErr
	}

	if !DoesProfileGroupExist(profileGroupId) {
		return ProfileGroupDoesNotExistErr
	}

	profileGroups[profileGroupId].Profiles.Set(profileId, true)

	return nil
}

// RemoveProfileFromProfileGroup removes a profile from a profile group
func RemoveProfileFromProfileGroup(profileId, profileGroupId string) error {
	if !DoesProfileExist(profileId) {
		return ProfileDoesNotExistErr
	}

	if !DoesProfileGroupExist(profileGroupId) {
		return ProfileGroupDoesNotExistErr
	}

	profileGroups[profileGroupId].Profiles.Delete(profileId)

	return nil
}