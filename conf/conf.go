package conf

// Conf set of ENV configs for app
type Conf struct {
	Addr string `env:"ADDR" envDefault:":3000"`

	DBUser     string `env:"DB_USER" envDefault:"postgres"`
	DBPassword string `env:"DB_PASS" envDefault:"2003"`
	DBName     string `env:"DB_NAME" envDefault:"postgres"`
	DBHost     string `env:"DB_HOST" envDefault:"localhost"`

	HTTPReadTimeout  int `env:"HTTP_Read_timeout" envDefault:"300"`
	HTTPWriteTimeout int `env:"HTTP_Write_timeout" envDefault:"300"`
}