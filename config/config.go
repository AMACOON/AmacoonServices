package config

type Config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	MongoDBUsername string
	MongoDBPassword string
	MongoDBHost     string
	MongoDBPort     string
	MongoDBName     string
	ServerPort string
}

func LoadConfig() *Config {
	return &Config{
		DBUsername: "amacoon001_add1",
		DBPassword: "armin013",
		DBHost:     "mysql.catclubsystem.com",
		DBPort:     "3306",
		DBName:     "amacoon01",
		MongoDBUsername: "amacoonservice",
		MongoDBPassword: "2010mainecoon2010",
		MongoDBHost:     "mongo",
		MongoDBPort:     "27017",
		MongoDBName:     "amacoon",
		ServerPort: "8080",
	}
}
//"amacoon01:2010amacoon2010@tcp(mysql20-farm1.kinghost.net:3306)/amacoon01"

//"amacoon001_add1:armin013@tcp(mysql20-farm1.kinghost.net:3306)/amacoon01"
//mysql.catclubsystem.com

//mongodb://amacoonservice:2010mainecoon2010@localhost:27017/amacoon