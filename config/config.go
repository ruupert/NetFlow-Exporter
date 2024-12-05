package config

import (
	"fmt"
	"time"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return fmt.Sprintf("%v", *i)
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func (i *arrayFlags) Contains(value string) bool {
	for _, item := range *i {
		if item == value {
			return true
		}
	}
	return false
}

// Config is the main configuration type
type Config struct {
	LogLevel        string        `default:"info" usage:"Log level."`
	LogFormat       string        `default:"text" usage:"Log format. It's text or json."`
	ListenAddress   string        `default:":2055" usage:"Network address to accept packets."`
	MetricAddress   string        `default:":9438" usage:"Network address to expose metrics."`
	MetricPath      string        `default:"/metrics" usage:"Network path to expose metrics."`
	Include         string        `default:"Count$" usage:"Types to include in collect process. It's regex."`
	Exclude         string        `default:"Time" usage:"Types to exclude in collect process. It's regex."`
	SampleExpire    time.Duration `default:"5m" usage:"How long a sample is valid for."`
	GeoipCities     arrayFlags
	GeoipAsns       arrayFlags
	GeoipResolveSrc bool `default:"false" usage:"Enables src IP geolocation"`
	GeoipResolveDst bool `default:"false" usage:"Enables dst IP geolocation"`
}
