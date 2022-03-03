package predefined

import (
	"github.com/willhu-commit/pango/objs/app"
	"github.com/willhu-commit/pango/objs/srvc"
	dlpft "github.com/willhu-commit/pango/predefined/dlp/filetype"
	tdbft "github.com/willhu-commit/pango/predefined/tdb/filetype"
	"github.com/willhu-commit/pango/predefined/threat"
	"github.com/willhu-commit/pango/util"
)

type Panorama struct {
	Application *app.Predefined
	DlpFileType *dlpft.Panorama
	Services    *srvc.Predefined
	TdbFileType *tdbft.Panorama
	Threat      *threat.Panorama
}

func PanoramaNamespace(x util.XapiClient) *Panorama {
	return &Panorama{
		Application: app.PredefinedNamespace(x),
		DlpFileType: dlpft.PanoramaNamespace(x),
		Services:    srvc.PredefinedNamespace(x),
		TdbFileType: tdbft.PanoramaNamespace(x),
		Threat:      threat.PanoramaNamespace(x),
	}
}
