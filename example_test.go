package csv_rotate_test

import (
	"log"

	"github.com/goex-top/csv-rotate"
)

// To use lumberjack with the standard library's log package, just pass it into
// the SetOutput function when your application starts.
func Example() {
	log.SetOutput(&csv_rotate.Csv{
		Filename:   "/var/log/myapp/foo.csv",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   // days
		Compress:   true, // disabled by default
	})
}
