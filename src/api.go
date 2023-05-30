package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type TimeService struct {
	privateIPBlocks []*net.IPNet
}

type checkIPRequest struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

func (t *TimeService) CheckUserIP(c *gin.Context) {
	userIP := c.RemoteIP()
	if t.isPrivateIP(net.ParseIP(userIP)) {
		c.JSON(http.StatusOK, gin.H{
			"userIP":       userIP,
			"timeZone":     time.Now().Location().String(),
			"timeReadable": time.Now(),
			"timeUnix":     time.Now().Unix(),
		})
		return
	}

	r, err := http.Get("http://ip-api.com/json/" + userIP)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, "cant make a request")
		return
	}

	var data checkIPRequest
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, "cant decode IPRequest response")
		return
	}

	if data.Status != "success" {
		c.JSON(http.StatusInternalServerError, "status: "+data.Status)
		return
	}

	loc, err := time.LoadLocation(data.Timezone)
	if err != nil {
		log.WithFields(log.Fields{
			"timezone": data.Timezone,
		}).Error(err)
		c.JSON(http.StatusInternalServerError, "cant load timezone")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userIP":       userIP,
		"timeZone":     data.Timezone,
		"timeReadable": time.Now().In(loc),
		"timeUnix":     time.Now().In(loc).Unix(),
	})
}

func (t *TimeService) isPrivateIP(ip net.IP) bool {
	if ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() {
		return true
	}

	for _, block := range t.privateIPBlocks {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}

func NewTimeService() *TimeService {
	var privateIPBlocks []*net.IPNet

	for _, cidr := range []string{
		"127.0.0.0/8",    // IPv4 loopback
		"10.0.0.0/8",     // RFC1918
		"172.16.0.0/12",  // RFC1918
		"192.168.0.0/16", // RFC1918
		"169.254.0.0/16", // RFC3927 link-local
		"::1/128",        // IPv6 loopback
		"fe80::/10",      // IPv6 link-local
		"fc00::/7",       // IPv6 unique local addr
	} {
		_, block, err := net.ParseCIDR(cidr)
		if err != nil {
			panic(fmt.Errorf("parse error on %q: %v", cidr, err))
		}
		privateIPBlocks = append(privateIPBlocks, block)
	}
	return &TimeService{privateIPBlocks: privateIPBlocks}
}
