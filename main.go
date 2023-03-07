package main

import "github.com/choisangh/image_crud_api/pkg/router"

func main() {
	r := router.Router()

	r.Run(":8080")

}
