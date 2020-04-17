// Contains common state functions.
// https://www.terraform.io/docs/extend/schemas/schema-behaviors.html#statefunc

package google

import (
	"crypto/sha256"
	"encoding/hex"
)

func sha256HashState(val interface{}) string {
	hash := sha256.Sum256([]byte(val.(string)))
	return hex.EncodeToString(hash[:])
}
