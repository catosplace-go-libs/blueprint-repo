package main

import (
	"fmt"

	"github.com/catosplace-go-libs/blueprint-repo/pkg/blueprint"
)

func main() {
	// Define the input string with the format [[ Repo-Template-Name ]]
	input := "This is a [[ Repo-Template-Name ]] template."

	// Define the title to replace the placeholder
	title := "GitHub Blueprint"

	// Replace the placeholder with the given title
	result := blueprint.ReplaceTitle(input, title)

	// Print the result
	fmt.Println(result)
	// Output: This is a GitHub Blueprint template.
}