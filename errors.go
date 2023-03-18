package vonagesms

import "fmt"

var ERRORS = map[string]map[string]string{
	"0": {
		"meaning":      "Generic ",
		"descritption": "Generic Error",
	},
	"1": {
		"meaning":      "Throttled ",
		"descritption": "You are sending SMS faster than the account limit (see What is the Throughput Limit for Outbound SMS?).",
	},
	"2": {
		"meaning":      "Missing Parameters ",
		"descritption": "Your request is missing one of the required parameters: from, to, api_key, api_secret or text.",
	},
	"3": {
		"meaning":      "Invalid Parameters ",
		"descritption": "The value of one or more parameters is invalid.",
	},
	"4": {
		"meaning":      "Invalid Credentials ",
		"descritption": "Your API key and/or secret are incorrect, invalid or disabled.",
	},
	"5": {
		"meaning":      "Internal Error ",
		"descritption": "An error has occurred in the platform whilst processing this message.",
	},
	"6": {
		"meaning":      "Invalid Message ",
		"descritption": "The platform was unable to process this message, for example, an unrecognized number prefix.",
	},
	"7": {
		"meaning":      "Number Barred ",
		"descritption": "The number you are trying to send messages to is on the list of barred numbers.",
	},
	"8": {
		"meaning":      "Partner Account Barred ",
		"descritption": "Your Vonage account has been suspended. Contact support@nexmo.com.",
	},
	"9": {
		"meaning":      "Partner Quota Violation ",
		"descritption": "You do not have sufficient credit to send the message. Top-up and retry.",
	},
	"10": {
		"meaning":      "Too Many Existing Binds ",
		"descritption": "The number of simultaneous connections to the platform exceeds your account allocation.",
	},
	"11": {
		"meaning":      "Account Not Enabled For HTTP ",
		"descritption": "This account is not provisioned for the SMS API, you should use SMPP instead.",
	},
	"12": {
		"meaning":      "Message Too Long ",
		"descritption": "The message length exceeds the maximum allowed.",
	},
	"14": {
		"meaning":      "Invalid Signature ",
		"descritption": "The signature supplied could not be verified.",
	},
	"15": {
		"meaning":      "Invalid Sender Address ",
		"descritption": "You are using a non-authorized sender ID in the from field. This is most commonly in North America, where a Vonage long virtual number or short code is required.",
	},
	"22": {
		"meaning":      "Invalid Network Code ",
		"descritption": "The network code supplied was either not recognized, or does not match the country of the destination address.",
	},
	"23": {
		"meaning":      "Invalid Callback URL ",
		"descritption": "The callback URL supplied was either too long or contained illegal characters.",
	},
	"29": {
		"meaning":      "Non-Whitelisted Destination ",
		"descritption": "Your Vonage account is still in demo mode. While in demo mode you must add target numbers to your whitelisted destination list. Top-up your account to remove this limitation.",
	},
	"32": {
		"meaning":      "Signature And API Secret Disallowed ",
		"descritption": "A signed request may not also present an api_secret.",
	},
	"33": {
		"meaning":      "Number De-activated ",
		"descritption": "The number you are trying to send messages to is de-activated and may not receive them.",
	},
}

type ResponseError struct {
	Status      string
	Meaning     string
	Description string
}

func (r *ResponseError) Error() string {
	msg := fmt.Sprintf("\nStatus: %s\nMeaning:%s\nDescription:%s", r.Status, r.Meaning, r.Description)
	return msg
}

func NewError(errorCode string) *ResponseError {
	err := &ResponseError{}

	if _, ok := ERRORS[errorCode]; !ok {
		errorCode = "0"
	}
	errDetail := ERRORS[errorCode]

	err.Status = errorCode
	err.Meaning = errDetail["meaning"]
	err.Description = errDetail["descritption"]
	return err
}
