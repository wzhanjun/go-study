module github.com/wzhanjun/go-gin-example

go 1.17

require (
	github.com/astaxie/beego v1.12.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.7.4
	github.com/go-ini/ini v1.63.2
	github.com/jinzhu/gorm v1.9.16
	github.com/robfig/cron/v3 v3.0.1
	github.com/swaggo/gin-swagger v1.3.2
	github.com/swaggo/swag v1.7.3
	github.com/unknwon/com v1.0.1
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/ugorji/go/codec v1.1.7 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/net v0.0.0-20210421230115-4e50805a0758 // indirect
	golang.org/x/sys v0.0.0-20210420072515-93ed5bcd2bfe // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.0 // indirect
	google.golang.org/protobuf v1.23.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/wzhanjun/go-gin-example/conf => ./conf
	github.com/wzhanjun/go-gin-example/docs => ./docs
	github.com/wzhanjun/go-gin-example/middleware => ./middleware
	github.com/wzhanjun/go-gin-example/models => ./models
	github.com/wzhanjun/go-gin-example/pkg/e => ./pkg/e
	github.com/wzhanjun/go-gin-example/pkg/logging => ./pkg/logging
	github.com/wzhanjun/go-gin-example/pkg/setting => ./pkg/setting
	github.com/wzhanjun/go-gin-example/pkg/util => ./pkg/util
	github.com/wzhanjun/go-gin-example/routers => ./routers

)
