package poli

import (
	"github.com/willhu-commit/pango/util"

	"github.com/willhu-commit/pango/poli/decryption"
	"github.com/willhu-commit/pango/poli/nat"
	"github.com/willhu-commit/pango/poli/pbf"
	"github.com/willhu-commit/pango/poli/security"
)

// Firewall is the client.Policies namespace.
type Firewall struct {
	Decryption            *decryption.Firewall
	Nat                   *nat.Firewall
	PolicyBasedForwarding *pbf.Firewall
	Security              *security.Firewall
}

func FirewallNamespace(x util.XapiClient) *Firewall {
	return &Firewall{
		Decryption:            decryption.FirewallNamespace(x),
		Nat:                   nat.FirewallNamespace(x),
		PolicyBasedForwarding: pbf.FirewallNamespace(x),
		Security:              security.FirewallNamespace(x),
	}
}
