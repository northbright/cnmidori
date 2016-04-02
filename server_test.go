package cnmidori_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/northbright/cnmidori"
	"github.com/northbright/pathhelper"
)

const (
	settingsStr string = `
{
    "redis-servers":[
        {"name":"user", "addr":"localhost:6379", "password":"123456"},
        {"name":"data", "addr":"localhost:6380", "password":"123456"}
    ]
}`
)

func ExampleNewServer() {
	serverRoot, _ := pathhelper.GetCurrentExecDir()

	settingsFile := path.Join(serverRoot, "settings.json")
	if err := ioutil.WriteFile(settingsFile, []byte(settingsStr), 0755); err != nil {
		fmt.Fprintf(os.Stderr, "ioutil.WriteFile() error: %v\n", err)
		return
	}

	server, err := cnmidori.NewServer(settingsFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "NewServer(%v) error: %v\n", settingsFile, err)
		return
	}
	fmt.Fprintf(os.Stderr, "NewServer() OK. server = %v\n", server)
	// Output:
}
