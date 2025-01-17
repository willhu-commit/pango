package syslog

import (
	"encoding/xml"

	"github.com/willhu-commit/pango/version"
)

// Entry is a normalized, version independent representation of an syslog profile.
//
// PAN-OS 7.1+.
type Entry struct {
	Name              string
	Config            string
	System            string
	Threat            string
	Traffic           string
	HipMatch          string
	Url               string // 8.0+
	Data              string // 8.0+
	Wildfire          string // 8.0+
	Tunnel            string // 8.0+
	UserId            string // 8.0+
	Gtp               string // 8.0+
	Auth              string // 8.0+
	Sctp              string // 8.1+
	Iptag             string // 9.0+
	EscapedCharacters string
	EscapeCharacter   string
	Servers           []Server
}

// Server is a syslog server.
type Server struct {
	Name         string
	Server       string
	Transport    string
	Port         int
	SyslogFormat string
	Facility     string
}

// Copy copies the information from source Entry `s` to this object.  As the
// Name field relates to the XPATH of this object, this field is not copied.
func (o *Entry) Copy(s Entry) {
	o.Config = s.Config
	o.System = s.System
	o.Threat = s.Threat
	o.Traffic = s.Traffic
	o.HipMatch = s.HipMatch
	o.Url = s.Url
	o.Data = s.Data
	o.Wildfire = s.Wildfire
	o.Tunnel = s.Tunnel
	o.UserId = s.UserId
	o.Gtp = s.Gtp
	o.Auth = s.Auth
	o.Sctp = s.Sctp
	o.Iptag = s.Iptag
	o.EscapedCharacters = s.EscapedCharacters
	o.EscapeCharacter = s.EscapeCharacter
	if s.Servers == nil {
		o.Servers = nil
	} else {
		list := make([]Server, 0, len(s.Servers))
		for _, x := range s.Servers {
			list = append(list, Server{
				Name:         x.Name,
				Server:       x.Server,
				Transport:    x.Transport,
				Port:         x.Port,
				SyslogFormat: x.SyslogFormat,
				Facility:     x.Facility,
			})
		}
		o.Servers = list
	}
}

/** Structs / functions for this namespace. **/

func (o Entry) Specify(v version.Number) (string, interface{}) {
	_, fn := versioning(v)
	return o.Name, fn(o)
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
		Name: o.Name,
	}

	if o.Server != nil {
		list := make([]Server, 0, len(o.Server.Entries))
		for _, x := range o.Server.Entries {
			list = append(list, Server{
				Name:         x.Name,
				Server:       x.Server,
				Transport:    x.Transport,
				Port:         x.Port,
				SyslogFormat: x.SyslogFormat,
				Facility:     x.Facility,
			})
		}
		ans.Servers = list
	}

	if o.Format != nil {
		ans.Config = o.Format.Config
		ans.System = o.Format.System
		ans.Threat = o.Format.Threat
		ans.Traffic = o.Format.Traffic
		ans.HipMatch = o.Format.HipMatch

		if o.Format.Esc != nil {
			ans.EscapedCharacters = o.Format.Esc.EscapedCharacters
			ans.EscapeCharacter = o.Format.Esc.EscapeCharacter
		}
	}

	return ans
}

type entry_v1 struct {
	XMLName xml.Name   `xml:"entry"`
	Name    string     `xml:"name,attr"`
	Server  *servers   `xml:"server"`
	Format  *format_v1 `xml:"format"`
}

type servers struct {
	Entries []serverEntry `xml:"entry"`
}

type serverEntry struct {
	XMLName      xml.Name `xml:"entry"`
	Name         string   `xml:"name,attr"`
	Server       string   `xml:"server"`
	Transport    string   `xml:"transport,omitempty"`
	Port         int      `xml:"port,omitempty"`
	SyslogFormat string   `xml:"format,omitempty"`
	Facility     string   `xml:"facility,omitempty"`
}

func (e *serverEntry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type local serverEntry
	ans := local{
		Transport:    TransportUdp,
		SyslogFormat: SyslogFormatBsd,
	}
	if err := d.DecodeElement(&ans, &start); err != nil {
		return err
	}
	*e = serverEntry(ans)
	return nil
}

type format_v1 struct {
	Config   string `xml:"config,omitempty"`
	System   string `xml:"system,omitempty"`
	Threat   string `xml:"thread,omitempty"`
	Traffic  string `xml:"traffic,omitempty"`
	HipMatch string `xml:"hip-match,omitempty"`
	Esc      *esc   `xml:"escaping"`
}

type esc struct {
	EscapedCharacters string `xml:"escaped-characters"`
	EscapeCharacter   string `xml:"escape-character"`
}

