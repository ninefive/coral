package version

import (
	"fmt"
	"runtime"
)

var (
	gitTag string = ""
	//sha1 from git,output of $(git rev-parse HEAD)
	gitCommit string = "$Format:%H$"
	//state of git tree,either "clean" or "dirty"
	gitTreeState string = "not a git tree"
	//build date in ISO08601 format
	//output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
	buildDate string = "1970-01-01T00:00:00Z"
)

type Info struct {
	GitTag       string `json:"gitTag"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

func (info Info) String() string {
	return info.GitTag
}

func GetInfo() Info {
	return Info{
		GitTag:       gitTag,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
