package email

import "fmt"

func Send(contents string, to string) {

}

func Contents(title string, name string, nights uint) string {
	text := "Dear %s %s,\n your room reservation for %d night(s) is confirmed. Have a nice day !"
	return fmt.Sprintf(text,
		title,
		name,
		nights)

}
