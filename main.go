package main

import (
	"config/env"
	"config/middlewares"
	"controller/auth/checkoutcontroller"
	"controller/auth/logincontroller"
	"controller/auth/signupcontroller"
	"controller/cart/cart-manager/createcartcontroller"
	"controller/cart/cart-manager/getcartcontroller"
	"controller/cart/cart-manager/updatecartcontroller"
	"databases"
	"os"
	"sync"
	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/cors"
	// "github.com/gin-gonic/gin"
	// "github.com/stripe/stripe-go"
	// "github.com/stripe/stripe-go/v72"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go databases.ConnectDB(&wg)
	wg.Wait()

	env.Loadenvfile()
	PORT := os.Getenv("PORT")
	router := gin.New()

	//middlewares
	router.Use(middlewares.Recovery(), middlewares.Logger())
	router.Use(middlewares.CORSMiddleware("*"))

	router.POST("/signup", signupcontroller.Signup)
	router.POST("/login", logincontroller.Login)
	router.POST("/create-payment-intent", checkoutcontroller.Checkout)
	router.POST("/addusercart", createcartcontroller.Addusercart)
	router.POST("/addcartitem", updatecartcontroller.Addcartitem)
	router.GET("/getusercart", getcartcontroller.Getcartdata)

	router.Run(PORT)

}
