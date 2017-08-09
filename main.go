package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/hashicorp/hcl"
	"github.com/mitchellh/mapstructure"
)

func readFile(n string) string {
	d, err := ioutil.ReadFile(filepath.Join(".", n))
	if err != nil {
		panic(err)
	}

	return string(d)
}

func main() {
	type Certificate struct {
		Path string
		Type string
	}
	type Server struct {
		Hostname    string
		Port        int
		Secure      bool
		Certificate Certificate `mapstructure:",squash",hcl:"certificate,squash"`
	}

	type HTTPServers struct {
		Server []Server `hcl:"server,expand"`
	}
	/* r, err := os.Open("config.hcl")
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer

	if _, err := io.Copy(&buf, r); err != nil {
		panic(err)
	}
	*/

	hclParseTree, err := hcl.Parse(readFile("config.hcl"))
	if err != nil {
		panic(err)
	}
	var m map[string]interface{}
	err = hcl.DecodeObject(&m, hclParseTree)
	if err != nil {
		panic(err)
	}
	httpservers := HTTPServers{}
	fmt.Println(m)
	var metadata mapstructure.Metadata

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: &metadata,
		Result:   &httpservers,
	})

	if err != nil {
		panic(err)
	}
	if err := decoder.Decode(m); err != nil {
		panic(err)
	}

	fmt.Printf("http servers %+v\n", httpservers)
	fmt.Printf("%#v\n", metadata)
	/*
		err := hcl.Decode(&httpservers, readFile("config.hcl"))
		if err != nil {
			panic(err)
		}*/
	//fmt.Println(length(httpservers.HTTPServers)
	for i, k := range httpservers.Server {
		fmt.Printf(" hello %d : %+v", i, k)
	}

}
