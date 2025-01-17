package variable

import (
	"fmt"

	"github.com/willhu-commit/pango/namespace"
	"github.com/willhu-commit/pango/util"
)

// Panorama is the client.Panorama.TemplateVariable namespace.
type Panorama struct {
	ns *namespace.Standard
}

// GetList performs GET to retrieve a list of all objects.
func (c *Panorama) GetList(tmpl, ts string) ([]string, error) {
	ans := c.container()
	return c.ns.Listing(util.Get, c.pather(tmpl, ts), ans)
}

// ShowList performs a SHOW to retrieve a list of all objects.
func (c *Panorama) ShowList(tmpl, ts string) ([]string, error) {
	ans := c.container()
	return c.ns.Listing(util.Show, c.pather(tmpl, ts), ans)
}

// Get performs GET to retrieve configuration for the given object.
func (c *Panorama) Get(tmpl, ts, name string) (Entry, error) {
	ans := c.container()
	err := c.ns.Object(util.Get, c.pather(tmpl, ts), name, ans)
	return first(ans, err)
}

// Show performs SHOW to retrieve configuration for the given object.
func (c *Panorama) Show(tmpl, ts, name string) (Entry, error) {
	ans := c.container()
	err := c.ns.Object(util.Show, c.pather(tmpl, ts), name, ans)
	return first(ans, err)
}

// GetAll performs GET to retrieve all objects configured.
func (c *Panorama) GetAll(tmpl, ts string) ([]Entry, error) {
	ans := c.container()
	err := c.ns.Objects(util.Get, c.pather(tmpl, ts), ans)
	return all(ans, err)
}

// ShowAll performs SHOW to retrieve all objects configured.
func (c *Panorama) ShowAll(tmpl, ts string) ([]Entry, error) {
	ans := c.container()
	err := c.ns.Objects(util.Show, c.pather(tmpl, ts), ans)
	return all(ans, err)
}

// Set performs SET to configure the specified objects.
func (c *Panorama) Set(tmpl, ts string, e ...Entry) error {
	return c.ns.Set(c.pather(tmpl, ts), specifier(e...))
}

// Edit performs EDIT to configure the specified object.
func (c *Panorama) Edit(tmpl, ts string, e Entry) error {
	return c.ns.Edit(c.pather(tmpl, ts), e)
}

// Delete performs DELETE to remove the specified objects.
//
// Objects can be either a string or an Entry object.
func (c *Panorama) Delete(tmpl, ts string, e ...interface{}) error {
	names, nErr := toNames(e)
	return c.ns.Delete(c.pather(tmpl, ts), names, nErr)
}

func (c *Panorama) pather(tmpl, ts string) namespace.Pather {
	return func(v []string) ([]string, error) {
		return c.xpath(tmpl, ts, v)
	}
}

func (c *Panorama) xpath(tmpl, ts string, vals []string) ([]string, error) {
	if tmpl == "" && ts == "" {
		return nil, fmt.Errorf("tmpl or ts must be specified")
	}

	ans := make([]string, 0, 7)
	ans = append(ans, util.TemplateXpathPrefix(tmpl, ts)...)
	ans = append(ans,
		"variable",
		util.AsEntryXpath(vals),
	)

	return ans, nil
}

func (c *Panorama) container() normalizer {
	return container(c.ns.Client.Versioning())
}
