package fiber

type Config struct {
	Port                  int `json:"HTTP_PORT" env:"HTTP_PORT" required:"true"`
	DisableStartupMessage bool
}
