package config

type DatabaseConfig struct {
	DatabaseAddress    string
	DatabaseName       string
	CollectionUserName string
}

type JWTConfig struct {
	JWTSecret          string
	JWTBackStageSecret string
	JWTSigningMethod   string
	JWTTokenLife       int64
}

type RedisConfig struct {
	RedisAddress      string
	RedisTokenLife    int32
	RedisHistoryLimit int64
}

type MailConfig struct {
	MailAddress string
	SMTPAddress string
	Name		string
	Password	string
	MailPort	int
}

type ConfigObject struct {
	DataBase DatabaseConfig
	JWT      JWTConfig
	Redis    RedisConfig
	Mail	MailConfig
}

var Config ConfigObject

func init() {
	databaseConfig := DatabaseConfig{
		DatabaseAddress:    "mongodb://@localhost:27017",
		DatabaseName:       "warmup",
		CollectionUserName: "user",
	}

	jwtConfig := JWTConfig{
		JWTSigningMethod: "HS256",
		JWTTokenLife:     4,
	}

	redisConfig := RedisConfig{
		RedisAddress: "localhost:6379",
		RedisTokenLife:    600,
		RedisHistoryLimit: 6,
	}

	mailConfig := MailConfig{
		MailAddress: "3102131813@qq.com",
		SMTPAddress: "smtp.qq.com",
		Name: "3102131813",
		Password: "xx",

	}

	Config.DataBase = databaseConfig
	Config.JWT = jwtConfig
	Config.Redis = redisConfig
	Config.Mail = mailConfig
}
