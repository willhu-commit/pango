package predefined

import (
	"github.com/willhu-commit/pango/objs/app"
	"github.com/willhu-commit/pango/objs/srvc"
	dlpft "github.com/willhu-commit/pango/predefined/dlp/filetype"
	tdbft "github.com/willhu-commit/pango/predefined/tdb/filetype"
	"github.com/willhu-commit/pango/predefined/threat"
	"github.com/willhu-commit/pango/util"
)

type Firewall struct {
	Application *app.Predefined
	DlpFileType *dlpft.Firewall
	Services    *srvc.Predefined
	TdbFileType *tdbft.Firewall
	Threat      *threat.Firewall
}

func FirewallNamespace(x util.XapiClient) *Firewall {
	return &Firewall{
		Application: app.PredefinedNamespace(x),
		DlpFileType: dlpft.FirewallNamespace(x),
		Services:    srvc.PredefinedNamespace(x),
		TdbFileType: tdbft.FirewallNamespace(x),
		Threat:      threat.FirewallNamespace(x),
	}
}
