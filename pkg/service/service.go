package service

import (
	"example.com/enterprise_grade_go_project/pkg/repository"
)

// Service represents the application service.
type Service struct {
	UserRepository *repository.UserRepository
}

// NewService returns a new instance of the Service.
func NewService() *Service {
	return &Service{UserRepository: repository.NewUserRepository(getDBConnection())}
}

func getDBConnection() *sql.DB {
	// Create a new database connection
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error opening database connection: ", err)
	}
	return conn
}