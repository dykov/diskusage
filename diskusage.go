package diskusage

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
)

func FileSpaceUsage(path string) (uint64, error) {

	cmd := exec.Command("du", "--bytes", path)

	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	if err := cmd.Run(); err != nil {
		return 0, err
	}

	splitted := strings.Split(string(cmdOutput.Bytes()), "\t")

	diskUsageBytes, err := strconv.ParseUint(splitted[0], 10, 64)
	if err != nil {
		return 0, err
	}

	return diskUsageBytes, nil

}

type FileInfo struct {
	Filesystem        string  `json:"filesystem"`
	Total             uint64  `json:"total"`
	Used              uint64  `json:"used"`
	Available         uint64  `json:"available"`
	UsedPercentString string  `json:"used_percent_string"`
	UsedPercentFloat  float64 `json:"used_percent_float"`
	MountedOn         string  `json:"mounted_on"`
}

func DiskSpaceUsage() (*FileInfo, error) {

	cmd := exec.Command("df", "/")

	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	// get line with values
	valLine := (strings.Split(string(cmdOutput.Bytes()), "\n"))[1]

	// split valLine to string values
	splitted2 := strings.Split(valLine, " ")

	var values []string
	for _, v := range splitted2 {
		if v != "" {
			values = append(values, v)
		}
	}

	filesystem := values[0]
	total, err := strconv.ParseUint(values[1], 10, 64)
	if err != nil {
		return nil, err
	}
	used, err := strconv.ParseUint(values[2], 10, 64)
	if err != nil {
		return nil, err
	}
	available, err := strconv.ParseUint(values[3], 10, 64)
	if err != nil {
		return nil, err
	}
	usedPercentString := values[4]
	usedPercentFloat := (float64(used) / float64(total)) * 100.0
	mountedOn := values[5]

	fileInfo := &FileInfo{
		filesystem,
		total,
		used,
		available,
		usedPercentString,
		usedPercentFloat,
		mountedOn,
	}

	return fileInfo, nil

}
