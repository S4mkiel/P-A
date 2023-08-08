package cfg

import "fmt"

type Config struct {
	Username   string `json:"username" env:"DB_USERNAME" required:"true"`
	Password   string `json:"password" env:"DB_PASSWORD" required:"true"`
	Host       string `json:"host" env:"DB_HOST" required:"true"`
	Port       string `json:"port" env:"DB_PORT" required:"true"`
	Database   string `json:"database" env:"DB_DATABASE" required:"true"`
	Connection string `json:"connection" env:"DB_CONNECTION" required:"true"`
}

func (c Config) ConnectionString() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.Host, c.Username, c.Password, c.Database, c.Port,
	)
}
