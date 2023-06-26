package output

import (
	"os"
	"strings"
)

func WriteToFile(modulesToBuild []string) error {
	modulesToBuildString := strings.Join(modulesToBuild, ",")

	err := os.WriteFile("/tmp/build.dag", []byte(modulesToBuildString), 0400)
	if err != nil {
		return err
	}

	return nil
}
