package server

import (
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
}