func specify_v1(e Entry) interface{} {
	ans := entry_v1{
		Name: e.Name,
	}

	if len(e.Servers) > 0 {
		list := make([]serverEntry, 0, len(e.Servers))
		for _, x := range e.Servers {
			list = append(list, serverEntry{
				Name:         x.Name,
				Server:       x.Server,
				Transport:    x.Transport,
				Port:         x.Port,
				SyslogFormat: x.SyslogFormat,
				Facility:     x.Facility,
			})
		}
		ans.Server = &servers{Entries: list}
	}

	hasEsc := e.EscapedCharacters != "" || e.EscapeCharacter != ""
	if e.Config != "" || e.System != "" || e.Threat != "" || e.Traffic != "" || e.HipMatch != "" || hasEsc {
		ans.Format = &format_v1{
			Config:   e.Config,
			System:   e.System,
			Threat:   e.Threat,
			Traffic:  e.Traffic,
			HipMatch: e.HipMatch,
		}

		if hasEsc {
			ans.Format.Esc = &esc{
				EscapedCharacters: e.EscapedCharacters,
				EscapeCharacter:   e.EscapeCharacter,
			}
		}
	}

	return ans
}

// PAN-OS 8.0
type container_v2 struct {
	Answer []entry_v2 `xml:"entry"`
}

func (o *container_v2) Normalize() []Entry {
	ans := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].normalize())
	}

	return ans
}

func (o *container_v2) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *entry_v2) normalize() Entry {
	ans := Entry{
		Name: o.Name,
	}

	if o.Server != nil {
		list := make([]Server, 0, len(o.Server.Entries))
		for _, x := range o.Server.Entries {
			list = append(list, Server{
				Name:         x.Name,
				Server:       x.Server,
				Transport:    x.Transport,
				Port:         x.Port,
				SyslogFormat: x.SyslogFormat,
				Facility:     x.Facility,
			})
		}
		ans.Servers = list
	}

	if o.Format != nil {
		ans.Config = o.Format.Config
		ans.System = o.Format.System
		ans.Threat = o.Format.Threat
		ans.Traffic = o.Format.Traffic
		ans.HipMatch = o.Format.HipMatch
		ans.Url = o.Format.Url
		ans.Data = o.Format.Data
		ans.Wildfire = o.Format.Wildfire
		ans.Tunnel = o.Format.Tunnel
		ans.UserId = o.Format.UserId
		ans.Gtp = o.Format.Gtp
		ans.Auth = o.Format.Auth

		if o.Format.Esc != nil {
			ans.EscapedCharacters = o.Format.Esc.EscapedCharacters
			ans.EscapeCharacter = o.Format.Esc.EscapeCharacter
		}
	}

	return ans
}

type entry_v2 struct {
	XMLName xml.Name   `xml:"entry"`
	Name    string     `xml:"name,attr"`
	Server  *servers   `xml:"server"`
	Format  *format_v2 `xml:"format"`
}

type format_v2 struct {
	Config   string `xml:"config,omitempty"`
	System   string `xml:"system,omitempty"`
	Threat   string `xml:"thread,omitempty"`
	Traffic  string `xml:"traffic,omitempty"`
	HipMatch string `xml:"hip-match,omitempty"`
	Url      string `xml:"url,omitempty"`
	Data     string `xml:"data,omitempty"`
	Wildfire string `xml:"wildfire,omitempty"`
	Tunnel   string `xml:"tunnel,omitempty"`
	UserId   string `xml:"userid,omitempty"`
	Gtp      string `xml:"gtp,omitempty"`
	Auth     string `xml:"auth,omitempty"`
	Esc      *esc   `xml:"escaping"`
}

func specify_v2(e Entry) interface{} {
	ans := entry_v2{
		Name: e.Name,
	}

	if len(e.Servers) > 0 {
		list := make([]serverEntry, 0, len(e.Servers))
		for _, x := range e.Servers {
			list = append(list, serverEntry{
				Name:         x.Name,
				Server:       x.Server,
				Transport:    x.Transport,
				Port:         x.Port,
				SyslogFormat: x.SyslogFormat,
				Facility:     x.Facility,
			})
		}
		ans.Server = &servers{Entries: list}
	}

	hasEsc := e.EscapedCharacters != "" || e.EscapeCharacter != ""
	if e.Config != "" || e.System != "" || e.Threat != "" || e.Traffic != "" || e.HipMatch != "" || e.Url != "" || e.Data != "" || e.Wildfire != "" || e.Tunnel != "" || e.UserId != "" || e.Gtp != "" || e.Auth != "" || hasEsc {
		ans.Format = &format_v2{
			Config:   e.Config,
			System:   e.System,
			Threat:   e.Threat,
			Traffic:  e.Traffic,
			HipMatch: e.HipMatch,
			Url:      e.Url,
			Data:     e.Data,
			Wildfire: e.Wildfire,
			Tunnel:   e.Tunnel,
			UserId:   e.UserId,
			Gtp:      e.Gtp,
			Auth:     e.Auth,
		}

		if hasEsc {
			ans.Format.Esc = &esc{
				EscapedCharacters: e.EscapedCharacters,
				EscapeCharacter:   e.EscapeCharacter,
			}
		}
	}

	return ans
}

