package main

import (
	"context"
	"fmt"
	"log"
	"os"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
	"github.com/syui/card-api/ent/migrate"
)

func main() {
    dataSourceName := fmt.Sprintf("file:/data/ent.sqlite?_fk=1")
				//client, err := ent.Open(dialect.SQLite, "file:/data/ent.sqlite?_fk=1")
    ctx := context.Background()

    // Create a local migration directory able to understand Atlas migration file format for replay.
    dir, err := atlas.NewLocalDir("ent/migrate/migrations")
    if err != nil {
        log.Fatalf("failed creating atlas migration directory: %v", err)
    }

    // Migrate diff options.
    opts := []schema.MigrateOption{
        schema.WithDir(dir),                         // provide migration directory
        schema.WithMigrationMode(schema.ModeInspect), // provide migration mode
        schema.WithDialect(dialect.SQLite),           // Ent dialect to use
        schema.WithFormatter(atlas.DefaultFormatter),
    }

    // Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).
    err = migrate.NamedDiff(ctx, dataSourceName, os.Args[1], opts...)
    if err != nil {
        log.Fatalf("failed generating migration file: %v", err)
    }
}
