package main

import (
	"context"
	"goimdemo/routine"

	"goimdemo/utils"
)

func main() {
	err := utils.Init()
	if err != nil {
		panic(err)
	}
	ctx, can := context.WithCancel(context.Background())
	defer can()
	defer utils.Close()
	go utils.Subscribes(ctx, 10)
	r := routine.Router()
	r.Run() //listen an serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
