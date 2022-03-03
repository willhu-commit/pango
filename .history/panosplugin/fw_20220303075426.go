package panosplugin

import (
	"github.com/willhu-commit/pango/panosplugin/cloudwatch"
	"github.com/willhu-commit/pango/util"
)

type Firewall struct {
	AwsCloudWatch *cloudwatch.Firewall
}

func FirewallNamespace(x util.XapiClient) *Firewall {
	return &Firewall{
		AwsCloudWatch: cloudwatch.FirewallNamespace(x),
	}
}
