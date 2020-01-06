package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"

	raft "github.com/hashicorp/raft"
	log "github.com/sirupsen/logrus"
)

// "github.com/hashicorp/raft" "github.com/hashicorp/raft-boltdb"

const (
	toolVersion = "1.1.0"
)

var genesisPtr *bool
var dataDirectoryPtr *string
var nodeIPPtr, nodeHTTPPortPtr, nodeTCPPortPtr *string
var networkIPPtr, networkTCPPortPtr *string

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

func init() {

	// SET FORMAT
	//log.SetFormatter(&log.TextFormatter{})
	// log.SetFormatter(&log.JSONFormatter{})

	// SET OUTPUT (DEFAULT stderr)
	//log.SetOutput(os.Stdout)

	// DATA DIRECTORY
	dataDirectoryPtr = flag.String("dataDirectory", "/raft", "Directory to store Raft data")
	// CREATE FIRST NODE (GENISIS)
	genesisPtr = flag.Bool("genesis", false, "Create your first Node")
	// LOG LEVEL
	logLevelPtr := flag.String("loglevel", "info", "LogLevel (info, debug or trace)")
	// NETWORK IP
	networkIPPtr = flag.String("networkip", "127.0.0.1", "Network IP")
	// NETWORK TCP PORT
	networkTCPPortPtr = flag.String("networktcpport", "3000", "Network TCP Port")
	// NODE HTTP PORT
	nodeHTTPPortPtr = flag.String("nodehttpport", "2000", "Node HTTP Port")
	// NODE IP
	nodeIPPtr = flag.String("nodeip", "127.0.0.1", "Node IP")
	// NODE TCP PORT
	nodeTCPPortPtr = flag.String("nodetcpport", "3000", "Node TCP Port")
	// VERSION FLAG
	versionPtr := flag.Bool("v", false, "prints current version")

	// Parse the flags
	flag.Parse()

	// SET LOG LEVEL
	if *logLevelPtr == "trace" {
		log.SetLevel(log.TraceLevel)
	} else if *logLevelPtr == "debug" {
		log.SetLevel(log.DebugLevel)
	} else if *logLevelPtr == "info" {
		log.SetLevel(log.InfoLevel)
	} else {
		log.Error("Error setting log levels")
		os.Exit(0)
	}

	// CHECK VERSION
	if *versionPtr {
		fmt.Println(toolVersion)
		os.Exit(0)
	}

}

func main() {

	fmt.Println("hi")

	// CREATE DATA DIRECTORY IF DOESN'T EXIST
	err := os.MkdirAll(*dataDirectoryPtr, 0700)
	checkErr(err)

	// LOAD RAFT
	raftConfig := raft.DefaultConfig()
	raftConfig.LocalID = raft.ServerID(*nodeIPPtr)

	address, err := net.ResolveTCPAddr("tcp", *nodeIPPtr)
	checkErr(err)

	// THIS IS NORMAL LOGGER
	// transportLogger := log.With().Str("component", "raft-transport").Logger()
	transportLogger := log.WithFields(log.Fields{"component": "s", "raft-transport": "ss"})

	transport, err := raft.NewTCPTransport(address.String(), address, 3, 10*time.Second, transportLogger.Info(???????))

	//snapshotStoreLogger := log.With().Str("component", "raft-snapshots").Logger()
	//snapshotStore, err := raft.NewFileSnapshotStore(config.DataDir, 1, snapshotStoreLogger)

	//logStore, err := raftboltdb.NewBoltStore(filepath.Join(config.DataDir, "raft-log.bolt"))

	//stableStore, err := raftboltdb.NewBoltStore(filepath.Join(config.DataDir, "raft-stable.bolt"))

}
