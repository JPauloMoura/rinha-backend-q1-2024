package date

import (
	"fmt"
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

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Println("==> ", name, elapsed.Milliseconds())
}
