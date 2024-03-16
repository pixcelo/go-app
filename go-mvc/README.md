# Go MVC Application

`gin`
```bash
go get -u github.com/gin-gonic/gin
```

`GORM`
```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

## 設定ファイル
`viper`
```bash
go get -u github.com/spf13/viper
```

設定は`config.yml`に記述してviperで読み込む
```yml
# Database
database:
  host: "127.0.0.1"
  port: "1433"
  user: ""
  password: ""
  name: ""

# Facebook
facebook:
  clientId: ""
  clientSecret: ""
  redirectURL: "http://localhost:8080/callback"
```


```bash
go get -u golang.org/x/oauth2
```

https://pkg.go.dev/golang.org/x/oauth2