module github.com/danargh/golang-crud-apis

go 1.23.3

require (
	// validator
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2
	// request router & dispatcher
	github.com/gorilla/mux v1.8.0
	// extension to sql
	github.com/jmoiron/sqlx v1.3.5
	// for loading dotenv file
	github.com/joho/godotenv v1.5.1
	// manage config
	github.com/kelseyhightower/envconfig v1.4.0
	// for postgre
	github.com/lib/pq v1.10.9
	github.com/rs/cors v1.10.1
	// to logger
	github.com/sirupsen/logrus v1.9.3
)

require golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
