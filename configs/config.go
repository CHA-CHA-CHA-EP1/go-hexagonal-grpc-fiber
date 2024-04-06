package configs

type Config struct {
    Database DatabaseConfig
}

type DatabaseConfig struct {
    Name        string
    Uri         string
    Username    string
    Password    string
}

func GetConfig() Config {
    return Config{
        Database: DatabaseConfig{
            Name: "auth_db",
            Uri: "mongodb://root:123456@0.0.0.0:27017",
            Username: "root",
            Password: "123456",
        },
    }
}
