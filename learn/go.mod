module github.com/Sunshine9d/golang-example/learn

go 1.23.2

require (
	github.com/Sunshine9d/golang-example/cryptit v0.0.0-00010101000000-000000000000
	github.com/pborman/uuid v1.2.1
)

require github.com/google/uuid v1.0.0 // indirect

replace github.com/Sunshine9d/golang-example/cryptit => ../cryptit
