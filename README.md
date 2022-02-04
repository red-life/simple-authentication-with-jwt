# Simple Authentication WebApp
A simple authentication web app in Go (using JWT)

## Routes
| Path                   | Method      | Data      |
| ---------------------- | ----------- | --------- |
| /api/v1/auth/register  |    POST     | {"firstname":,"lastname":,"email":,"username":,"password":} |
| /api/v1/auth/login     |    POST     | {"email":,"password":} |
| /apiv1/auth/delete     |    GET      | None |
## Install Dependencies
```
go get -u github.com/dgrijalva/jwt-go
go get -u github.com/gin-gonic/gin
go get -u github.com/go-playground/validator/v10
go get -u gopkg.in/yaml.v3
go get -u gorm.io/driver/postgres
go get -u gorm.io/gorm
```

## Usage
Put your postgres connection details in **config.yaml**<br>
Run this command below:
```
go run
```

## License
GNU General Public License, version 3
