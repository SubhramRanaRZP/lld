package call_services

import (
	"fmt"
)

// NormalCall contains external call service
type NormalCall struct {
	// here contains external call service whose method will be consumed by this struct
	phoneNumbers []string
}

func (call *NormalCall) Call() {
	var calledNumber string
	for _, num := range call.phoneNumbers {
		// try to call the 'num'...
		// if it is current call is successful, then break
		calledNumber = num
	}
	fmt.Printf("calling %v", calledNumber)
}

func (call *NormalCall) AddEmergencyPhoneNumber(num string) {
	call.phoneNumbers = append(call.phoneNumbers, num)
}
