package client

import (
	"regexp"
)

var imageRegex = regexp.MustCompile(`^\S*:\S*$`)

// Search looks for a an image. An error is returned
// if it cannot be found.
func Search(image string) (string, error) {
	// If the image text didn't have a tag with it, add
	// :latest
	if !imageRegex.MatchString(image) {
		image += ":latest"
	}

	summaries, err := Repos()
	if err != nil {
		return "", err
	}

	imageFound := false
	for _, summary := range summaries {
		tags, err := Tags(summary.Name)
		if err != nil {
			return "", err
		}

		for _, tag := range tags {
			name := summary.Name + ":" + tag
			if name == image {
				imageFound = true
				break
			}
		}
	}

	if !imageFound {
		return "", ErrImageNotFound
	}

	return image, nil
}
