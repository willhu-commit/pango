package dev

import (
	"github.com/willhu-commit/pango/util"

	cert "github.com/willhu-commit/pango/dev/certificate"
	"github.com/willhu-commit/pango/dev/general"
	"github.com/willhu-commit/pango/dev/ha"
	halink "github.com/willhu-commit/pango/dev/ha/monitor/link"
	hapath "github.com/willhu-commit/pango/dev/ha/monitor/path"
	"github.com/willhu-commit/pango/dev/localuserdb/group"
	"github.com/willhu-commit/pango/dev/localuserdb/user"
	"github.com/willhu-commit/pango/dev/profile/certificate"
	"github.com/willhu-commit/pango/dev/profile/email"
	"github.com/willhu-commit/pango/dev/profile/http"
	"github.com/willhu-commit/pango/dev/profile/snmp"
	"github.com/willhu-commit/pango/dev/profile/syslog"
	"github.com/willhu-commit/pango/dev/ssldecrypt"
	"github.com/willhu-commit/pango/dev/telemetry"
	"github.com/willhu-commit/pango/dev/vminfosource"
)

// Firewall is the client.Device namespace.
type Firewall struct {
	Certificate         *cert.Firewall
	CertificateProfile  *certificate.Firewall
	EmailServerProfile  *email.Firewall
	GeneralSettings     *general.Firewall
	HaConfig            *ha.Firewall
	HaLinkMonitorGroup  *halink.Firewall
	HaPathMonitorGroup  *hapath.Firewall
	HttpServerProfile   *http.Firewall
	LocalUserDbGroup    *group.Firewall
	LocalUserDbUser     *user.Firewall
	SslDecrypt          *ssldecrypt.Firewall
	SnmpServerProfile   *snmp.Firewall
	SyslogServerProfile *syslog.Firewall
	Telemetry           *telemetry.Firewall
	VmInfoSource        *vminfosource.Firewall
}

func FirewallNamespace(x util.XapiClient) *Firewall {
	return &Firewall{
		Certificate:         cert.FirewallNamespace(x),
		CertificateProfile:  certificate.FirewallNamespace(x),
		EmailServerProfile:  email.FirewallNamespace(x),
		GeneralSettings:     general.FirewallNamespace(x),
		HaConfig:            ha.FirewallNamespace(x),
		HaLinkMonitorGroup:  halink.FirewallNamespace(x),
		HaPathMonitorGroup:  hapath.FirewallNamespace(x),
		HttpServerProfile:   http.FirewallNamespace(x),
		LocalUserDbGroup:    group.FirewallNamespace(x),
		LocalUserDbUser:     user.FirewallNamespace(x),
		SslDecrypt:          ssldecrypt.FirewallNamespace(x),
		SnmpServerProfile:   snmp.FirewallNamespace(x),
		SyslogServerProfile: syslog.FirewallNamespace(x),
		Telemetry:           telemetry.FirewallNamespace(x),
		VmInfoSource:        vminfosource.FirewallNamespace(x),
	}
}
