package runner

import (
	"github.com/projectdiscovery/gologger"
	"ids.app/common"
)

// showBanner is used to show the banner to the user
func showBanner() {
	gologger.Print().Msgf("%s\n\n", common.Banner)
	gologger.Print().Msgf("\t%s\n\n", common.Email)
	
}
