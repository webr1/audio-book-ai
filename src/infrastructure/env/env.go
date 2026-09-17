package env

import (
	"os"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Env struct {
	Port string `env:"PORT" envDefault:"8080"`

	DBHost     string `env:"DB_HOST" envDefault:"localhost"`
	DBPort     string `env:"DB_PORT" envDefault:"5432"`
	DBUser     string `env:"DB_USER" envDefault:"postgres"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"postgres"`
	DBName     string `env:"DB_NAME" envDefault:"audiobook"`
	DBSSLMode  string `env:"DB_SSLMODE" envDefault:"disable"`

	RedisAddr     string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	RedisPassword string `env:"REDIS_PASSWORD" envDefault:""`

	DataDir       string `env:"DATA_DIR" envDefault:"./data"`
	MaxChunkChars int    `env:"MAX_CHUNK_CHARS" envDefault:"1000"`
	MaxUploadMB   int64  `env:"MAX_UPLOAD_MB" envDefault:"20"`
	TTSProvider   string `env:"TTS_PROVIDER" envDefault:"mock"`

	PricePerChar float64 `env:"PRICE_PER_CHAR" envDefault:"0.5"`

	GoogleClientID string `env:"GOOGLE_CLIENT_ID" envDefault:""`

	JwtSecret               string `env:"JWT_SECRET" envDefault:"dev-secret-change-me"`
	JwtAccessExpireMinutes  int64  `env:"JWT_ACCESS_EXPIRE_MINUTES" envDefault:"15"`
	JwtRefreshExpireMinutes int64  `env:"JWT_REFRESH_EXPIRE_MINUTES" envDefault:"10080"`
}

func loadDotenv() {
	if os.Getenv("CONTAINER_MODE") != "1" {
		_ = godotenv.Load("env/.env.local")
		_ = godotenv.Load("env/.env")
	}
}

// @inject
func NewEnv() *Env {
	loadDotenv()

	cfg := &Env{}
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
