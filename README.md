# Simple Authentication WebApp
A simple authentication web app in Go (using JWT)

## Routes
| Path        | Description |
| ----------- | ----------- |
| /api/v1/register      | Signup Router       |
| /api/v1/login   | Login Router        |

## Install Dependencies
```go get -u github.com/dgrijalva/jwt-go
go get -u github.com/gin-gonic/gin v1.7.7
go get -u github.com/go-playground/validator/v10 v10.10.0
go get -u gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
go get -u gorm.io/driver/postgres v1.2.3
go get -u gorm.io/gorm v1.22.5
```

## Usage
Put your postgres connection details in **config.yaml**<br>
Run this command below:
```
go run ./cmd/serve.go
```

## License
GNU General Public License, version 3
