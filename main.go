package main

import "os"

func main() {
	sender := NewSender(os.Getenv("_email"), os.Getenv("_password"))

	//The receiver needs to be in slice as the receive supports multiple receiver
	Receiver := []string{"rabiul@gmail.com", "rabiulislam993@live.com"}

	Subject := "Testing HTLML Email from golang"
	message := `
	<!DOCTYPE HTML PULBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
	<html>
	<head>
	<meta http-equiv="content-type" content="text/html"; charset=ISO-8859-1">
	</head>
	<body>This is the body<br>
	<div class="moz-signature"><i><br>
	<br>
	Regards<br>
	Rabiul Islam<br>
	<i></div>
	</body>
	</html>
	`
	bodyMessage := sender.WriteHTMLEmail(Receiver, Subject, message)

	sender.SendMail(Receiver, Subject, bodyMessage)
}
