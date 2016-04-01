package server

import (
	"fmt"
	"github.com/northbright/pathhelper"
	"path"
)

var (
	// Server Root Directory.
	ServerRoot string = ""
	// Directories
	Dirs map[string]string = map[string]string{}
)

// Initialize Global Variables
func init() {
	ServerRoot, _ = pathhelper.GetCurrentExecDir()
	Dirs["static"] = path.Join(ServerRoot, "static")
	Dirs["js"] = path.Join(Dirs["static"], "js")
	Dirs["css"] = path.Join(Dirs["static"], "css")
	fmt.Printf("ServerRoot = %v\n", ServerRoot)
	for k, v := range Dirs {
		fmt.Printf("Dirs[\"%v\"] = %v\n", k, v)
	}
}
