user-role-app/
│── main.go
│── go.mod
│── config/
│ └── database.go
│── models/
│ └── user.go
│── repositories/
│ ├── user_repository.go
│ └── role_repository.go
│── services/
│ ├── user_service.go
│ └── role_service.go
│── controllers/
│ ├── user_controller.go
│ └── role_controller.go
│── routes/
│ ├── user_routes.go
│ └── role_routes.go
│── validators/
│ ├── user_validator.go
│ └── role_validator.go
│── routes.go

# Install Gin framework for handling HTTP requests

go get -u github.com/gin-gonic/gin

# Install GORM for ORM support

go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres # For PostgreSQL database driver

# Install UUID package for unique IDs

go get -u github.com/google/uuid

# Install Validator package for input validation

go get -u github.com/go-playground/validator/v10
