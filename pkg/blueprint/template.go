package blueprint

import (
	"github.com/valyala/fasttemplate"
)

// ReplaceTitle takes an input string with the format [[ Repo-Template-Name ]] and replaces it with the given title
func ReplaceTitle(input, title string) string {
	// Define the start and end markers for the template.
	tmpl := fasttemplate.New(input, "[[ ", " ]]")

	// Replace the Title placeholder with the given title
	result := tmpl.ExecuteString(map[string]interface{}{
		"Repo-Template-Name": title,
	})

	return result
}