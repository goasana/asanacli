package version

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	path "path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/goasana/asana/cmd/commands"
	asanaLogger "github.com/goasana/asana/logger"
	"github.com/goasana/asana/logger/colors"
	"github.com/goasana/asana/utils"
	"gopkg.in/yaml.v2"
)

const verboseVersionBanner string = `%s%s______
    ___   _____ ___    _   _____ 
   /   | / ___//   |  / | / /   |
  / /| | \__ \/ /| | /  |/ / /| |
 / ___ |___/ / ___ |/ /|  / ___ |
/_/  |_/____/_/  |_/_/ |_/_/  |_| v{{ .AsanaVersion }}%s
%s%s
├── Asana     : {{ .AsanaFrameworkVersion }}
├── GoVersion : {{ .GoVersion }}
├── GOOS      : {{ .GOOS }}
├── GOARCH    : {{ .GOARCH }}
├── NumCPU    : {{ .NumCPU }}
├── GOPATH    : {{ .GOPATH }}
├── GOROOT    : {{ .GOROOT }}
├── Compiler  : {{ .Compiler }}
└── Date      : {{ Now "Monday, 2 Jan 2006" }}%s
`

const shortVersionBanner = `______
    ___   _____ ___    _   _____ 
   /   | / ___//   |  / | / /   |
  / /| | \__ \/ /| | /  |/ / /| |
 / ___ |___/ / ___ |/ /|  / ___ |
/_/  |_/____/_/  |_/_/ |_/_/  |_|  v{{ .AsanaVersion }}
`

var CmdVersion = &commands.Command{
	UsageLine: "version",
	Short:     "Prints the current Asana version",
	Long: `
Prints the current Asana, Asana and Go version alongside the platform information.
`,
	Run: versionCmd,
}
var outputFormat string

const version = "1.0"

func init() {
	fs := flag.NewFlagSet("version", flag.ContinueOnError)
	fs.StringVar(&outputFormat, "o", "", "Set the output format. Either json or yaml.")
	CmdVersion.Flag = *fs
	commands.AvailableCommands = append(commands.AvailableCommands, CmdVersion)
}

func versionCmd(cmd *commands.Command, args []string) int {

	cmd.Flag.Parse(args)
	stdout := cmd.Out()

	if outputFormat != "" {
		runtimeInfo := RuntimeInfo{
			GetGoVersion(),
			runtime.GOOS,
			runtime.GOARCH,
			runtime.NumCPU(),
			os.Getenv("GOPATH"),
			runtime.GOROOT(),
			runtime.Compiler,
			version,
			GetAsanaVersion(),
		}
		switch outputFormat {
		case "json":
			{
				b, err := json.MarshalIndent(runtimeInfo, "", "    ")
				if err != nil {
					asanaLogger.Log.Error(err.Error())
				}
				fmt.Println(string(b))
				return 0
			}
		case "yaml":
			{
				b, err := yaml.Marshal(&runtimeInfo)
				if err != nil {
					asanaLogger.Log.Error(err.Error())
				}
				fmt.Println(string(b))
				return 0
			}
		}
	}

	coloredBanner := fmt.Sprintf(verboseVersionBanner, "\x1b[35m", "\x1b[1m",
		"\x1b[0m", "\x1b[32m", "\x1b[1m", "\x1b[0m")
	InitBanner(stdout, bytes.NewBufferString(coloredBanner))
	return 0
}

// ShowShortVersionBanner prints the short version banner.
func ShowShortVersionBanner() {
	output := colors.NewColorWriter(os.Stdout)
	InitBanner(output, bytes.NewBufferString(colors.MagentaBold(shortVersionBanner)))
}

func GetAsanaVersion() string {
	re, err := regexp.Compile(`VERSION = "([0-9.]+)"`)
	if err != nil {
		return ""
	}
	wgopath := utils.GetGOPATHs()
	if len(wgopath) == 0 {
		asanaLogger.Log.Error("You need to set GOPATH environment variable")
		return ""
	}
	for _, wg := range wgopath {
		wg, _ = path.EvalSymlinks(path.Join(wg, "src", "github.com", "goasana", "framework"))
		filename := path.Join(wg, "asana.go")
		_, err := os.Stat(filename)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			asanaLogger.Log.Error("Error while getting stats of 'asana.go'")
		}
		fd, err := os.Open(filename)
		if err != nil {
			asanaLogger.Log.Error("Error while reading 'asana.go'")
			continue
		}
		reader := bufio.NewReader(fd)
		for {
			byteLine, _, er := reader.ReadLine()
			if er != nil && er != io.EOF {
				return ""
			}
			if er == io.EOF {
				break
			}
			line := string(byteLine)
			s := re.FindStringSubmatch(line)
			if len(s) >= 2 {
				return s[1]
			}
		}

	}
	return "Asana is not installed. Please do consider installing it first: https://github.com/goasana/framework"
}

func GetGoVersion() string {
	var (
		cmdOut []byte
		err    error
	)

	if cmdOut, err = exec.Command("go", "version").Output(); err != nil {
		asanaLogger.Log.Fatalf("There was an error running 'go version' command: %s", err)
	}
	return strings.Split(string(cmdOut), " ")[2]
}
