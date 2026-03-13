package db

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"sort"
	"strings"

	"github.com/charmbracelet/log"
)

// DB is the shared database connection pool.
var DB *sql.DB

//go:embed migrations/*.sql
var migrationsFS embed.FS

// Connect opens a database connection based on the engine and name/DSN.
// Supported engines: "sqlite", "sqlite3", "postgres", "postgresql", "mysql".
func Connect(engine, nameOrDSN string) error {
	driver, err := resolveDriver(engine)
	if err != nil {
		return err
	}

	DB, err = sql.Open(driver, nameOrDSN)
	if err != nil {
		return fmt.Errorf("open db: %w", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("ping db: %w", err)
	}

	log.Info("database connected", "engine", engine)
	return nil
}

// Migrate runs all pending SQL migration files from the embedded
// migrations directory. Migrations are tracked in the schema_migrations table
// and applied in lexicographic filename order. Only Up migrations are
// supported — files must end in .sql and will be executed once.
func Migrate() error {
	if DB == nil {
		return fmt.Errorf("database not connected, call Connect first")
	}

	// Ensure the tracking table exists
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS schema_migrations (
		filename  TEXT PRIMARY KEY,
		applied_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		return fmt.Errorf("create schema_migrations: %w", err)
	}

	// Collect all .sql files from the embedded FS
	entries, err := fs.ReadDir(migrationsFS, "migrations")
	if err != nil {
		return fmt.Errorf("read migrations dir: %w", err)
	}

	var files []string
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".sql") {
			files = append(files, e.Name())
		}
	}
	sort.Strings(files)

	applied := 0
	for _, filename := range files {
		// Check if already applied
		var count int
		DB.QueryRow(`SELECT COUNT(*) FROM schema_migrations WHERE filename = ?`, filename).Scan(&count)
		if count > 0 {
			continue
		}

		content, err := migrationsFS.ReadFile("migrations/" + filename)
		if err != nil {
			return fmt.Errorf("read migration %s: %w", filename, err)
		}

		// Extract the Up section (between "-- +goose Up" and "-- +goose Down")
		sql := extractUpSQL(string(content))
		if sql == "" {
			sql = string(content) // no goose annotations — run the whole file
		}

		// Execute the migration
		if _, err := DB.Exec(sql); err != nil {
			return fmt.Errorf("apply migration %s: %w", filename, err)
		}

		// Record it
		if _, err := DB.Exec(`INSERT INTO schema_migrations (filename) VALUES (?)`, filename); err != nil {
			return fmt.Errorf("record migration %s: %w", filename, err)
		}

		log.Info("migration applied", "file", filename)
		applied++
	}

	if applied == 0 {
		log.Info("no new migrations")
	}
	return nil
}

// Close closes the database connection pool.
func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

func resolveDriver(engine string) (string, error) {
	switch strings.ToLower(engine) {
	case "sqlite", "sqlite3":
		return "sqlite3", nil
	case "postgres", "postgresql":
		return "postgres", nil
	case "mysql":
		return "mysql", nil
	default:
		return "", fmt.Errorf("unsupported db engine %q (supported: sqlite, postgres, mysql)", engine)
	}
}

// extractUpSQL parses the Up section from a goose-annotated SQL file.
func extractUpSQL(content string) string {
	lines := strings.Split(content, "\n")
	var inUp bool
	var sb strings.Builder

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "-- +goose Up" {
			inUp = true
			continue
		}
		if trimmed == "-- +goose Down" {
			break
		}
		if trimmed == "-- +goose StatementBegin" || trimmed == "-- +goose StatementEnd" {
			continue
		}
		if inUp {
			sb.WriteString(line)
			sb.WriteByte('\n')
		}
	}
	return strings.TrimSpace(sb.String())
}
