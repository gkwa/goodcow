package embedded

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type RepositoryInfo struct {
	BrowseURL string    `json:"browse_url"`
	CreatedAt time.Time `json:"created_at"`
	GitCommit string    `json:"git_commit"`
	GitURL    string    `json:"git_url"`
	IndexedAt time.Time `json:"indexed_at"`
	Path      string    `json:"path"`
	Release   string    `json:"release"`
	Subpath   string    `json:"subpath"`
	Version   string    `json:"version"`
}

type CustomizedRepositoryInfo struct {
	RepositoryInfo
	Author        string `json:"author"`
	AuthorRepoURL string `json:"author_url"`
	// Additional fields...
}

func NewCustomizedRepositoryInfo(baseRepoInfo RepositoryInfo, additionalFields map[string]interface{}) CustomizedRepositoryInfo {
	customRepoInfo := CustomizedRepositoryInfo{
		RepositoryInfo: baseRepoInfo,
	}

	customValue := reflect.ValueOf(&customRepoInfo).Elem()

	for fieldName, value := range additionalFields {
		field := customValue.FieldByName(fieldName)
		if field.IsValid() && field.CanSet() {
			field.Set(reflect.ValueOf(value))
		}
	}

	return customRepoInfo
}

func RunEmbedded() error {
	jsonData := `{
	    "browse_url": "https://github.com/fluent-ci-templates/shuttle-pipeline",
	    "created_at": "2024-01-29T08:13:15Z",
	    "git_commit": "0e073730b9f80be6f2c50631840b153e31c331c8",
	    "git_url": "https://github.com/fluent-ci-templates/shuttle-pipeline",
	    "indexed_at": "2024-01-29T08:14:39.002483Z",
	    "path": "github.com/fluent-ci-templates/shuttle-pipeline",
	    "release": "",
	    "version": "0e073730b9f80be6f2c50631840b153e31c331c8"
	}`

	var baseRepoInfo RepositoryInfo

	err := json.Unmarshal([]byte(jsonData), &baseRepoInfo)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %s", err)
	}

	fmt.Printf("%+v\n", baseRepoInfo)

	additionalFields := map[string]interface{}{
		"Author":        "John Doe",
		"AuthorRepoURL": "https://github.com/johndoe",
		// Add more additional fields as needed...
	}

	customRepoInfo := NewCustomizedRepositoryInfo(baseRepoInfo, additionalFields)

	fmt.Println(baseRepoInfo.String())
	fmt.Println(customRepoInfo.String())

	return nil
}

func (r *RepositoryInfo) String() string {
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "Error marshalling JSON: " + err.Error()
	}
	return string(b)
}

func (r *CustomizedRepositoryInfo) String() string {
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "Error marshalling JSON: " + err.Error()
	}
	return string(b)
}
