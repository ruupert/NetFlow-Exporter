package mmdb

import (
	"fmt"
	"log"
	"net/netip"

	mmdb "github.com/oschwald/maxminddb-golang/v2"
)

type Mmdb struct {
	GeoIpReader *mmdb.Reader
	GeoIpDst    bool
	GeoIpSrc    bool
}

type MmdbStruct struct {
	City struct {
		Names struct {
			En string `maxminddb:"en"`
		} `maxminddb:"names"`
	} `maxminddb:"city"`
	Location struct {
		Latitude  float64 `maxminddb:"latitude"`
		Longitude float64 `maxminddb:"longitude"`
	} `maxminddb:"location"`
	Country struct {
		IsoCode string `maxminddb:"iso_code"`
	} `maxminddb:"country"`
}

type ASN struct {
	AutonomousSystemNumber       uint   `maxminddb:"autonomous_system_number"`
	AutonomousSystemOrganization string `maxminddb:"autonomous_system_organization"`
}

func (c *Mmdb) NewMmdb(filepath string, src bool, dst bool) {
	c.GeoIpDst = dst
	c.GeoIpSrc = src
	db, err := mmdb.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	c.GeoIpReader = db
}

func (c *Mmdb) Mmdb(ip string) *MmdbStruct {
	addr := netip.MustParseAddr(ip)
	record := new(MmdbStruct)
	result := c.GeoIpReader.Lookup(addr)
	err := result.Decode(&record)
	if err != nil {
		fmt.Println("errorrrrein")
		fmt.Println(err)
	}
	return record
}

func (c *Mmdb) Asn(ip string) *ASN {
	addr := netip.MustParseAddr(ip)
	record := new(ASN)
	result := c.GeoIpReader.Lookup(addr)
	err := result.Decode(&record)
	if err != nil {
		fmt.Println("errorrrrein")
		fmt.Println(err)
	}
	return record
}
