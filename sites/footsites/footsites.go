package footsites

import "github.com/EdwinJ0124/bot-base/internal/task"

var (
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36"
)

func Initialize() {
	taskType := task.RegisterTaskType("footsites")

	taskType.SetFirstHandlerState(INITIALIZE)

	taskType.AddHandlers(task.TaskHandlerMap{
		INITIALIZE:       initialize,
		GET_SESSION:      getSession,
		WAIT_FOR_MONITOR: waitForMonitor,
		ADD_TO_CART:      addToCart,
		VERIFY_EMAIL: verifyEmail,
		SUBMIT_SHIPPING: submitShipping,
		SUBMIT_BILLING: submitBilling,
		SUBMIT_ORDER: submitOrder,
	})
}
