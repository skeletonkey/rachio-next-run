package healthcheck

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime/debug"
)

func Healthcheck() {
	config := getConfig()
	http.HandleFunc(config.EndPoint, hcResponse)
	err := http.ListenAndServe(config.Addr, nil)
	if err != nil {
		panic(err)
	}
}

func hcResponse(w http.ResponseWriter, r *http.Request) {
	config := getConfig()
	hcData := make(map[string]map[string]string, 0)
	hcData["App"] = make(map[string]string)
	buildInfo, ok := debug.ReadBuildInfo()
	if ok {
		hcData["App"]["GoVersion"] = buildInfo.GoVersion
		if config.Visible.Dependencies {
			hcData["Dependencies"] = make(map[string]string)
			for _, v := range buildInfo.Deps {
				hcData["Dependencies"][v.Path] = fmt.Sprintf("%s:%s", v.Version, v.Sum)
			}
		}
		if config.Visible.Settings {
			hcData["Settings"] = make(map[string]string)
			for _, v := range buildInfo.Settings {
				hcData["Settings"][v.Key] = v.Value
			}
		}
	}
	hc, err := json.Marshal(hcData)
	if err != nil {
		panic(err)
	}
	io.WriteString(w, string(hc))
}
