1. CA用の秘密鍵を生成
```bash
$ openssl genrsa -out localCA.key 2048
```

2. 秘密鍵と識別情報から CSR (certificate signing request 証明書署名要求) を生成
```bash
$ openssl req -batch -new -key localCA.key  -out localCA.csr \
  -subj "/C=jp/ST=Osaka/L=Osaka-shi/O=\"Example Inc\"/OU=Foo/CN=tek"
```

3. 公開鍵が入ったCSRを秘密鍵で自己署名してCA証明書を作成
```bash
$ openssl x509 -req -days 3650 -signkey localCA.key -in localCA.csr -out localCA.crt
```

4. サーバー用の秘密鍵を生成
```
$ openssl genrsa -out localhost.key 2048
```

5. 秘密鍵と識別情報から CSR (certificate signing request 証明書署名要求) を生成
```bash
$  openssl req -batch -new -key localhost.key  -out localhost.csr \
  -subj "/C=jp/ST=Osaka/L=Osaka-shi/O=\"Example Inc\"/OU=Foo/CN=localhost"
```

6. 証明書に付加する SAN (Subject Alternative Name サブジェクト代替名) を入れたテキストファイルを作成
```bash
$ echo 'subjectAltName = DNS:localhost, DNS:localhost.localdomain, IP:127.0.0.1, DNS:app, DNS:app.localdomain' > localhost.csx
```

7. CSRにSAN情報も付加し、CA秘密鍵とCA証明書とで署名してサーバー証明書を作成
```bash
$ openssl x509 -req -days 1825 -CA localCA.crt -CAkey localCA.key -CAcreateserial -in localhost.csr -extfile localhost.csx -out localhost.crt
```

8. CA証明書をクライアントにインストール
- Windowsキーを押して、`certlm.msc`を実行
- 信頼されたルート証明機関に`localCA.crt`をインポート

9. サーバー側で秘密鍵とサーバー証明書を指定して起動
```go
package main

import "github.com/gin-gonic/gin"

func main() {
    router := gin.Default()
    router.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", nil)
    })
    router.RunTLS(PORT, "ssl/localhost.crt", "ssl/localhost.key")
}
```

作成されたファイル
```bash
$ ll
total 30
-rw-r--r-- 1 user 197609 1232 Mar 20 06:20 localCA.crt
-rw-r--r-- 1 user 197609 1018 Mar 20 06:20 localCA.csr
-rw-r--r-- 1 user 197609 1732 Mar 20 06:20 localCA.key
-rw-r--r-- 1 user 197609   42 Mar 20 06:20 localCA.srl
-rw-r--r-- 1 user 197609 1596 Mar 20 06:20 localhost.crt
-rw-r--r-- 1 user 197609 1026 Mar 20 06:20 localhost.csr
-rw-r--r-- 1 user 197609  102 Mar 20 06:20 localhost.csx
-rw-r--r-- 1 user 197609 1732 Mar 20 06:20 localhost.key
```