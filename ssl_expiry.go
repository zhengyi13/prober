package main

import (
	"flag"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"github.com/zhengyi13/prober"
)

// Probes is a list of HostPorts
type Probes []prober.HostPort

// And a ProbeConfig struct is a mapping of a Probes list to the  "probes:" heading in a YAML config
type ProbeConfig struct {
	Probes `yaml:"probes"`
}

func main() {

	confFile := flag.String("config", "./probes.yaml", "a YAML config file of host:port pairs")
	flag.Parse()

	data, err := ioutil.ReadFile(*confFile)
	if err != nil {
		log.Printf("ERROR cannot read file %s: %v\n", *confFile, err)
	}
	config := ProbeConfig{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
        log.Fatalf("Unable to unmarshal your config: %v\n", err)
	}
	log.Printf("--- config:\n%v\n\n", config)
	for i, hp := range config.Probes {
		log.Printf("Probe %d: %s\n", i, hp)
		ts, err := prober.Probe(hp)
		if err != nil {
			log.Printf("ERROR Failed to probe %s\n", hp)
		}
		log.Printf("Probe %d timestamp: %d\n", i, ts)
	}

}
