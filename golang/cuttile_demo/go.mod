module cuttile_demo

go 1.15

require (
	github.com/BurntSushi/toml v0.4.1 // indirect
	github.com/airmap/gdal v0.0.5
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-spatial/geom v0.0.0-20190821234737-802ab2533ab4
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/lib/pq v1.10.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.5.1
	go.uber.org/zap v1.16.0
	golang.org/x/net v0.0.0-20200625001655-4c5254603344 // indirect
	golang.org/x/sys v0.0.0-20201214210602-f9fddec55a1e // indirect
	golang.org/x/tools v0.0.0-20200507205054-480da3ebd79c // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.20.12
)

replace github.com/airmap/gdal v0.0.5 => github.com/eanxu/gdal v0.0.8
