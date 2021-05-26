package footsites

import (
	"fmt"
	"github.com/EdwinJ0124/bot-base/internal/profile"
	"github.com/EdwinJ0124/bot-base/internal/task"
)

func verifyEmail(t *task.Task) task.TaskState {
	internal := t.Internal.(*footsites)

	// get profile data
	if !internal.ProfileRetrieved {

		// profile is in a queue system
		profile, err := profile.GetProfileFromProfileGroup(t.ProfileGroupID)

		if err != nil {
			// handle error and retry
			return VERIFY_EMAIL
		}

		internal.Profile = profile
		internal.ProfileRetrieved = true
	}

	_, err := t.Client.NewRequest().
		SetURL(fmt.Sprintf("https://%s/api/users/carts/current/email/%s", internal.Host, internal.Profile.ShippingAddress.Email)).
		SetMethod("PUT").
		SetHeader("user-agent", userAgent).
		SetHeader("accept", "application/json").
		SetHeader("content-type", "application/json").
		Do()

	if err != nil {
		// handle error and retry
		return VERIFY_EMAIL
	}

	return handleVerifyEmailResponse(t)
}

func handleVerifyEmailResponse(t *task.Task) task.TaskState {
	if t.Client.LatestResponse.StatusCode() > 201 {
		// message := HandleStatusCodes(resp.StatusCode())

		// handle error and retry
		return VERIFY_EMAIL
	}

	return SUBMIT_SHIPPING
}