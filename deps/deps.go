package deps

type DepsMap = map[string]interface{}
type BuildMap = map[string]bool

func GetDepsMap(changedFiles []string) (DepsMap, error) {
	return DepsMap{}, nil
}

func GetModulesToBuild(changedFiles []string, depsMap DepsMap) ([]string, error) {
	return []string{"Hello!"}, nil
}
