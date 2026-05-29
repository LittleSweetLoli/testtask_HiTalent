package main

import (
	"context"
	"database/sql"
	"log/slog"
	"os"
	"test_task_hitalent/internal/config"
	migrationfs "test_task_hitalent/migrations"

	_ "github.com/lib/pq"
	goose "github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	cfg := config.Load()

	db := connectDB(cfg)

	runMigrations(db, cfg)

	// mux := http.NewServeMux()

	// Создать подразделение
	// mux.HandleFunc("POST /departments/")

	// Создать сотрудника в подразделении
	// mux.HandleFunc("POST /departments/{id}/employees/")

	// Получить подразделение
	// mux.HandleFunc("GET /departments/{id}")

	// Переместить подразделение в другое
	// mux.HandleFunc("PATCH /departments/{id}")

	// Удалить подразделение
	// mux.HandleFunc("DELETE /departments/{id}")
}

func connectDB(cfg *config.Config) *gorm.DB {
	var db *gorm.DB
	var err error

	// TODO: Add retry mechanism
	db, err = gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err == nil {
		sqlDB, _ := db.DB()
		if pingErr := sqlDB.Ping(); pingErr == nil {
			slog.Info("connected to database")
			return db
		}
	}

	slog.Error("could not connect to database", "err", err)
	os.Exit(1)
	return nil
}

func runMigrations(db *gorm.DB, cfg *config.Config) {
	sqlDB, err := sql.Open("postgres", cfg.DSN())
	if err != nil {
		slog.Error("migration: open db failed", "err", err)
		os.Exit(1)
	}
	defer sqlDB.Close()

	goose.SetBaseFS(migrationfs.FS)
	if err := goose.SetDialect("postgres"); err != nil {
		slog.Error("migration: set dialect failed", "err", err)
		os.Exit(1)
	}
	ctx := context.Background()
	if err := goose.RunContext(ctx, "up", sqlDB, "."); err != nil {
		slog.Error("migration failed", "err", err)
		os.Exit(1)
	}
	slog.Info("migrations applied")
}
