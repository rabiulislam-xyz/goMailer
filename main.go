package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "os"

	_ "github.com/gin-gonic/gin"
)

//func main() {
//	email := os.Getenv("email")
//	password := os.Getenv("password")
//	port := os.Getenv("PORT")
//
//	if email == "" {
//		log.Fatal("$EMAIL must be set")
//	}
//	if password == "" {
//		log.Fatal("$PASSWORD must be set")
//	}
//	if port == "" {
//		port = "8000"
//	}
//
//
//	//The receiver needs to be in slice as the receive supports multiple receiver
//	Receiver := []string{"rabiul@aadibd.com"}
//	Subject := "Testing HTLML Email from golang"
//	message := `
//	<!DOCTYPE HTML PULBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
//	<html>
//	<head>
//	<meta http-equiv="content-type" content="text/html"; charset=ISO-8859-1">
//	</head>
//	<body>This is the body<br>
//	<div class="moz-signature"><i><br>
//	<br>
//	Regards<br>
//	Rabiul Islam<br>
//	<i></div>
//	</body>
//	</html>
//	`
//
//	router := gin.New()
//	router.Use(gin.Logger())
//	//router.LoadHTMLGlob("templates/*.tmpl.html")
//	//router.Static("/static", "static")
//
//	router.GET("/", func(c *gin.Context) {
//		c.String(http.StatusOK, string(`
//										for sending email, POST request to current URL with json like:
//
//										{
//											"emails": ["mail1@go.com", "mail2@go.com"],
//											"subject": "mail subject",
//											"body": "mail body, can be html",
//											"token": "v@rificationtoken"
//										}
//
//										`))
//	})
//
//	router.POST("/", func(c *gin.Context) {
//		c.JSON(http.StatusOK, string(`{"status": "success", "msg": "mail sent"`))
//	})
//
//	router.GET("/send", func(c *gin.Context) {
//		sender := NewSender(email, password)
//		bodyMessage := sender.WriteHTMLEmail(Receiver, Subject, message)
//		sender.SendMail(Receiver, Subject, bodyMessage)
//
//		c.String(http.StatusOK, string("email sent"))
//	})
//
//	router.Run(":" + port)
//}


type data struct {
	Emails *[]string `json:"emails"`
	Subject *string    `json:"subject"`
	Body  *string    `json:"body"`
	Token  *string    `json:"token"`
}

func index(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		log.Println("get req")
		text := `
			for sending email, POST request to current URL with json like:
			
			{
				"emails": ["mail1@go.com", "mail2@go.com"],
				"subject": "mail subject",
				"body": "mail body, can be html",
				"token": "v@rificationtoken"
			}
			
			`
		w.Write([]byte(text))

	case "POST":
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		var d data
		err := decoder.Decode(&d)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(struct {Message, Status string}{"Mail sent","success"})

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %v is not allowd", r.Method)
	}
}

func main() {
	http.HandleFunc("/", index)

	log.Println("Go!")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
