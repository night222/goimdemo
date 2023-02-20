package main

import (
	"goimdemo/routine"

	"goimdemo/utils"
)

func main() {
	err := utils.Init()
	if err != nil {
		panic(err)
	}
	r := routine.Router()
	r.Run() //listen an serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
