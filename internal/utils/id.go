package utils

import "github.com/jaevor/go-nanoid"

// Generate random a 12-character ID.
func GenerateID() string {
	canonicID, err := nanoid.Standard(12)
	if err != nil {
		panic(err)
	}

	return canonicID()
}
