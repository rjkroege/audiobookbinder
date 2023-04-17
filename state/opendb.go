package state

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"

	"github.com/rjkroege/audiobookbinder/global"
	_ "modernc.org/sqlite"
)

//go:embed schema.sql
var ddl string

func OpenDb(gctx *global.Context) error {
	dsn := "file:" + gctx.Dbname
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return fmt.Errorf("OpenDb can't open db %q: %v", gctx.Dbname, err)
	}

	// Actually make the tables if they don't exist.
	ctx := context.Background()
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return fmt.Errorf("OpenDb can't make tables: %v", err)
	}

	// Now db is ready for action.
	gctx.Db = db
	return nil
}
