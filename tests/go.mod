module github.com/iznauy/gorm/tests

go 1.14

require (
	github.com/google/uuid v1.1.1
	github.com/jinzhu/now v1.1.1
	github.com/lib/pq v1.6.0
	github.com/iznauy/driver/mysql v0.2.9
	github.com/iznauy/driver/postgres v0.2.5
	github.com/iznauy/driver/sqlite v1.0.8
	github.com/iznauy/driver/sqlserver v0.2.4
	github.com/iznauy/gorm v0.2.19
)

replace github.com/iznauy/gorm => ../
