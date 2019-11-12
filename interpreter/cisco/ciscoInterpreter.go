// Package interpreter provides functions for configuration files ans returns readable analyses for Cisco router configuration files
package interpreter

import (
	"io/ioutil"
	"path/filepath"
	"regexp"
)

// enumeration replacement to be used by external functions to access the right parameters from the configuration.
const (
	// Interfaces parameter
	Interfaces = "interfaces"
	// Vlans parameter
	Vlans = "vlans"
	// Dhcp parameter
	Dhcp = "dhcp"
)

var (
	aconfig  string
	fpath, _ = filepath.Abs("../golangAPPS/interpreter/cisco/configurationExample.conf")
	regexMap = map[string]string{
		Interfaces: "interface.*",
		Vlans:      ".*vlan.*",
		Dhcp:       ".*dhcp.*\n/!"}
)

// Check checks for errors and if so throws a panic
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Analyse analyses the given configuration file and outputs it.
func Analyse() string {
	if aconfig == "" {
		dat, err := ioutil.ReadFile(fpath)
		check(err)
		aconfig = string(dat)
		return aconfig
	}
	return aconfig
}

// FindAll returns all patterns in a map
func FindAll() map[string][]string {
	output := make(map[string][]string)
	for k, v := range regexMap {
		s := Find(v, Analyse())
		output[k] = s
	}
	return output
}

// FindAllInMap returns all the findfunctions output in a map
func FindAllInMap() map[string][]string {
	output := map[string][]string{
		Interfaces: FindInterfaces(),
		Vlans:      FindVlans()}
	return output
}

// Find returns the matches according to the regexpattern given.
func Find(pattern string, config string) []string {
	r := regexp.MustCompile(pattern)
	matches := r.FindAllString(config, -1)
	return matches
}

// FindInterfaces returns every match with the interfaces in the configruration file.
func FindInterfaces() []string {
	s := regexMap["interfaces"]
	f := Find(s, Analyse())
	return f
}

// FindVlans returs every match with the vlans in the configuration file.
func FindVlans() []string {
	return Find(regexMap["vlans"], Analyse())
}
