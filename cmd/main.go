package main

import (
	"flag"
	_ "homework/docs"
	"homework/infra"
	"homework/routes"
	"log"
)

// @title Project Chat App
// @version 1.0
// @description Homework CHAPTER 50 -55
// @termsOfService http://example.com/terms/
// @contact.name Paulus Otto Harman
// @contact.url https://academy.lumoshive.com/contact-us
// @contact.email paulus.otto.harman@gmail.com
// @license.name Lumoshive Academy
// @license.url https://academy.lumoshive.com
// @host localhost:8080
// @schemes http
// @BasePath /
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description <h5>Please type "Bearer", following with a space, and then the JWT token</h5>

func main() {
	ctx, err := infra.NewServiceContext()
	if err != nil {
		log.Fatal("can't init service context %w", err)
	}

	if shouldNotLaunchServer() {
		return
	}

	routes.NewRoutes(*ctx)
}

func shouldNotLaunchServer() bool {
	shouldNotLaunch := false

	flag.Parse()
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "m" {
			shouldNotLaunch = true
		}

		if f.Name == "s" {
			shouldNotLaunch = true
		}
	})

	return shouldNotLaunch
}