// PAN-OS 8.1
type container_v3 struct {
	Answer []entry_v3 `xml:"entry"`
}

func (o *container_v3) Normalize() []Entry {
	ans := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].normalize())
	}

	return ans
}

func (o *container_v3) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *entry_v3) normalize() Entry {
	ans := Entry{
		Name: o.Name,
	}

	if o.Server != nil {
		list := make([]Server, 0, len(o.Server.Entries))
		for _, x := range o.Server.Entries {
			list = append(list, Server{
				Name:         x.Name,
				Server:       x.Server,
				Transport:    x.Transport,
				Port:         x.Port,
				SyslogFormat: x.SyslogFormat,
				Facility:     x.Facility,
			})
		}
		ans.Servers = list
	}

	if o.Format != nil {
		ans.Config = o.Format.Config
		ans.System = o.Format.System
		ans.Threat = o.Format.Threat
		ans.Traffic = o.Format.Traffic
		ans.HipMatch = o.Format.HipMatch
		ans.Url = o.Format.Url
		ans.Data = o.Format.Data
		ans.Wildfire = o.Format.Wildfire
		ans.Tunnel = o.Format.Tunnel
		ans.UserId = o.Format.UserId
		ans.Gtp = o.Format.Gtp
		ans.Auth = o.Format.Auth
		ans.Sctp = o.Format.Sctp

		if o.Format.Esc != nil {
			ans.EscapedCharacters = o.Format.Esc.EscapedCharacters
			ans.EscapeCharacter = o.Format.Esc.EscapeCharacter
		}
	}

	return ans
}

type entry_v3 struct {
	XMLName xml.Name   `xml:"entry"`
	Name    string     `xml:"name,attr"`
	Server  *servers   `xml:"server"`
	Format  *format_v3 `xml:"format"`
}

type format_v3 struct {
	Config   string `xml:"config,omitempty"`
	System   string `xml:"system,omitempty"`
	Threat   string `xml:"thread,omitempty"`
	Traffic  string `xml:"traffic,omitempty"`
	HipMatch string `xml:"hip-match,omitempty"`
	Url      string `xml:"url,omitempty"`
	Data     string `xml:"data,omitempty"`
	Wildfire string `xml:"wildfire,omitempty"`
	Tunnel   string `xml:"tunnel,omitempty"`
	UserId   string `xml:"userid,omitempty"`
	Gtp      string `xml:"gtp,omitempty"`
	Auth     string `xml:"auth,omitempty"`
	Sctp     string `xml:"sctp,omitempty"`
	Esc      *esc   `xml:"escaping"`
}

func specify_v3(e Entry) interface{} {
	ans := entry_v3{
		Name: e.Name,
	}

	if len(e.Servers) > 0 {
		list := make([]serverEntry, 0, len(e.Servers))
		for _, x := range e.Servers {
			list = append(list, serverEntry{
				Name:         x.Name,
				Server:       x.Server,
				Transport:    x.Transport,
				Port:         x.Port,
				SyslogFormat: x.SyslogFormat,
				Facility:     x.Facility,
			})
		}
		ans.Server = &servers{Entries: list}
	}

	hasEsc := e.EscapedCharacters != "" || e.EscapeCharacter != ""
	if e.Config != "" || e.System != "" || e.Threat != "" || e.Traffic != "" || e.HipMatch != "" || e.Url != "" || e.Data != "" || e.Wildfire != "" || e.Tunnel != "" || e.UserId != "" || e.Gtp != "" || e.Auth != "" || e.Sctp != "" || hasEsc {
		ans.Format = &format_v3{
			Config:   e.Config,
			System:   e.System,
			Threat:   e.Threat,
			Traffic:  e.Traffic,
			HipMatch: e.HipMatch,
			Url:      e.Url,
			Data:     e.Data,
			Wildfire: e.Wildfire,
			Tunnel:   e.Tunnel,
			UserId:   e.UserId,
			Gtp:      e.Gtp,
			Auth:     e.Auth,
			Sctp:     e.Sctp,
		}

		if hasEsc {
			ans.Format.Esc = &esc{
				EscapedCharacters: e.EscapedCharacters,
				EscapeCharacter:   e.EscapeCharacter,
			}
		}
	}

	return ans
}

