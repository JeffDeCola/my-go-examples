// my-go-examples complex_function.go

package monkey

import "log"

type (
	version struct {
		Ref string `json:"ref"`
	}
	// InputJSON ...
	InputJSON struct {
		Params  map[string]string `json:"params"`
		Source  map[string]string `json:"source"`
		Version version           `json:"version"`
	}

	checkOutputJSON []version
)

func getversions() []string {

	return []string{
		"123",
		"3de",
		"456",
		"777",
	}

}

// Check will return the NEW versions of a resource.
func Check(input InputJSON, logger *log.Logger) (checkOutputJSON, error) {

	var output = checkOutputJSON{}
	for _, ver := range getversions() {
		output = append(output, version{Ref: ver})
	}

	return output, nil

}
