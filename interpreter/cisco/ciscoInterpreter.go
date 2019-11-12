// Package interpreter provides functions for configuration files ans returns readable analyses for Cisco router configuration files
package interpreter

import (
	"io/ioutil"
	"path/filepath"
	"regexp"
)

var (
	aconfig  string = "github.com/kimdepypere/golangAPPS/interpreter/cisco/configurationExample.conf"
	fpath, _        = filepath.Abs("../golangAPPS/interpreter/cisco/configurationExample.conf")

	regexMap = map[string]string{
		"interfaces": ".*interface.*"}
)

// Check checks for errors and if so throws a panic
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Analyse analyses the given configuration file and outputs it.
func Analyse() string {
	dat, err := ioutil.ReadFile(fpath)
	check(err)
	return string(dat)
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
	//return Find(regexMap["interfaces"], Analyse())
}
