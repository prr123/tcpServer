// program that creates a socket with an event loop
// author: prr, azul software
// date: 4/7/2023
// copyright: 2023 prr azul softwre
//

package main

import (
	"os"
	"fmt"
	"log"

	ev "server/tcp/evLoopLib"
 //   yaml "github.com/goccy/go-yaml"
    util "github.com/prr123/utility/utilLib"
	)


func main() {

    numarg := len(os.Args)
    dbg := true
    flags:=[]string{"dbg","cfg"}

	cfgFilnam := "evloop.yaml"

    useStr := "./evloop [/cfg=config yaml] [/dbg]"
    helpStr := "program that creates a server listening on port [portnumber]\n"

    if numarg > 4 {
        fmt.Println("too many arguments in cl!")
        fmt.Println("usage: %s\n", useStr)
        os.Exit(-1)
	}

   if numarg > 1 && os.Args[1] == "help" {
		fmt.Printf("help:\n%s\n", helpStr)
		fmt.Printf("\nusage is: %s\n", useStr)
		os.Exit(1)
	}


	flagMap, err := util.ParseFlags(os.Args, flags)
	if err != nil {log.Fatalf("util.ParseFlags: %v\n", err)}

	_, ok := flagMap["dbg"]
	if ok {dbg = true}
	if dbg {
		for k, v :=range flagMap {
			fmt.Printf("flag: %s value: %s\n", k, v)
		}
	}

	val, ok := flagMap["cfg"]
	if !ok {
		log.Printf("default cfgList: %s\n", cfgFilnam)
	} else {
		if val.(string) == "none" {log.Fatalf("no yaml file provided with /cfg  flag!")}
		cfgFilnam = val.(string)
		log.Printf("cfg File: %s\n", cfgFilnam)
	}

	cfg, err := ev.InitCfg(cfgFilnam)
	if err != nil {log.Fatalf("InitCfg: %v\n", err)}
	ev.PrintCfg(cfg)

	log.Printf("server ending successfully!\n")
}
