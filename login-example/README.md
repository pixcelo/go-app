
# Login

```bash
go mod init login-example
go get github.com/labstack/echo/v4
```

hot reload
```bash
docker compose up
```

MySQL
```bash
go get github.com/go-sql-driver/mysql
go get github.com/jmoiron/sqlx
```

環境変数の確認
```bash
docker compose exec api bash
root@a361c9466756:/app# echo $DB_USER
# login-user
```

## Reference
- [Echo](https://echo.labstack.com/)
- [Air](https://github.com/cosmtrek/air)
- [sqlx](https://github.com/jmoiron/sqlx)