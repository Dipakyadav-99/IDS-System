package runner

import (
	"github.com/projectdiscovery/gologger"
	"ids.app/pkg/cache"
	"ids.app/pkg/errors"
)

func rmCache() {
	cache.Purge()
	gologger.Info().Msg("All local cached resources have been removed.")
	errors.Abort(9)
}
