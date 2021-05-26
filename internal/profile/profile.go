package profile

import (
	"errors"
	"github.com/lithammer/shortuuid"
)

var (
	ProfileDoesNotExistErr = errors.New("profile does not exist")
	ProfileNotAssignedErr  = errors.New("profile not assigned")
	profiles               = make(map[string]*Profile)
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
