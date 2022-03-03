package tunnel

import (
	"encoding/xml"

	"github.com/willhu-commit/pango/util"
	"github.com/willhu-commit/pango/version"
)

// Entry is a normalized, version independent representation of
// a VLAN interface.
type Entry struct {
	Name              string
	Comment           string
	NetflowProfile    string
	StaticIps         []string // ordered
	ManagementProfile string
	Mtu               int
	EnableIpv6        bool
	Ipv6InterfaceId   string

	raw map[string]string
}

// Copy copies the information from source Entry `s` to this object.  As the
// Name field relates to the XPATH of this object, this field is not copied.
func (o *Entry) Copy(s Entry) {
	o.Comment = s.Comment
	o.NetflowProfile = s.NetflowProfile
	o.StaticIps = s.StaticIps
	o.ManagementProfile = s.ManagementProfile
	o.Mtu = s.Mtu
	o.EnableIpv6 = s.EnableIpv6
	o.Ipv6InterfaceId = s.Ipv6InterfaceId
}

/** Structs / functions for this namespace. **/

func (o Entry) Specify(v version.Number) (string, string, interface{}) {
	_, fn := versioning(v)
	return o.Name, o.Name, fn(o)
}

type normalizer interface {
	Normalize() []Entry
	Names() []string
}

type container_v1 struct {
	Answer []entry_v1 `xml:"entry"`
}

func (o *container_v1) Normalize() []Entry {
	ans := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].normalize())
	}

	return ans
}

func (o *container_v1) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *entry_v1) normalize() Entry {
	ans := Entry{
		Name:              o.Name,
		Comment:           o.Comment,
		NetflowProfile:    o.NetflowProfile,
		StaticIps:         util.EntToStr(o.StaticIps),
		Mtu:               int(o.Mtu),
		ManagementProfile: o.ManagementProfile,
	}

	raw := make(map[string]string)

	if o.Ipv6 != nil {
		ans.EnableIpv6 = util.AsBool(o.Ipv6.EnableIpv6)
		ans.Ipv6InterfaceId = o.Ipv6.Ipv6InterfaceId

		if o.Ipv6.Addresses != nil {
			raw["v6a"] = util.CleanRawXml(o.Ipv6.Addresses.Text)
		}
	}

	if len(raw) != 0 {
		ans.raw = raw
	}

	return ans
}

type entry_v1 struct {
	XMLName           xml.Name        `xml:"entry"`
	Name              string          `xml:"name,attr"`
	Comment           string          `xml:"comment,omitempty"`
	NetflowProfile    string          `xml:"netflow-profile,omitempty"`
	StaticIps         *util.EntryType `xml:"ip"`
	Mtu               int             `xml:"mtu,omitempty"`
	ManagementProfile string          `xml:"interface-management-profile,omitempty"`
	Ipv6              *ipv6           `xml:"ipv6"`
}

type ipv6 struct {
	EnableIpv6      string       `xml:"enabled"`
	Ipv6InterfaceId string       `xml:"interface-id,omitempty"`
	Addresses       *util.RawXml `xml:"address"`
}

func specify_v1(e Entry) interface{} {
	ans := entry_v1{
		Name:              e.Name,
		Comment:           e.Comment,
		NetflowProfile:    e.NetflowProfile,
		StaticIps:         util.StrToEnt(e.StaticIps),
		Mtu:               e.Mtu,
		ManagementProfile: e.ManagementProfile,
	}

	if e.raw["v6a"] != "" || e.EnableIpv6 || e.Ipv6InterfaceId != "" {
		ans.Ipv6 = &ipv6{
			EnableIpv6:      util.YesNo(e.EnableIpv6),
			Ipv6InterfaceId: e.Ipv6InterfaceId,
		}

		if text := e.raw["v6a"]; text != "" {
			ans.Ipv6.Addresses = &util.RawXml{text}
		}
	}

	return ans
}