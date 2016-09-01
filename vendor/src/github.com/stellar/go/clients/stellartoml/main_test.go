package stellartoml

import (
	"fmt"
	"log"
)

// ExampleGetTOML gets the stellar.toml file for stellar.org
func ExampleGetTOML() {
	resp, err := DefaultClient.GetStellarToml("stellar.org")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Signing key: %s\n", resp.SigningKey)
	// Output: Signing key: GDZ2LHRX35XR7PEVVWYRMFP7OMRL7W2X5JGLDUWS6YBRIXPYU553TADF
}
