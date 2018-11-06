package main

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/basicauth"
	"os"
	"time"
)

func newApp() *iris.Application {

	app := iris.New()

	authConfig := basicauth.Config{
		Users:   map[string]string{"wtq": "wtq", "www": "www"},
		Realm:   "Authorization Required", // defaults to "Authorization Required"
		Expires: time.Duration(30) * time.Minute,
	}

	authentication := basicauth.New(authConfig)

	app.RegisterView(iris.HTML("./views", ".html"))

	app.Get("/", func(ctx iris.Context) { ctx.Redirect("/admin") })
	// to party

	needAuth := app.Party("/admin", authentication)
	{
		//http://localhost:8080/admin
		needAuth.Get("/", func(ctx iris.Context) {
			//ctx.ViewData("Name", "iris") // the .Name inside the ./templates/hi.html
			ctx.Gzip(true)               // enable gzip for big files
			ctx.View("onepage.html")
		})

		needAuth.Get("/data", func(ctx iris.Context) {
			c, err := redis.Dial("tcp", "localhost:6379")
			errCheck(err)

			defer c.Close()
			// get a specific key, as string, if no found returns just an empty string
			I0T, getErr := redis.String(c.Do("get", "I0T"))
			errCheck(getErr)
			I0H, getErr := redis.String(c.Do("get", "I0H"))
			errCheck(getErr)
			mapD := map[string]map[string]string{"I0": {"Temperature": I0T, "Humidity": I0H}}
			mapB, _ := json.Marshal(mapD)

			ctx.Writef(string(mapB))
		})
	}

	return app
}

func main() {

	app := newApp()
	app.StaticWeb("/public", "./public")
	app.Run(iris.Addr(":8080"))
}

func errCheck(err error) {
	if err != nil {
		fmt.Println("sorry,has some error:", err)
		os.Exit(-1)
	}
}
