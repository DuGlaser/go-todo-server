package db

type Config struct {
	DB struct {
		Production struct {
			Host     string
			Username string
			Password string
			DBName   string
			Port     string
		}
		Test struct {
			Host     string
			Username string
			Password string
			DBName   string
			Port     string
		}
		Routing struct {
			Port string
		}
	}
}

func NewTodoConfig() *Config {
	c := new(Config)

	c.DB.Production.Host = "localhost"
	c.DB.Production.Username = "username"
	c.DB.Production.Password = "password"
	c.DB.Production.DBName = "db_name"
	c.DB.Production.Port = "3306"

	c.DB.Test.Host = "localhost"
	c.DB.Test.Username = "username"
	c.DB.Test.Password = "password"
	c.DB.Test.DBName = "db_name_test"
	c.DB.Test.Port = "3306"

	return c
}
