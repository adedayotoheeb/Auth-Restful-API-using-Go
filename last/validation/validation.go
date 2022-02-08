package validation

type ErrorMessage struct{
	Failed string
	Msg string
}

func MsgForTag(tag string) string {
    switch tag {
    case "required":
        return "This field is required"
    case "email":
        return "Invalid email"
    }
    return ""
}
