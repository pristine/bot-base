package footsites

import "github.com/EdwinJ0124/bot-base/internal/task"

var (
	INITIALIZE       task.TaskState = "initialize"
	GET_SESSION      task.TaskState = "get_session"
	WAIT_FOR_MONITOR task.TaskState = "wait_for_monitor"
	ADD_TO_CART      task.TaskState = "add_to_cart"
	VERIFY_EMAIL task.TaskState = "verify_email"
	SUBMIT_SHIPPING task.TaskState = "submit_shipping"
	SUBMIT_BILLING task.TaskState = "submit_billing"
	SUBMIT_ORDER task.TaskState = "submit_order"
)
