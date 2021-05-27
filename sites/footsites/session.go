package footsites

import (
	"fmt"
	"github.com/EdwinJ0124/bot-base/internal/task"
)

func getSession(t *task.Task, internal *FootsitesInternal) task.TaskState {
	_, err := t.Client.NewRequest().
		SetURL(fmt.Sprintf("https://%s/api/session", internal.Host)).
		SetMethod("GET").
		SetHeader("user-agent", userAgent).
		SetHeader("accept", "application/json").
		Do()

	if err != nil {
		// handle error and retry
		return GET_SESSION
	}

	return handleSessionResponse(t, internal)
}

func handleSessionResponse(t *task.Task, internal *FootsitesInternal) task.TaskState {
	if t.Client.LatestResponse.StatusCode() > 201 {
		// message := HandleStatusCodes(resp.StatusCode())

		// handle error and retry
		return GET_SESSION
	}

	sessionResponse := SessionResponse{}

	err := t.Client.LatestResponse.BodyAsJSON(&sessionResponse)

	if err != nil {
		// handle error and retry
		return GET_SESSION
	}

	internal.CSRFToken = sessionResponse.Data.CSRFToken

	return WAIT_FOR_MONITOR
}
