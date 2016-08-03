package main

import (
	"encoding/json"
	"fmt"
	"github.com/containous/flaeg"
	"os"
	"time"
)

//Configuration is a struct which contains all differents type to field
//using parsers on string, time.Duration, pointer, bool, int, int64, time.Time, float64
type Configuration struct {
	VersionName string        //no description struct tag, it will not be flaged
	LogLevel    string        `short:"l" description:"Log level"`      //string type field, short flag "-l"
	Timeout     time.Duration `description:"Timeout duration"`         //time.Duration type field
	Db          *DatabaseInfo `description:"Enable database"`          //pointer type field (on DatabaseInfo)
	Owner       *OwnerInfo    `description:"Enable Owner description"` //another pointer type field (on OwnerInfo)
}

//DatabaseInfo is a Sub-struct embeded in Configuration
type DatabaseInfo struct {
	ConnectionMax   uint   `long:"comax" description:"Number max of connections on database"` //uint type field, long flag "--comax"
	ConnectionMax64 uint64 `description:"Number max of connections on database"`              //uint64 type field, same description just to be sure it works
	Watch           bool   `description:"Watch device"`                                       //bool type
	IP              string `description:"Server ip address"`
}

//OwnerInfo is a Sub-struct embeded in Configuration
type OwnerInfo struct {
	Name        string    `description:"Owner name"`                     //pointer type field on string
	DateOfBirth time.Time `long:"dob" description:"Owner date of birth"` //time.Time type field, long flag "--dob"
	Rate        float64   `description:"Owner rate"`                     //float64 type field
}

func main() {
	//config contains the default configuration of the program
	config := &Configuration{
		VersionName: "Rebloch",
		LogLevel:    "INFO",
		Timeout:     time.Second,
		Owner: &OwnerInfo{
			Name:        "default owner",
			DateOfBirth: time.Now(),
			Rate:        0.5,
		},
	}
	defaultConfigOfPointerFields := &Configuration{
		Db: &DatabaseInfo{
			ConnectionMax:   100,
			ConnectionMax64: 6400000000000000000,
			Watch:           true,
			IP:              "192.168.0.1",
		},
		Owner: &OwnerInfo{
			Name: "admin",
			Rate: 1,
		},
	}

	flaegCmd := &flaeg.Command{
		Name:                  "flaegdemo",
		Description:           "flaegdemo is a golang program...(here the program description)",
		Config:                &config,
		DefaultPointersConfig: &defaultConfigOfPointerFields,
		Run: func() error {
			printableConfig, _ := json.Marshal(config)
			fmt.Printf("%s\n", printableConfig)
			return nil
		},
	}

	flaegVersionCmd := &flaeg.Command{
		Name:                  "version",
		Description:           "Print version",
		Config:                struct{}{},
		DefaultPointersConfig: struct{}{},
		Run: func() error {
			fmt.Printf("Version %s\n", config.VersionName)
			return nil
		},
	}

	f := flaeg.New(flaegCmd, os.Args[1:])
	f.AddCommand(flaegVersionCmd)
	if err := f.Run(); err != nil {
		fmt.Println(err.Error())
	}
}
