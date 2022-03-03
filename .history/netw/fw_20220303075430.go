package netw

import (
	"github.com/willhu-commit/pango/netw/dhcp"
	"github.com/willhu-commit/pango/netw/ikegw"
	aggeth "github.com/willhu-commit/pango/netw/interface/aggregate"
	"github.com/willhu-commit/pango/netw/interface/arp"
	"github.com/willhu-commit/pango/netw/interface/eth"
	ipv6a "github.com/willhu-commit/pango/netw/interface/ipv6/address"
	ipv6n "github.com/willhu-commit/pango/netw/interface/ipv6/neighbor"
	"github.com/willhu-commit/pango/netw/interface/loopback"
	"github.com/willhu-commit/pango/netw/interface/subinterface/layer2"
	"github.com/willhu-commit/pango/netw/interface/subinterface/layer3"
	"github.com/willhu-commit/pango/netw/interface/tunnel"
	vli "github.com/willhu-commit/pango/netw/interface/vlan"
	"github.com/willhu-commit/pango/netw/ipsectunnel"
	tpiv4 "github.com/willhu-commit/pango/netw/ipsectunnel/proxyid/ipv4"
	"github.com/willhu-commit/pango/netw/profile/bfd"
	"github.com/willhu-commit/pango/netw/profile/ike"
	"github.com/willhu-commit/pango/netw/profile/ipsec"
	"github.com/willhu-commit/pango/netw/profile/mngtprof"
	"github.com/willhu-commit/pango/netw/profile/monitor"
	redist4 "github.com/willhu-commit/pango/netw/routing/profile/redist/ipv4"
	"github.com/willhu-commit/pango/netw/routing/protocol/bgp"
	"github.com/willhu-commit/pango/netw/routing/protocol/bgp/aggregate"
	agaf "github.com/willhu-commit/pango/netw/routing/protocol/bgp/aggregate/filter/advertise"
	"github.com/willhu-commit/pango/netw/routing/protocol/bgp/aggregate/filter/suppress"
	"github.com/willhu-commit/pango/netw/routing/protocol/bgp/conadv"
	"github.com/willhu-commit/pango/netw/routing/protocol/bgp/conadv/filter/advertise"
	"github.com/willhu-commit/pango/netw/routing/protocol/bgp/conadv/filter/nonexist"
	"github.com/willhu-commit/pango/netw/routing/protocol/bgp/exp"
	"github.com/willhu-commit/pango/netw/routing/protocol/bgp/imp"
	"github.com/willhu-commit/pango/netw/routing/protocol/bgp/peer"
	"github.com/willhu-commit/pango/netw/routing/protocol/bgp/peer/group"
	"github.com/willhu-commit/pango/netw/routing/protocol/bgp/profile/auth"
	"github.com/willhu-commit/pango/netw/routing/protocol/bgp/profile/dampening"
	bgpredist "github.com/willhu-commit/pango/netw/routing/protocol/bgp/redist"
	"github.com/willhu-commit/pango/netw/routing/protocol/ospf"
	ospfarea "github.com/willhu-commit/pango/netw/routing/protocol/ospf/area"
	ospfint "github.com/willhu-commit/pango/netw/routing/protocol/ospf/area/iface"
	ospfvlink "github.com/willhu-commit/pango/netw/routing/protocol/ospf/area/vlink"
	ospfexp "github.com/willhu-commit/pango/netw/routing/protocol/ospf/exp"
	ospfauth "github.com/willhu-commit/pango/netw/routing/protocol/ospf/profile/auth"
	"github.com/willhu-commit/pango/netw/routing/route/static/ipv4"
	ipv6sr "github.com/willhu-commit/pango/netw/routing/route/static/ipv6"
	"github.com/willhu-commit/pango/netw/routing/router"
	"github.com/willhu-commit/pango/netw/tunnel/gre"
	"github.com/willhu-commit/pango/netw/vlan"
	"github.com/willhu-commit/pango/netw/zone"
	"github.com/willhu-commit/pango/util"
)

