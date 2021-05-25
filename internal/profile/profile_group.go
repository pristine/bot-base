package profile

import (
	"errors"
	"github.com/iancoleman/orderedmap"
	"github.com/lithammer/shortuuid"
)

type ProfileGroup struct {
	Name, ID string
	Profiles *orderedmap.OrderedMap // ordered map to make sure our profile selection works
}

var (
	ProfileGroupEmptyErr = errors.New("profile group does not contain any profiles")
	ProfileGroupDoesNotExistErr = errors.New("profile group does not exist")
	profileGroups = make(map[string]*ProfileGroup)
)

// DoesProfileGroupExist checks if a profile group exists
func DoesProfileGroupExist(id string) bool {
	_, ok := profileGroups[id]
	return ok
}

// CreateProfileGroup creates a new profile group
func CreateProfileGroup(name string) string {
	id := shortuuid.New()

	profileGroups[id] = &ProfileGroup{
		Name: name,
		ID: id,
		Profiles: orderedmap.New(),
	}

	return id
}

// RemoveProfileGroup removes a profile group
func RemoveProfileGroup(id string) error {
	if !DoesProfileGroupExist(id) {
		return ProfileGroupDoesNotExistErr
	}

	delete(profileGroups, id)

	return nil
}

// GetProfileGroupById gets a profile group by id
func GetProfileGroupById(id string) (*ProfileGroup, error) {
	if !DoesProfileGroupExist(id) {
		return &ProfileGroup{}, ProfileGroupDoesNotExistErr
	}

	return profileGroups[id], nil
}

// GetProfileFromProfileGroup gets a profile from a profile group
// Gets the first profile, then moves it to the back
func GetProfileFromProfileGroup(id string) (*Profile, error) {
	if !DoesProfileGroupExist(id) {
		return &Profile{}, ProfileGroupDoesNotExistErr
	}

	profileGroup := profileGroups[id]

	profileIds := profileGroup.Profiles.Keys()

	if len(profileIds) == 0 {
		return &Profile{}, ProfileGroupEmptyErr
	}

	firstProfileId := profileIds[0]

	profile, err := GetProfileById(firstProfileId)

	if err != nil {
		return &Profile{}, err
	}

	profileGroup.Profiles.Delete(firstProfileId)

	profileGroup.Profiles.Set(firstProfileId, true)

	return profile, nil
}