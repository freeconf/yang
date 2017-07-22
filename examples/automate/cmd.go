package automate

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"time"

	"syscall"

	"github.com/c2stack/c2g/c2"
	"github.com/c2stack/c2g/meta"
	"github.com/c2stack/c2g/restconf"
)

type CmdSystem struct {
	Verbose      bool
	VarDir       string
	ExamplesDir  string
	YangPath     meta.StreamSource
	proxy        *Handle
	entries      map[string]cmdEntry
	nextPort     int
	nextDeviceId int
	proxyPort    int
}

var startup = `{
	"restconf" : {
		"web" : {
			"port" : ":{{.Port}}"
		},
		"callHome" : {
			"deviceId" : "{{.Device}}",
			"localAddress" : "http://{REQUEST_ADDRESS}:{{.Port}}/restconf",
			"address" : "http://127.0.0.1:{{.ProxyPort}}",
			"endpoint" : "/restconf"
		}
	}	
}`

type cmdEntry struct {
	deviceId  string
	port      int
	proxyPort int
	role      string
	cmd       *exec.Cmd
}

func (self *cmdEntry) address() string {
	s := fmt.Sprintf("http://127.0.0.1:%d/restconf", self.proxyPort)
	if self.role != "proxy" {
		s += "=" + self.deviceId
	}
	return s
}

func (self *CmdSystem) Startup(e *cmdEntry, out io.Writer) error {
	t, err := template.New("startup").Parse(startup)
	if err != nil {
		return err
	}
	return t.Execute(out, struct {
		ProxyPort int
		Port      int
		Device    string
	}{
		ProxyPort: self.proxyPort,
		Port:      e.port,
		Device:    e.deviceId,
	})
}

func (self *CmdSystem) New(role string) (*Handle, error) {
	if self.proxy == nil && role != "proxy" {
		return nil, c2.NewErrC("need to create proxy first", 400)
	}

	e := self.nextEntry(role)
	varDir, _ := filepath.Abs(self.VarDir)
	startup := fmt.Sprintf("%s/%s-startup.json", varDir, e.deviceId)
	c2.Debug.Print(startup)
	configFile, err := os.Create(startup)
	if err != nil {
		return nil, err
	}
	self.Startup(e, configFile)
	configFile.Close()
	cmd := exec.Command("go", "run", "main.go", "-startup", startup)

	// create a parent group pid to kill all children of process, because
	//  go run ... creates at least one child process.
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	log := fmt.Sprintf("%s/%s.log", self.VarDir, e.deviceId)
	cmd.Stdout, err = os.Create(log)
	cmd.Stderr = cmd.Stdout
	cmd.Dir = fmt.Sprintf("%s/%s/cmd", self.ExamplesDir, role)
	go func() {
		if err := cmd.Run(); err != nil {
			panic(err)
		}
	}()

	// hack : wait a bit, then try to connect to restconf service via proxy
	// better option would be to poll until avail or timeout
	<-time.After(2 * time.Second)
	d, err := restconf.ProtocolHandler(self.YangPath)(e.address())
	if err != nil {
		return nil, err
	}

	h := &Handle{
		Id:     e.deviceId,
		Device: d,
		Close: func() {
			syscall.Kill(-cmd.Process.Pid, 15)
		},
	}
	if role == "proxy" {
		self.proxy = h
	}
	return h, nil
}

func (self *CmdSystem) nextEntry(role string) *cmdEntry {
	e := cmdEntry{
		role:      role,
		deviceId:  fmt.Sprintf("%s%d", role, self.nextDeviceId),
		proxyPort: self.proxyPort,
	}
	if role == "proxy" {
		e.port = self.proxyPort
	} else {
		e.port = self.nextPort
		self.nextPort++
	}
	self.nextDeviceId++
	return &e
}
