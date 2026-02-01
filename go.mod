module github.com/yuriy-kazanov/arch-go-example

go 1.24.7

replace (
	gitlab.tbank.ru/ca-business-common => ../ca-business-common
)

require (
	github.com/jackc/pgx/v5 v5.8.0
	go.opentelemetry.io/otel v1.39.0
)

require (
	gitlab.tbank.ru/ca-business-common v0.0.0
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/otel/metric v1.39.0 // indirect
	go.opentelemetry.io/otel/trace v1.39.0 // indirect
	golang.org/x/text v0.29.0 // indirect
)
