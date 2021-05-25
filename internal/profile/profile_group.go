package profile

import "errors"

type ProfileGroup struct {
	Name, ID string
	Profiles map[string]bool
}

var (
	ProfileGroupDoesNotExistErr = errors.New("profile group does not exist")
	profileGroups = make(map[string]*Profile)
)