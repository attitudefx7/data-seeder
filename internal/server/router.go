package server

import (
	"data-seeder/internal/api/sd"
	"fmt"
	"net/http"
)

func registerRouter(router *http.ServeMux)  {
	// test
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("test")
	})

	// sd-seeder
	invoke := new(sd.Invoke)
	router.HandleFunc("/sd_seeder", invoke.Invoke)

}