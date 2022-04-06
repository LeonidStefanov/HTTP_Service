package option

type Options struct {
	Port   string `envconfig:"PORT" required:"true" default:"3000"`
	DBPort string `envconfig:"DB_PORT" required:"true" default:"27017"`
	DBHost string `envconfig:"DB_HOST" required:"true" default:"localhost"`
}
