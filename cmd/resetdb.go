package cmd

import (
	"fmt"
	"os"

	"github.com/rjkroege/audiobookbinder/global"
	"github.com/rjkroege/audiobookbinder/state"
)

func ResetDatabase(gctx *global.Context) error {
	if _, err := os.Stat(gctx.Dbname); err == nil {
		if err := os.Remove(gctx.Dbname); err != nil {
			return fmt.Errorf("ResetDatabase: can't remove %q: %v", gctx.Dbname, err)
		}
	}

	// At this point, there is no database file so make one.
	return state.OpenDb(gctx)
}
