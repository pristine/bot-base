package footsites

import (
	"fmt"
	"github.com/EdwinJ0124/bot-base/internal/task"
)

func submitOrder(t *task.Task) task.TaskState {
	internal := t.Internal.(*footsites)

	requestBody := OrderRequest{
		Cartid: "",
		Deviceid: "",
		Encryptedcardnumber: "",
		Encryptedexpirymonth: "",
		Encryptedexpiryyear: "",
		Encryptedsecuritycode: "",
		Paymentmethod: "CREDITCARD",
		Preferredlanguage: "en",
		Returnurl: fmt.Sprintf("https://%s/adyen/checkout", internal.Host),
		Termsandcondition: false,
	}

	_, err := t.Client.NewRequest().
		SetURL(fmt.Sprintf("https://%s/api/v2/users/orders", internal.Host)).
		SetMethod("POST").
		SetHeader("user-agent", userAgent).
		SetHeader("accept", "application/json").
		SetHeader("content-type", "application/json").
		SetJSONBody(requestBody).
		Do()

	if err != nil {
		// handle error and retry
		return SUBMIT_ORDER
	}

	return handleSubmitOrderResponse(t)
}

func handleSubmitOrderResponse(t *task.Task) task.TaskState {
	if t.Client.LatestResponse.StatusCode() > 201 {
		// message := HandleStatusCodes(resp.StatusCode())

		// handle error and retry
		return SUBMIT_ORDER
	}

	return task.DoneTaskState
}