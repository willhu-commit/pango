package poli

import (
	"github.com/willhu-commit/pango/util"

	"github.com/willhu-commit/pango/poli/decryption"
	"github.com/willhu-commit/pango/poli/nat"
	"github.com/willhu-commit/pango/poli/pbf"
	"github.com/willhu-commit/pango/poli/security"
)

// Panorama is the client.Policies namespace.
type Panorama struct {
	Decryption            *decryption.Panorama
	Nat                   *nat.Panorama
	PolicyBasedForwarding *pbf.Panorama
	Security              *security.Panorama
}

func PanoramaNamespace(x util.XapiClient) *Panorama {
	return &Panorama{
		Decryption:            decryption.PanoramaNamespace(x),
		Nat:                   nat.PanoramaNamespace(x),
		PolicyBasedForwarding: pbf.PanoramaNamespace(x),
		Security:              security.PanoramaNamespace(x),
	}
}
