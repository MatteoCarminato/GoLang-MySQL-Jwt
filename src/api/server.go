package api

import (
	"fmt"	"log"
)


func  Run()  {
	fmt.Printf("running ...")
	r := router.New();
	log.Fatal(httpListenAndServe(":3000", r))
}
