package config

type (
	Config struct {
		Db Db
	}

	Db struct {
		Url string
	}
)

func GetConfig() Config {
	return Config{
		Db: Db{
			Url: "mongodb://root:123456@0.0.0.0:27019",
		},
	}
}
