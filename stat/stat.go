package stat

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ZacharyCalvert/medman/args"
	"github.com/ZacharyCalvert/medman/db"
)

func RunCommand() {
	managed := args.GetManagedFolder()
	if managed == "" {
		managed = "." // current working directory by default
	}
	data, err := db.LoadDatabase(managed)
	if err != nil {
		fmt.Printf("Failed to load database from %s: %v", managed, err)
		os.Exit(1)
	}
	lastUpdated := time.Unix(data.LastUpdated/1000, 0)
	var total, notReviewed, reviewed, ignored uint
	var earliest, latest int64
	counts := make(map[string]int)
	for _, rec := range data.Media {
		total++
		if rec.IsReviewed() {
			reviewed++
		} else {
			notReviewed++
		}
		if rec.IsIgnoredMedia() {
			ignored++
		}
		if earliest == 0 {
			earliest = rec.Earliest
		} else if earliest > rec.Earliest {
			earliest = rec.Earliest
		}
		if latest == 0 {
			latest = rec.Earliest
		} else if latest < rec.Earliest {
			latest = rec.Earliest
		}
		ext := strings.ToLower(rec.Extensions[0])
		if count, ok := counts[ext]; ok {
			counts[ext] = count + 1
		} else {
			counts[ext] = 1
		}
	}
	delim := strings.Repeat("*", 40)
	fmt.Printf(delim + "\n")
	fmt.Printf("The media directory at \n\t%s\n\t was last updated on %v.\n", managed, lastUpdated)
	fmt.Printf("Media includes content from %v to %v.\n", time.Unix(earliest/1000, 0), time.Unix(latest/1000, 0))
	fmt.Printf("Total of %d media files identified, with %d of them ignored (deleted).\n", total, ignored)
	fmt.Printf("%d have been reviewed, while %d have not.\n", reviewed, notReviewed)
	for extension, count := range counts {
		fmt.Printf("\tThere are %d %s files.\n", count, extension)
	}
	fmt.Printf(delim + "\n")
}
