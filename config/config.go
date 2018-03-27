package config

type (
	DB struct {
		Host       string
		Port       string
		Username   string
		Password   string
		Dbname     string
		Collection string
	}
)

var (
	DBCONFIG *DB
)

func init() {
	DBCONFIG = &DB{
		Host:       "localhost",
		Port:       ":27017",
		Username:   "",
		Password:   "",
		Dbname:     "golang",
		Collection: "employee",
	}
}
