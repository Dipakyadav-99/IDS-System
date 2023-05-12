package runner

import (
	"fmt"
	"os"

	"ids.app/common"
)

func showVersion() {
	fmt.Printf("ids %s\n", common.Version)
	os.Exit(2)
}
