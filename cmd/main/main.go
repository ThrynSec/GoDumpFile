package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/ThrynSec/GoDumpFile/cmd/internals/api"
	"github.com/ThrynSec/GoDumpFile/cmd/internals/config"
)

func main() {
	log.Println("Booting up GoDumpFile")

	flagFile := flag.String("file", "", "Path and name of the file you want to dump (ex : /home/JeanJean/script.py)")
	flagName := flag.String("name", "", "The name of the endpoint (ex : -name cve.py will create 127.0.0.1/cve.py")
	flagPort := flag.Int("port", 80, "OPTIONAL : Port you want to put GoDumpFile on")

	flag.Parse()

	// OPTIONAL - No flag catcher
	if (*flagFile == "") && (*flagName == "") {
		log.Fatal("Error, -file and -name are both mandatory to use GoDumpFile")
	}

	config.EndpointName = *flagName
	config.FilePath = *flagFile
	config.Port = *flagPort

	log.Println("parsed GoDumpFile flags :" +
		"\n - Filepath is : " + config.FilePath +
		"\n - Endpoint name is : " + config.EndpointName +
		"\n - Using port : " + strconv.Itoa(config.Port))

	log.Println("Creating the endpoint")
	api.MapUrls()
}