// PAN-OS 9.0
type container_v4 struct {
	Answer []entry_v4 `xml:"entry"`
}

func (o *container_v4) Normalize() []Entry {
	ans := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].normalize())
	}

	return ans
}

func (o *container_v4) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *entry_v4) normalize() Entry {
	ans := Entry{
		Name: o.Name,
	}

	if o.Server != nil {
		list := make([]Server, 0, len(o.Server.Entries))
		for _, x := range o.Server.Entries {
			list = append(list, Server{
				Name:         x.Name,
				Server:       x.Server,
				Transport:    x.Transport,
				Port:         x.Port,
				SyslogFormat: x.SyslogFormat,
				Facility:     x.Facility,
			})
		}
		ans.Servers = list
	}

	if o.Format != nil {
		ans.Config = o.Format.Config
		ans.System = o.Format.System
		ans.Threat = o.Format.Threat
		ans.Traffic = o.Format.Traffic
		ans.HipMatch = o.Format.HipMatch
		ans.Url = o.Format.Url
		ans.Data = o.Format.Data
		ans.Wildfire = o.Format.Wildfire
		ans.Tunnel = o.Format.Tunnel
		ans.UserId = o.Format.UserId
		ans.Gtp = o.Format.Gtp
		ans.Auth = o.Format.Auth
		ans.Sctp = o.Format.Sctp
		ans.Iptag = o.Format.Iptag

		if o.Format.Esc != nil {
			ans.EscapedCharacters = o.Format.Esc.EscapedCharacters
			ans.EscapeCharacter = o.Format.Esc.EscapeCharacter
		}
	}

	return ans
}

type entry_v4 struct {
	XMLName xml.Name   `xml:"entry"`
	Name    string     `xml:"name,attr"`
	Server  *servers   `xml:"server"`
	Format  *format_v4 `xml:"format"`
}

type format_v4 struct {
	Config   string `xml:"config,omitempty"`
	System   string `xml:"system,omitempty"`
	Threat   string `xml:"thread,omitempty"`
	Traffic  string `xml:"traffic,omitempty"`
	HipMatch string `xml:"hip-match,omitempty"`
	Url      string `xml:"url,omitempty"`
	Data     string `xml:"data,omitempty"`
	Wildfire string `xml:"wildfire,omitempty"`
	Tunnel   string `xml:"tunnel,omitempty"`
	UserId   string `xml:"userid,omitempty"`
	Gtp      string `xml:"gtp,omitempty"`
	Auth     string `xml:"auth,omitempty"`
	Sctp     string `xml:"sctp,omitempty"`
	Iptag    string `xml:"iptag,omitempty"`
	Esc      *esc   `xml:"escaping"`
}

func specify_v4(e Entry) interface{} {
	ans := entry_v4{
		Name: e.Name,
	}

	if len(e.Servers) > 0 {
		list := make([]serverEntry, 0, len(e.Servers))
		for _, x := range e.Servers {
			list = append(list, serverEntry{
				Name:         x.Name,
				Server:       x.Server,
				Transport:    x.Transport,
				Port:         x.Port,
				SyslogFormat: x.SyslogFormat,
				Facility:     x.Facility,
			})
		}
		ans.Server = &servers{Entries: list}
	}

	hasEsc := e.EscapedCharacters != "" || e.EscapeCharacter != ""
	if e.Config != "" || e.System != "" || e.Threat != "" || e.Traffic != "" || e.HipMatch != "" || e.Url != "" || e.Data != "" || e.Wildfire != "" || e.Tunnel != "" || e.UserId != "" || e.Gtp != "" || e.Auth != "" || e.Sctp != "" || e.Iptag != "" || hasEsc {
		ans.Format = &format_v4{
			Config:   e.Config,
			System:   e.System,
			Threat:   e.Threat,
			Traffic:  e.Traffic,
			HipMatch: e.HipMatch,
			Url:      e.Url,
			Data:     e.Data,
			Wildfire: e.Wildfire,
			Tunnel:   e.Tunnel,
			UserId:   e.UserId,
			Gtp:      e.Gtp,
			Auth:     e.Auth,
			Sctp:     e.Sctp,
			Iptag:    e.Iptag,
		}

		if hasEsc {
			ans.Format.Esc = &esc{
				EscapedCharacters: e.EscapedCharacters,
				EscapeCharacter:   e.EscapeCharacter,
			}
		}
	}

	return ans
}
