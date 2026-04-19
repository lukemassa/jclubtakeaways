package main

import (
	"fmt"
	"time"

	"github.com/jessevdk/go-flags"
	log "github.com/lukemassa/clilog"
	"github.com/lukemassa/jclubtakeaways/internal/schedule"
)

func parseDate(dateFmt string) (time.Time, error) {
	if dateFmt == "" {
		return time.Now(), nil
	}
	return time.Parse("2006-01-02", dateFmt)
}

func main() {
	var opts struct {
		// Date
		Date string `short:"d" long:"date" description:"A file"`
	}

	_, err := flags.Parse(&opts)

	if err != nil {
		log.Error(err)
	}
	date, err := parseDate(opts.Date)
	if err != nil {
		log.Errorf("Cannot parse date: %v", err)
	}

	s := schedule.New(date)
	deadlines := s.Deadlines()
	fmt.Println("Current month")
	for _, deadline := range deadlines.ThisMonth {
		fmt.Printf("  %s: %s\n", deadline.Date.Format("2006-01-02 Mon"), deadline.Description)
	}

	fmt.Println("\nNext month")
	for _, deadline := range deadlines.NextMonth {
		fmt.Printf("  %s: %s\n", deadline.Date.Format("2006-01-02 Mon"), deadline.Description)
	}

}
