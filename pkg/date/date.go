package date

import (
	"log"
	"time"
	_ "time/tzdata"
)

var DATE_BR_WITH_HOURS = "02/01/2006 03:04:05.00"

func LocationBR() *time.Location {
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		log.Fatal(err)
	}

	return loc
}
