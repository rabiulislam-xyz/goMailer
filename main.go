package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	email := os.Getenv("email")
	password := os.Getenv("password")
	port := os.Getenv("PORT")

	if email == "" {
		log.Fatal("$EMAIL must be set")
	}
	if password == "" {
		log.Fatal("$PASSWORD must be set")
	}
	if port == "" {
		log.Fatal("$PORT must be set")
	}


	//The receiver needs to be in slice as the receive supports multiple receiver
	Receiver := []string{"rabiul@aadibd.com"}
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

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/send", func(c *gin.Context) {
		sender := NewSender(email, password)
		bodyMessage := sender.WriteHTMLEmail(Receiver, Subject, message)
		sender.SendMail(Receiver, Subject, bodyMessage)

		c.String(http.StatusOK, string("email sent"))
	})

	router.Run(":" + port)
}
