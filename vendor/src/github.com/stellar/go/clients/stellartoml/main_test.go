package stellartoml

import "log"

// ExampleGetTOML gets the stellar.toml file for coins.asia
func ExampleGetTOML() {
	_, err := DefaultClient.GetStellarToml("coins.asia")
	if err != nil {
		log.Fatal(err)
	}
}
