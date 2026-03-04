package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"web3-recruitment-admin/internal/handler"
	"web3-recruitment-admin/internal/middleware"
	"web3-recruitment-admin/internal/repository"
	"web3-recruitment-admin/internal/service"
)

func main() {
	cfg := loadConfig()
	dsn := buildDSN(cfg)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Println("Connected to PostgreSQL database:", cfg.DBName)

	// Run migrations
	if err := runMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize layers
	jobViewRepo := repository.NewJobViewRepository(db)
	applicationRepo := repository.NewApplicationRepository(db)
	jobRepo := repository.NewJobRepository(db)

	jobViewService := service.NewJobViewService(jobViewRepo)
	applicationService := service.NewApplicationService(applicationRepo)
	jobService := service.NewJobService(jobRepo)

	jobViewHandler := handler.NewJobViewHandler(jobViewService)
	applicationHandler := handler.NewApplicationHandler(applicationService)
	jobHandler := handler.NewJobHandler(jobService)

	// Setup Gin router
	router := gin.Default()
	
	// CORS middleware
	router.Use(middleware.CORS())

	// API routes (with basic auth for admin)
	admin := router.Group("/api/admin")
	admin.Use(middleware.BasicAuth(cfg.AdminUser, cfg.AdminPass))
	{
		jobViewHandler.RegisterRoutes(admin)
		applicationHandler.RegisterRoutes(admin)
		jobHandler.RegisterRoutes(admin)
	}

	// Public tracking endpoint (no auth required)
	router.POST("/api/track/view", jobViewHandler.TrackView)

	log.Printf("Server starting on :%s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	Port       string
	AdminUser  string
	AdminPass  string
}

func loadConfig() *Config {
	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "joy"),
		DBPassword: getEnv("DB_PASSWORD", "ServBay.dev"),
		DBName:     getEnv("DB_NAME", "web3_recruitment"),
		Port:       getEnv("PORT", "8081"),
		AdminUser:  getEnv("ADMIN_USER", "admin"),
		AdminPass:  getEnv("ADMIN_PASS", "admin123"),
	}
}

func buildDSN(cfg *Config) string {
	return "host=" + cfg.DBHost +
		" port=" + cfg.DBPort +
		" user=" + cfg.DBUser +
		" password=" + cfg.DBPassword +
		" dbname=" + cfg.DBName +
		" sslmode=disable"
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func runMigrations(db *sql.DB) error {
	migrations := []string{
		`CREATE TABLE IF NOT EXISTS job_views (
			id SERIAL PRIMARY KEY,
			job_id INTEGER NOT NULL,
			ip_address VARCHAR(50),
			user_agent TEXT,
			viewed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE INDEX IF NOT EXISTS idx_job_views_job_id ON job_views(job_id)`,
		`CREATE INDEX IF NOT EXISTS idx_job_views_viewed_at ON job_views(viewed_at)`,
	}

	for _, m := range migrations {
		if _, err := db.Exec(m); err != nil {
			return err
		}
	}

	log.Println("Database migrations completed")
	return nil
}
