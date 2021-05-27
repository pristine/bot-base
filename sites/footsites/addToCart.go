package footsites

import (
	"fmt"
	"github.com/EdwinJ0124/bot-base/internal/task"
)

func addToCart(t *task.Task, internal *FootsitesInternal) task.TaskState {
	requestBody := AddToCartRequest{
		ProductQuantity: 1,
		ProductID:       internal.VariantID,
	}

	_, err := t.Client.NewRequest().
		SetURL(fmt.Sprintf("https://%s/api/users/carts/current/entries", internal.Host)).
		SetMethod("POST").
		SetHeader("user-agent", userAgent).
		SetHeader("accept", "application/json").
		SetHeader("content-type", "application/json").
		SetJSONBody(requestBody).
		Do()

	if err != nil {
		// handle error and retry
		return ADD_TO_CART
	}

	return handleAddToCartResponse(t)
}

func handleAddToCartResponse(t *task.Task) task.TaskState {
	if t.Client.LatestResponse.StatusCode() > 201 {
		// message := HandleStatusCodes(resp.StatusCode())

		// handle error and retry
		return ADD_TO_CART
	}

	return VERIFY_EMAIL
}
