package client

import (
	"errors"
	"regexp"
	"strings"
)

var imageRegex = regexp.MustCompile(`^\S*:\S*$`)

// Search looks for a an image. An error is returned
// if it cannot be found.
func Search(image string) (string, error) {
	image = strings.TrimSpace(image)

	// There's a special case where just a colon can be
	// provided, this fixes that
	if image == ":" {
		return "", errors.New("improper image format")
	}

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
