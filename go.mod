module github.com/hatamiarash7/netflow-exporter

go 1.23

toolchain go1.23.3

replace github.com/hatamiarash7/netflow-exporter => ./

require (
	github.com/itzg/go-flagsfiller v1.15.0
	github.com/prometheus/client_golang v1.20.5
	github.com/prometheus/client_model v0.6.1
	github.com/sirupsen/logrus v1.9.3
	github.com/tehmaze/netflow v0.0.0-20240303214733-8c13bb004068
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/oschwald/maxminddb-golang/v2 v2.0.0-beta.2 // indirect
	github.com/prometheus/common v0.60.1 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	golang.org/x/sys v0.27.0 // indirect
	google.golang.org/protobuf v1.35.2 // indirect
)
