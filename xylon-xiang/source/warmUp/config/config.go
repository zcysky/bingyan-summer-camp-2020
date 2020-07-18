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
	JWTTokenLife       int32
}

type RedisConfig struct {
	RedisAddress      string
	RedisTokenLife    int32
	RedisHistoryLimit int64
}

type ConfigObject struct {
	DataBase DatabaseConfig
	JWT      JWTConfig
	Redis    RedisConfig
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
		RedisTokenLife:    600,
		RedisHistoryLimit: 6,
	}

	Config.DataBase = databaseConfig
	Config.JWT = jwtConfig
	Config.Redis = redisConfig
}