// Firewall is the client.Network namespace.
type Firewall struct {
	AggregateInterface       *aggeth.Firewall
	Arp                      *arp.Firewall
	BfdProfile               *bfd.Firewall
	BgpAggregate             *aggregate.Firewall
	BgpAggAdvertiseFilter    *agaf.Firewall
	BgpAggSuppressFilter     *suppress.Firewall
	BgpAuthProfile           *auth.Firewall
	BgpConAdvAdvertiseFilter *advertise.Firewall
	BgpConAdvNonExistFilter  *nonexist.Firewall
	BgpConditionalAdv        *conadv.Firewall
	BgpConfig                *bgp.Firewall
	BgpDampeningProfile      *dampening.Firewall
	BgpExport                *exp.Firewall
	BgpImport                *imp.Firewall
	BgpPeer                  *peer.Firewall
	BgpPeerGroup             *group.Firewall
	BgpRedistRule            *bgpredist.Firewall
	Dhcp                     *dhcp.Firewall
	EthernetInterface        *eth.Firewall
	GreTunnel                *gre.Firewall
	IkeCryptoProfile         *ike.Firewall
	IkeGateway               *ikegw.Firewall
	IpsecCryptoProfile       *ipsec.Firewall
	IpsecTunnel              *ipsectunnel.Firewall
	IpsecTunnelProxyId       *tpiv4.Firewall
	Ipv6Address              *ipv6a.Firewall
	Ipv6NeighborDiscovery    *ipv6n.Firewall
	Ipv6StaticRoute          *ipv6sr.Firewall
	Layer2Subinterface       *layer2.Firewall
	Layer3Subinterface       *layer3.Firewall
	LoopbackInterface        *loopback.Firewall
	ManagementProfile        *mngtprof.Firewall
	MonitorProfile           *monitor.Firewall
	OspfArea                 *ospfarea.Firewall
	OspfAreaInterface        *ospfint.Firewall
	OspfAreaVirtualLink      *ospfvlink.Firewall
	OspfAuthProfile          *ospfauth.Firewall
	OspfConfig               *ospf.Firewall
	OspfExport               *ospfexp.Firewall
	RedistributionProfile    *redist4.Firewall
	StaticRoute              *ipv4.Firewall
	TunnelInterface          *tunnel.Firewall
	VirtualRouter            *router.Firewall
	Vlan                     *vlan.Firewall
	VlanInterface            *vli.Firewall
	Zone                     *zone.Firewall
}

func FirewallNamespace(x util.XapiClient) *Firewall {
	return &Firewall{
		AggregateInterface:       aggeth.FirewallNamespace(x),
		Arp:                      arp.FirewallNamespace(x),
		BfdProfile:               bfd.FirewallNamespace(x),
		BgpAggregate:             aggregate.FirewallNamespace(x),
		BgpAggAdvertiseFilter:    agaf.FirewallNamespace(x),
		BgpAggSuppressFilter:     suppress.FirewallNamespace(x),
		BgpAuthProfile:           auth.FirewallNamespace(x),
		BgpConAdvAdvertiseFilter: advertise.FirewallNamespace(x),
		BgpConAdvNonExistFilter:  nonexist.FirewallNamespace(x),
		BgpConditionalAdv:        conadv.FirewallNamespace(x),
		BgpConfig:                bgp.FirewallNamespace(x),
		BgpDampeningProfile:      dampening.FirewallNamespace(x),
		BgpExport:                exp.FirewallNamespace(x),
		BgpImport:                imp.FirewallNamespace(x),
		BgpPeer:                  peer.FirewallNamespace(x),
		BgpPeerGroup:             group.FirewallNamespace(x),
		BgpRedistRule:            bgpredist.FirewallNamespace(x),
		Dhcp:                     dhcp.FirewallNamespace(x),
		EthernetInterface:        eth.FirewallNamespace(x),
		GreTunnel:                gre.FirewallNamespace(x),
		IkeCryptoProfile:         ike.FirewallNamespace(x),
		IkeGateway:               ikegw.FirewallNamespace(x),
		IpsecCryptoProfile:       ipsec.FirewallNamespace(x),
		IpsecTunnel:              ipsectunnel.FirewallNamespace(x),
		IpsecTunnelProxyId:       tpiv4.FirewallNamespace(x),
		Ipv6Address:              ipv6a.FirewallNamespace(x),
		Ipv6NeighborDiscovery:    ipv6n.FirewallNamespace(x),
		Ipv6StaticRoute:          ipv6sr.FirewallNamespace(x),
		Layer2Subinterface:       layer2.FirewallNamespace(x),
		Layer3Subinterface:       layer3.FirewallNamespace(x),
		LoopbackInterface:        loopback.FirewallNamespace(x),
		ManagementProfile:        mngtprof.FirewallNamespace(x),
		MonitorProfile:           monitor.FirewallNamespace(x),
		OspfArea:                 ospfarea.FirewallNamespace(x),
		OspfAreaInterface:        ospfint.FirewallNamespace(x),
		OspfAreaVirtualLink:      ospfvlink.FirewallNamespace(x),
		OspfAuthProfile:          ospfauth.FirewallNamespace(x),
		OspfConfig:               ospf.FirewallNamespace(x),
		OspfExport:               ospfexp.FirewallNamespace(x),
		RedistributionProfile:    redist4.FirewallNamespace(x),
		StaticRoute:              ipv4.FirewallNamespace(x),
		TunnelInterface:          tunnel.FirewallNamespace(x),
		VirtualRouter:            router.FirewallNamespace(x),
		Vlan:                     vlan.FirewallNamespace(x),
		VlanInterface:            vli.FirewallNamespace(x),
		Zone:                     zone.FirewallNamespace(x),
	}
}
