package footsites

import (
	"fmt"
	"github.com/EdwinJ0124/bot-base/internal/task"
	"github.com/EdwinJ0124/bot-base/internal/utils"
	"strings"
)

func submitShipping(t *task.Task) task.TaskState {
	internal := t.Internal.(*footsites)

	requestBody := ShippingRequest{
		Shippingaddress: ShippingAddress{
			Loqatesearch: "",
			Country: Country{
				Isocode: utils.CountryToISO[internal.Profile.ShippingAddress.Country],
				Name:    internal.Profile.ShippingAddress.Country,
			},
			Email:      false,
			Firstname:  strings.Split(internal.Profile.ShippingAddress.Name, " ")[0],
			ID:         nil,
			Lastname:   strings.Split(internal.Profile.ShippingAddress.Name, " ")[1],
			Line1:      internal.Profile.ShippingAddress.Line1,
			Line2:      internal.Profile.ShippingAddress.Line2,
			Phone:      internal.Profile.ShippingAddress.Phone,
			Postalcode: internal.Profile.ShippingAddress.Postcode,
			Recordtype: "S",
			Region: Region{
				Countryiso:   utils.CountryToISO[internal.Profile.ShippingAddress.Country],
				Isocode:      fmt.Sprintf("%s:%s", utils.CountryToISO[internal.Profile.ShippingAddress.Country], utils.StateToISO[internal.Profile.ShippingAddress.State]),
				Isocodeshort: utils.StateToISO[internal.Profile.ShippingAddress.State],
				Name:         internal.Profile.ShippingAddress.State,
			},
			Regionfpo:            nil,
			Saveinaddressbook:    false,
			Setasbilling:         true,
			Setasdefaultshipping: false,
			Setasdefaultbilling:  false,
			Shippingaddress:      true,
			Town:                 internal.Profile.ShippingAddress.City,
			Type:                 "default",
		},
	}

	_, err := t.Client.NewRequest().
		SetURL(fmt.Sprintf("https://%s/api/users/carts/current/addresses/shipping", internal.Host)).
		SetMethod("POST").
		SetHeader("user-agent", userAgent).
		SetHeader("accept", "application/json").
		SetHeader("content-type", "application/json").
		SetJSONBody(requestBody).
		Do()

	if err != nil {
		// handle error and retry
		return SUBMIT_SHIPPING
	}

	return handleSubmitShippingResponse(t)
}

func handleSubmitShippingResponse(t *task.Task) task.TaskState {
	if t.Client.LatestResponse.StatusCode() > 201 {
		// message := HandleStatusCodes(resp.StatusCode())

		// handle error and retry
		return SUBMIT_SHIPPING
	}

	return SUBMIT_BILLING
}
