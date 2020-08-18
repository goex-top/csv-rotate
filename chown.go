// +build !linux

package csv_rotate

import (
	"os"
)

func chown(_ string, _ os.FileInfo) error {
	return nil
}
