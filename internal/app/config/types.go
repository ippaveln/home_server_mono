package config

type Config struct {
	TgBot botConfig

	Debug bool

	Server serverConfig

	HA ha
}

type botConfig struct {
	Token      string
	Debug      bool
	IppavelnID int64
}

type serverConfig struct {
	Debug bool
	Ports ports
}

type ports struct {
	Http    int64
	Grpc    int64
	Swagger int64
}

type ha struct {
	Token string
	Host  string
}
