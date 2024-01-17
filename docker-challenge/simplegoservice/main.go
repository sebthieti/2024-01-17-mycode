/* Alta3 Research | RZFeeser
   @rzfeeser      | https://alta3.com

   Endpoints:
   /
   /ping
   /spock
   /env
   /health
   /alta3
   /info

   // coming soon!
   /talkingparrot
       ?say=Hello%20Parrot (say=string)      */

package main

import (
    "net/http"
	"os"
	"time"
    "strconv"
    "net"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// set the version of the Go Server
	version := os.Getenv("VERSION")
	if version == "" {
		version = "0.1"
	}

	// set the HEALTH_DELAY
        healthDelay := 0
        val, ok := os.LookupEnv("HEALTH_DELAY")
        if ok {
            healthDelay, _ = strconv.Atoi(val)
        } 


	// set the port to listen on
	//httpPort := os.Getenv("HTTP_PORT")
	httpPort := os.Getenv("PORT0")
	if httpPort == "" {
		httpPort = "9876"
	}

	// env - environment response
	type Env struct {
		Version string `json:"version"`
		Environmentals []string `json:"env"`
	}

	// health - health response
	type Health struct {
		Healthy bool   `json:"healthy"`
		Version string `json:"version"`
		Delay   int    `json:"delay in seconds"`
	}

    // alta3 - embedded within the alta3 endpoint response
	type Info struct {
		Homepage string `json:"homepage"`
		Youtube  string `json:"youtube"`
		Posters  string `json:"posters"`
	}

	// alta3 - alta3 response
	type Alta struct {
		Version string `json:"version"`
		Thanks  string `json:"thanks"`
		Info    Info   `json:"info"`
	}

	type Host struct {
		   Ip   []string `json:"ip"`
		   Port string `json:"port"`
	}

	// header info
	type Ipinfo struct {
		Version string `json:"version"`
		From    string `json:"from"`
		Host    Host   `json:"host"`
	}
	
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// respond with 200 to "/"
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<3")
	})

	// respond with 200 to "/ping"
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	// respond with 200 to "/spock" - Return JSON - "Status": "Live long and prosper"
	e.GET("/spock", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "Live long and prosper"})
	})

	// respond with 200 to "/env" - return version and all environmental variables"
	e.GET("/env", func(c echo.Context) error {
		r := &Env{
			Version: version,
			Environmentals: os.Environ(),
		}
		return c.JSON(http.StatusOK, r)
	})

	// respond with 200 to "/health" after healthDelay in seconds
	e.GET("/health", func(c echo.Context) error {
		r := &Health{
			Healthy: true,
			Version: version,
			Delay:   healthDelay,
		}
		time.Sleep(time.Duration(healthDelay) * time.Second)
		return c.JSON(http.StatusOK, r)
	})

	// respond with 200 to "/alta3" with info about Alta3
	e.GET("/alta3", func(c echo.Context) error {

		r := &Alta{
			Version: version,
			Thanks:  "Thank you for training with Alta3 Research!",
			Info: Info{
				Homepage: "https://alta3.com",
				Youtube:  "https://youtube.com",
				Posters:  "https://alta3.com/posters",
			},
		}
		return c.JSON(http.StatusOK, r)
	})

	e.GET("/info", func(c echo.Context) error {

	var myips []string

	// get list of available addresses
	addr, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	for _, addr := range addr {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			// check if IPv4 or IPv6 is not nil
			if ipnet.IP.To4() != nil {   // || ipnet.IP.To16 != nil {   // just grab the IPv4 address (es) 
				// print available addresses
				myips = append(myips, ipnet.IP.String())
			}
		}
	}

		r := &Ipinfo{
			Version: version,
			From: c.RealIP(),
			Host: Host{
			  Ip: myips,
			  Port: httpPort,
			},
		}		
        return c.JSON(http.StatusOK, r)	
    })

	// start the webservice on the designated port
	e.Logger.Fatal(e.Start(":" + httpPort))
}


