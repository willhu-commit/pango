package ipsectunnel

import (
	"github.com/willhu-commit/pango/version"
)

type testCase struct {
	desc    string
	version version.Number
	conf    Entry
}

func getTests() []testCase {
	return []testCase{
		{"v1 auto key", version.Number{6, 1, 0, ""}, Entry{
			Name:                       "auto key",
			TunnelInterface:            "tunnel.1",
			AntiReplay:                 true,
			CopyTos:                    false,
			Type:                       TypeAutoKey,
			AkIkeGateway:               "my gateway",
			AkIpsecCryptoProfile:       "my profile",
			EnableTunnelMonitor:        true,
			TunnelMonitorDestinationIp: "10.1.1.1",
			TunnelMonitorSourceIp:      "192.168.1.1",
			TunnelMonitorProfile:       "tun prof",
		}},
		{"v1 manual key esp md5", version.Number{6, 1, 0, ""}, Entry{
			Name:                       "manual key esp md5",
			TunnelInterface:            "tunnel.1",
			AntiReplay:                 false,
			CopyTos:                    true,
			Type:                       TypeManualKey,
			MkLocalSpi:                 "local spi",
			MkRemoteSpi:                "remote spi",
			MkInterface:                "mk interface",
			MkLocalAddressIp:           "10.5.5.5",
			MkRemoteAddress:            "192.168.55.55",
			MkProtocol:                 MkProtocolEsp,
			MkEspEncryptionType:        MkEspEncryptionAes128,
			MkEspEncryptionKey:         "esp key",
			MkAuthType:                 MkAuthTypeMd5,
			MkAuthKey:                  "auth secret",
			EnableTunnelMonitor:        false,
			TunnelMonitorDestinationIp: "10.1.1.1",
			TunnelMonitorSourceIp:      "192.168.1.1",
			TunnelMonitorProfile:       "tun prof",
		}},
		{"v1 manual key esp sha1", version.Number{6, 1, 0, ""}, Entry{
			Name:                "manual key esp sha1",
			TunnelInterface:     "tunnel.1",
			AntiReplay:          true,
			CopyTos:             true,
			Type:                TypeManualKey,
			MkLocalSpi:          "local spi",
			MkRemoteSpi:         "remote spi",
			MkInterface:         "mk interface",
			MkLocalAddressIp:    "10.5.5.5",
			MkRemoteAddress:     "192.168.55.55",
			MkProtocol:          MkProtocolEsp,
			MkEspEncryptionType: MkEspEncryptionDes,
			MkEspEncryptionKey:  "enc key",
			MkAuthType:          MkAuthTypeSha1,
			MkAuthKey:           "auth key",
		}},
		{"v1 manual key esp sha256", version.Number{6, 1, 0, ""}, Entry{
			Name:                "manual key esp sha256",
			TunnelInterface:     "tunnel.1",
			AntiReplay:          true,
			CopyTos:             true,
			Type:                TypeManualKey,
			MkLocalSpi:          "local spi",
			MkRemoteSpi:         "remote spi",
			MkInterface:         "mk interface",
			MkLocalAddressIp:    "10.5.5.5",
			MkRemoteAddress:     "192.168.55.55",
			MkProtocol:          MkProtocolEsp,
			MkEspEncryptionType: MkEspEncryptionDes,
			MkEspEncryptionKey:  "enc key",
			MkAuthType:          MkAuthTypeSha256,
			MkAuthKey:           "auth key",
		}},
		{"v1 manual key esp sha384", version.Number{6, 1, 0, ""}, Entry{
			Name:                "manual key esp sha384",
			TunnelInterface:     "tunnel.1",
			AntiReplay:          true,
			CopyTos:             true,
			Type:                TypeManualKey,
			MkLocalSpi:          "local spi",
			MkRemoteSpi:         "remote spi",
			MkInterface:         "mk interface",
			MkLocalAddressIp:    "10.5.5.5",
			MkRemoteAddress:     "192.168.55.55",
			MkProtocol:          MkProtocolEsp,
			MkEspEncryptionType: MkEspEncryptionAes192,
			MkEspEncryptionKey:  "enc key",
			MkAuthType:          MkAuthTypeSha384,
			MkAuthKey:           "auth key",
		}},
		{"v1 manual key esp sha512", version.Number{6, 1, 0, ""}, Entry{
			Name:                "manual key esp sha512",
			TunnelInterface:     "tunnel.1",
			AntiReplay:          true,
			CopyTos:             true,
			Type:                TypeManualKey,
			MkLocalSpi:          "local spi",
			MkRemoteSpi:         "remote spi",
			MkInterface:         "mk interface",
			MkLocalAddressIp:    "10.5.5.5",
			MkRemoteAddress:     "192.168.55.55",
			MkProtocol:          MkProtocolEsp,
			MkEspEncryptionType: MkEspEncryptionAes256,
			MkEspEncryptionKey:  "enc key",
			MkAuthType:          MkAuthTypeSha512,
			MkAuthKey:           "auth key",
		}},
		{"v1 manual key esp none", version.Number{6, 1, 0, ""}, Entry{
			Name:                "manual key esp none",
			TunnelInterface:     "tunnel.1",
			AntiReplay:          true,
			CopyTos:             true,
			Type:                TypeManualKey,
			MkLocalSpi:          "local spi",
			MkRemoteSpi:         "remote spi",
			MkInterface:         "mk interface",
			MkLocalAddressIp:    "10.5.5.5",
			MkRemoteAddress:     "192.168.55.55",
			MkProtocol:          MkProtocolEsp,
			MkEspEncryptionType: MkEspEncryption3des,
			MkEspEncryptionKey:  "enc key",
			MkAuthType:          MkAuthTypeNone,
		}},
		{"v1 manual key ah md5", version.Number{6, 1, 0, ""}, Entry{
			Name:             "manual key ah md5",
			TunnelInterface:  "tunnel.1",
			AntiReplay:       true,
			CopyTos:          true,
			Type:             TypeManualKey,
			MkLocalSpi:       "local spi",
			MkRemoteSpi:      "remote spi",
			MkInterface:      "mk interface",
			MkLocalAddressIp: "10.5.5.5",
			MkRemoteAddress:  "192.168.55.55",
			MkProtocol:       MkProtocolAh,
			MkAuthType:       MkAuthTypeMd5,
			MkAuthKey:        "auth key",
		}},
		{"v1 manual key ah sha1", version.Number{6, 1, 0, ""}, Entry{
			Name:             "manual key ah sha1",
			TunnelInterface:  "tunnel.1",
			AntiReplay:       true,
			CopyTos:          true,
			Type:             TypeManualKey,
			MkLocalSpi:       "local spi",
			MkRemoteSpi:      "remote spi",
			MkInterface:      "mk interface",
			MkLocalAddressIp: "10.5.5.5",
			MkRemoteAddress:  "192.168.55.55",
			MkProtocol:       MkProtocolAh,
			MkAuthType:       MkAuthTypeSha1,
			MkAuthKey:        "auth key",
		}},
		{"v1 manual key ah sha256", version.Number{6, 1, 0, ""}, Entry{
			Name:             "manual key ah sha256",
			TunnelInterface:  "tunnel.1",
			AntiReplay:       true,
			CopyTos:          true,
			Type:             TypeManualKey,
			MkLocalSpi:       "local spi",
			MkRemoteSpi:      "remote spi",
			MkInterface:      "mk interface",
			MkLocalAddressIp: "10.5.5.5",
			MkRemoteAddress:  "192.168.55.55",
			MkProtocol:       MkProtocolAh,
			MkAuthType:       MkAuthTypeSha256,
			MkAuthKey:        "auth key",
		}},
		{"v1 manual key ah sha384", version.Number{6, 1, 0, ""}, Entry{
			Name:             "manual key ah sha384",
			TunnelInterface:  "tunnel.1",
			AntiReplay:       true,
			CopyTos:          true,
			Type:             TypeManualKey,
			MkLocalSpi:       "local spi",
			MkRemoteSpi:      "remote spi",
			MkInterface:      "mk interface",
			MkLocalAddressIp: "10.5.5.5",
			MkRemoteAddress:  "192.168.55.55",
			MkProtocol:       MkProtocolAh,
			MkAuthType:       MkAuthTypeSha384,
			MkAuthKey:        "auth key",
		}},
		{"v1 manual key ah sha512", version.Number{6, 1, 0, ""}, Entry{
			Name:             "manual key ah sha512",
			TunnelInterface:  "tunnel.1",
			AntiReplay:       true,
			CopyTos:          true,
			Type:             TypeManualKey,
			MkLocalSpi:       "local spi",
			MkRemoteSpi:      "remote spi",
			MkInterface:      "mk interface",
			MkLocalAddressIp: "10.5.5.5",
			MkRemoteAddress:  "192.168.55.55",
			MkProtocol:       MkProtocolAh,
			MkAuthType:       MkAuthTypeSha512,
			MkAuthKey:        "auth key",
		}},
		{"v1 gps no cert", version.Number{6, 1, 0, ""}, Entry{
			Name:                      "gps no cert",
			TunnelInterface:           "tunnel.1",
			AntiReplay:                true,
			CopyTos:                   true,
			Type:                      TypeGlobalProtectSatellite,
			GpsPublishConnectedRoutes: true,
			GpsPublishRoutes:          []string{"route1", "route2"},
			GpsInterface:              "gps interface",
			GpsInterfaceIpIpv4:        "gps ipv4 ip",
		}},
		{"v1 gps with cert", version.Number{6, 1, 0, ""}, Entry{
			Name:                      "gps with cert",
			TunnelInterface:           "tunnel.1",
			AntiReplay:                true,
			CopyTos:                   true,
			Type:                      TypeGlobalProtectSatellite,
			GpsPublishConnectedRoutes: false,
			GpsInterface:              "gps interface",
			GpsInterfaceIpIpv4:        "gps ipv4 ip",
			GpsLocalCertificate:       "my cert",
			GpsCertificateProfile:     "cert prof",
		}},
		{"v2 auto key", version.Number{7, 0, 0, ""}, Entry{
			Name:                       "auto key",
			TunnelInterface:            "tunnel.1",
			AntiReplay:                 true,
			CopyTos:                    false,
			EnableIpv6:                 true,
			Disabled:                   false,
			CopyFlowLabel:              true,
			Type:                       TypeAutoKey,
			AkIkeGateway:               "my gateway",
			AkIpsecCryptoProfile:       "my profile",
			EnableTunnelMonitor:        true,
			TunnelMonitorDestinationIp: "10.1.1.1",
			TunnelMonitorSourceIp:      "192.168.1.1",
			TunnelMonitorProfile:       "tun prof",
		}},
		{"v2 manual key esp md5", version.Number{7, 0, 0, ""}, Entry{
			Name:                       "manual key esp md5",
			TunnelInterface:            "tunnel.1",
			AntiReplay:                 false,
			CopyTos:                    true,
			EnableIpv6:                 false,
			Disabled:                   true,
			CopyFlowLabel:              true,
			Type:                       TypeManualKey,
			MkLocalSpi:                 "local spi",
			MkRemoteSpi:                "remote spi",
			MkInterface:                "mk interface",
			MkLocalAddressIp:           "10.5.5.5",
			MkRemoteAddress:            "192.168.55.55",
			MkProtocol:                 MkProtocolEsp,
			MkEspEncryptionType:        MkEspEncryptionAes128,
			MkEspEncryptionKey:         "esp key",
			MkAuthType:                 MkAuthTypeMd5,
			MkAuthKey:                  "auth secret",
			EnableTunnelMonitor:        false,
			TunnelMonitorDestinationIp: "10.1.1.1",
			TunnelMonitorSourceIp:      "192.168.1.1",
			TunnelMonitorProfile:       "tun prof",
		}},
		{"v2 manual key esp sha1", version.Number{7, 0, 0, ""}, Entry{
			Name:                     "manual key esp sha1",
			TunnelInterface:          "tunnel.1",
			AntiReplay:               true,
			CopyTos:                  true,
			EnableIpv6:               true,
			Disabled:                 true,
			CopyFlowLabel:            true,
			Type:                     TypeManualKey,
			MkLocalSpi:               "local spi",
			MkRemoteSpi:              "remote spi",
			MkInterface:              "mk interface",
			MkLocalAddressFloatingIp: "10.5.5.5",
			MkRemoteAddress:          "192.168.55.55",
			MkProtocol:               MkProtocolEsp,
			MkEspEncryptionType:      MkEspEncryptionDes,
			MkEspEncryptionKey:       "enc key",
			MkAuthType:               MkAuthTypeSha1,
			MkAuthKey:                "auth key",
			TunnelMonitorProxyId:     "tun proxy id",
		}},
		{"v2 manual key esp sha256", version.Number{7, 0, 0, ""}, Entry{
			Name:                "manual key esp sha256",
			TunnelInterface:     "tunnel.1",
			AntiReplay:          true,
			CopyTos:             true,
			Type:                TypeManualKey,
			MkLocalSpi:          "local spi",
			MkRemoteSpi:         "remote spi",
			MkInterface:         "mk interface",
			MkLocalAddressIp:    "10.5.5.5",
			MkRemoteAddress:     "192.168.55.55",
			MkProtocol:          MkProtocolEsp,
			MkEspEncryptionType: MkEspEncryptionDes,
			MkEspEncryptionKey:  "enc key",
			MkAuthType:          MkAuthTypeSha256,
			MkAuthKey:           "auth key",
		}},
		{"v2 manual key esp sha384", version.Number{7, 0, 0, ""}, Entry{
			Name:                "manual key esp sha384",
			TunnelInterface:     "tunnel.1",
			AntiReplay:          true,
			CopyTos:             true,
			Type:                TypeManualKey,
			MkLocalSpi:          "local spi",
			MkRemoteSpi:         "remote spi",
			MkInterface:         "mk interface",
			MkLocalAddressIp:    "10.5.5.5",
			MkRemoteAddress:     "192.168.55.55",
			MkProtocol:          MkProtocolEsp,
			MkEspEncryptionType: MkEspEncryptionAes192,
			MkEspEncryptionKey:  "enc key",
			MkAuthType:          MkAuthTypeSha384,
			MkAuthKey:           "auth key",
		}},
		{"v2 manual key esp sha512", version.Number{7, 0, 0, ""}, Entry{
			Name:                "manual key esp sha512",
			TunnelInterface:     "tunnel.1",
			AntiReplay:          true,
			CopyTos:             true,
			Type:                TypeManualKey,
			MkLocalSpi:          "local spi",
			MkRemoteSpi:         "remote spi",
			MkInterface:         "mk interface",
			MkLocalAddressIp:    "10.5.5.5",
			MkRemoteAddress:     "192.168.55.55",
			MkProtocol:          MkProtocolEsp,
			MkEspEncryptionType: MkEspEncryptionAes256,
			MkEspEncryptionKey:  "enc key",
			MkAuthType:          MkAuthTypeSha512,
			MkAuthKey:           "auth key",
		}},
		{"v2 manual key esp none", version.Number{7, 0, 0, ""}, Entry{
			Name:                "manual key esp none",
			TunnelInterface:     "tunnel.1",
			AntiReplay:          true,
			CopyTos:             true,
			Type:                TypeManualKey,
			MkLocalSpi:          "local spi",
			MkRemoteSpi:         "remote spi",
			MkInterface:         "mk interface",
			MkLocalAddressIp:    "10.5.5.5",
			MkRemoteAddress:     "192.168.55.55",
			MkProtocol:          MkProtocolEsp,
			MkEspEncryptionType: MkEspEncryption3des,
			MkEspEncryptionKey:  "enc key",
			MkAuthType:          MkAuthTypeNone,
		}},
		{"v2 manual key ah md5", version.Number{7, 0, 0, ""}, Entry{
			Name:             "manual key ah md5",
			TunnelInterface:  "tunnel.1",
			AntiReplay:       true,
			CopyTos:          true,
			Type:             TypeManualKey,
			MkLocalSpi:       "local spi",
			MkRemoteSpi:      "remote spi",
			MkInterface:      "mk interface",
			MkLocalAddressIp: "10.5.5.5",
			MkRemoteAddress:  "192.168.55.55",
			MkProtocol:       MkProtocolAh,
			MkAuthType:       MkAuthTypeMd5,
			MkAuthKey:        "auth key",
		}},
		{"v2 manual key ah sha1", version.Number{7, 0, 0, ""}, Entry{
			Name:             "manual key ah sha1",
			TunnelInterface:  "tunnel.1",
			AntiReplay:       true,
			CopyTos:          true,
			Type:             TypeManualKey,
			MkLocalSpi:       "local spi",
			MkRemoteSpi:      "remote spi",
			MkInterface:      "mk interface",
			MkLocalAddressIp: "10.5.5.5",
			MkRemoteAddress:  "192.168.55.55",
			MkProtocol:       MkProtocolAh,
			MkAuthType:       MkAuthTypeSha1,
			MkAuthKey:        "auth key",
		}},
		{"v2 manual key ah sha256", version.Number{7, 0, 0, ""}, Entry{
			Name:             "manual key ah sha256",
			TunnelInterface:  "tunnel.1",
			AntiReplay:       true,
			CopyTos:          true,
			Type:             TypeManualKey,
			MkLocalSpi:       "local spi",
			MkRemoteSpi:      "remote spi",
			MkInterface:      "mk interface",
			MkLocalAddressIp: "10.5.5.5",
			MkRemoteAddress:  "192.168.55.55",
			MkProtocol:       MkProtocolAh,
			MkAuthType:       MkAuthTypeSha256,
			MkAuthKey:        "auth key",
		}},
		{"v2 manual key ah sha384", version.Number{7, 0, 0, ""}, Entry{
			Name:             "manual key ah sha384",
			TunnelInterface:  "tunnel.1",
			AntiReplay:       true,
			CopyTos:          true,
			Type:             TypeManualKey,
			MkLocalSpi:       "local spi",
			MkRemoteSpi:      "remote spi",
			MkInterface:      "mk interface",
			MkLocalAddressIp: "10.5.5.5",
			MkRemoteAddress:  "192.168.55.55",
			MkProtocol:       MkProtocolAh,
			MkAuthType:       MkAuthTypeSha384,
			MkAuthKey:        "auth key",
		}},
		{"v2 manual key ah sha512", version.Number{7, 0, 0, ""}, Entry{
			Name:             "manual key ah sha512",
			TunnelInterface:  "tunnel.1",
			AntiReplay:       true,
			CopyTos:          true,
			Type:             TypeManualKey,
			MkLocalSpi:       "local spi",
			MkRemoteSpi:      "remote spi",
			MkInterface:      "mk interface",
			MkLocalAddressIp: "10.5.5.5",
			MkRemoteAddress:  "192.168.55.55",
			MkProtocol:       MkProtocolAh,
			MkAuthType:       MkAuthTypeSha512,
			MkAuthKey:        "auth key",
		}},
		{"v2 gps no cert", version.Number{7, 0, 0, ""}, Entry{
			Name:                      "gps no cert",
			TunnelInterface:           "tunnel.1",
			AntiReplay:                true,
			CopyTos:                   true,
			Type:                      TypeGlobalProtectSatellite,
			GpsPublishConnectedRoutes: true,
			GpsPublishRoutes:          []string{"route1", "route2"},
			GpsInterface:              "gps interface",
			GpsInterfaceIpIpv4:        "gps ipv4 ip",
		}},
		{"v2 gps with cert", version.Number{7, 0, 0, ""}, Entry{
			Name:                       "gps with cert",
			TunnelInterface:            "tunnel.1",
			AntiReplay:                 true,
			CopyTos:                    true,
			Type:                       TypeGlobalProtectSatellite,
			GpsPublishConnectedRoutes:  false,
			GpsInterface:               "gps interface",
			GpsInterfaceFloatingIpIpv4: "gps ipv4 floating ip",
			GpsLocalCertificate:        "my cert",
			GpsCertificateProfile:      "cert prof",
		}},
		{"v3 auto key", version.Number{8, 0, 0, ""}, Entry{
			Name:                       "auto key",
			TunnelInterface:            "tunnel.1",
			AntiReplay:                 true,
			CopyTos:                    false,
			EnableIpv6:                 true,
			Disabled:                   false,
			CopyFlowLabel:              true,
			Type:                       TypeAutoKey,
			AkIkeGateway:               "my gateway",
			AkIpsecCryptoProfile:       "my profile",
			EnableTunnelMonitor:        true,
			TunnelMonitorDestinationIp: "10.1.1.1",
			TunnelMonitorSourceIp:      "192.168.1.1",
			TunnelMonitorProfile:       "tun prof",
		}},
		{"v3 manual key esp md5", version.Number{8, 0, 0, ""}, Entry{
			Name:                       "manual key esp md5",
			TunnelInterface:            "tunnel.1",
			AntiReplay:                 false,
			CopyTos:                    true,
			EnableIpv6:                 false,
			Disabled:                   true,
			CopyFlowLabel:              true,
			Type:                       TypeManualKey,
			MkLocalSpi:                 "local spi",
			MkRemoteSpi:                "remote spi",
			MkInterface:                "mk interface",
			MkLocalAddressIp:           "10.5.5.5",
			MkRemoteAddress:            "192.168.55.55",
			MkProtocol:                 MkProtocolEsp,
			MkEspEncryptionType:        MkEspEncryptionAes128,
			MkEspEncryptionKey:         "esp key",
			MkAuthType:                 MkAuthTypeMd5,
			MkAuthKey:                  "auth secret",
			EnableTunnelMonitor:        false,
			TunnelMonitorDestinationIp: "10.1.1.1",
			TunnelMonitorSourceIp:      "192.168.1.1",
			TunnelMonitorProfile:       "tun prof",
		}},
		{"v3 manual key esp sha1", version.Number{8, 0, 0, ""}, Entry{
			Name:                     "manual key esp sha1",
			TunnelInterface:          "tunnel.1",
			AntiReplay:               true,
			CopyTos:                  true,
			EnableIpv6:               true,
			Disabled:                 true,
			CopyFlowLabel:            true,
			Type:                     TypeManualKey,
			MkLocalSpi:               "local spi",
			MkRemoteSpi:              "remote spi",
			MkInterface:              "mk interface",
			MkLocalAddressFloatingIp: "10.5.5.5",
			MkRemoteAddress:          "192.168.55.55",
			MkProtocol:               MkProtocolEsp,
			MkEspEncryptionType:      MkEspEncryptionDes,
			MkEspEncryptionKey:       "enc key",
			MkAuthType:               MkAuthTypeSha1,
			MkAuthKey:                "auth key",
			TunnelMonitorProxyId:     "tun proxy id",
		}},
		{"v3 manual key esp sha256", version.Number{8, 0, 0, ""}, Entry{
			Name:                "manual key esp sha256",
			TunnelInterface:     "tunnel.1",
			AntiReplay:          true,
			CopyTos:             true,
			Type:                TypeManualKey,
			MkLocalSpi:          "local spi",
			MkRemoteSpi:         "remote spi",
			MkInterface:         "mk interface",
			MkLocalAddressIp:    "10.5.5.5",
			MkRemoteAddress:     "192.168.55.55",
			MkProtocol:          MkProtocolEsp,
			MkEspEncryptionType: MkEspEncryptionDes,
			MkEspEncryptionKey:  "enc key",
			MkAuthType:          MkAuthTypeSha256,
			MkAuthKey:           "auth key",
		}},
		{"v3 manual key esp sha384", version.Number{8, 0, 0, ""}, Entry{
			Name:                "manual key esp sha384",
			TunnelInterface:     "tunnel.1",
			AntiReplay:          true,
			CopyTos:             true,
			Type:                TypeManualKey,
			MkLocalSpi:          "local spi",
			MkRemoteSpi:         "remote spi",
			MkInterface:         "mk interface",
			MkLocalAddressIp:    "10.5.5.5",
			MkRemoteAddress:     "192.168.55.55",
			MkProtocol:          MkProtocolEsp,
			MkEspEncryptionType: MkEspEncryptionAes192,
			MkEspEncryptionKey:  "enc key",
			MkAuthType:          MkAuthTypeSha384,
			MkAuthKey:           "auth key",
		}},
		{"v3 manual key esp sha512", version.Number{8, 0, 0, ""}, Entry{
			Name:                "manual key esp sha512",
			TunnelInterface:     "tunnel.1",
			AntiReplay:          true,
			CopyTos:             true,
			Type:                TypeManualKey,
			MkLocalSpi:          "local spi",
			MkRemoteSpi:         "remote spi",
			MkInterface:         "mk interface",
			MkLocalAddressIp:    "10.5.5.5",
			MkRemoteAddress:     "192.168.55.55",
			MkProtocol:          MkProtocolEsp,
			MkEspEncryptionType: MkEspEncryptionAes256,
			MkEspEncryptionKey:  "enc key",
			MkAuthType:          MkAuthTypeSha512,
			MkAuthKey:           "auth key",
		}},
		{"v3 manual key esp none", version.Number{8, 0, 0, ""}, Entry{
			Name:                "manual key esp none",
			TunnelInterface:     "tunnel.1",
			AntiReplay:          true,
			CopyTos:             true,
			Type:                TypeManualKey,
			MkLocalSpi:          "local spi",
			MkRemoteSpi:         "remote spi",
			MkInterface:         "mk interface",
			MkLocalAddressIp:    "10.5.5.5",
			MkRemoteAddress:     "192.168.55.55",
			MkProtocol:          MkProtocolEsp,
			MkEspEncryptionType: MkEspEncryption3des,
			MkEspEncryptionKey:  "enc key",
			MkAuthType:          MkAuthTypeNone,
		}},
		{"v3 manual key ah md5", version.Number{8, 0, 0, ""}, Entry{
			Name:             "manual key ah md5",
			TunnelInterface:  "tunnel.1",
			AntiReplay:       true,
			CopyTos:          true,
			Type:             TypeManualKey,
			MkLocalSpi:       "local spi",
			MkRemoteSpi:      "remote spi",
			MkInterface:      "mk interface",
			MkLocalAddressIp: "10.5.5.5",
			MkRemoteAddress:  "192.168.55.55",
			MkProtocol:       MkProtocolAh,
			MkAuthType:       MkAuthTypeMd5,
			MkAuthKey:        "auth key",
		}},
		{"v3 manual key ah sha1", version.Number{8, 0, 0, ""}, Entry{
			Name:             "manual key ah sha1",
			TunnelInterface:  "tunnel.1",
			AntiReplay:       true,
			CopyTos:          true,
			Type:             TypeManualKey,
			MkLocalSpi:       "local spi",
			MkRemoteSpi:      "remote spi",
			MkInterface:      "mk interface",
			MkLocalAddressIp: "10.5.5.5",
			MkRemoteAddress:  "192.168.55.55",
			MkProtocol:       MkProtocolAh,
			MkAuthType:       MkAuthTypeSha1,
			MkAuthKey:        "auth key",
		}},
		{"v3 manual key ah sha256", version.Number{8, 0, 0, ""}, Entry{
			Name:             "manual key ah sha256",
			TunnelInterface:  "tunnel.1",
			AntiReplay:       true,
			CopyTos:          true,
			Type:             TypeManualKey,
			MkLocalSpi:       "local spi",
			MkRemoteSpi:      "remote spi",
			MkInterface:      "mk interface",
			MkLocalAddressIp: "10.5.5.5",
			MkRemoteAddress:  "192.168.55.55",
			MkProtocol:       MkProtocolAh,
			MkAuthType:       MkAuthTypeSha256,
			MkAuthKey:        "auth key",
		}},
		{"v3 manual key ah sha384", version.Number{8, 0, 0, ""}, Entry{
			Name:             "manual key ah sha384",
			TunnelInterface:  "tunnel.1",
			AntiReplay:       true,
			CopyTos:          true,
			Type:             TypeManualKey,
			MkLocalSpi:       "local spi",
			MkRemoteSpi:      "remote spi",
			MkInterface:      "mk interface",
			MkLocalAddressIp: "10.5.5.5",
			MkRemoteAddress:  "192.168.55.55",
			MkProtocol:       MkProtocolAh,
			MkAuthType:       MkAuthTypeSha384,
			MkAuthKey:        "auth key",
		}},
		{"v3 manual key ah sha512", version.Number{8, 0, 0, ""}, Entry{
			Name:             "manual key ah sha512",
			TunnelInterface:  "tunnel.1",
			AntiReplay:       true,
			CopyTos:          true,
			Type:             TypeManualKey,
			MkLocalSpi:       "local spi",
			MkRemoteSpi:      "remote spi",
			MkInterface:      "mk interface",
			MkLocalAddressIp: "10.5.5.5",
			MkRemoteAddress:  "192.168.55.55",
			MkProtocol:       MkProtocolAh,
			MkAuthType:       MkAuthTypeSha512,
			MkAuthKey:        "auth key",
		}},
		{"v3 gps no cert", version.Number{8, 0, 0, ""}, Entry{
			Name:                      "gps no cert",
			TunnelInterface:           "tunnel.1",
			AntiReplay:                true,
			CopyTos:                   true,
			Type:                      TypeGlobalProtectSatellite,
			GpsPublishConnectedRoutes: true,
			GpsPublishRoutes:          []string{"route1", "route2"},
			GpsInterface:              "gps interface",
			GpsInterfaceIpIpv4:        "gps ipv4 ip",
			GpsInterfaceIpIpv6:        "gps ipv6 ip",
		}},
		{"v3 gps with cert", version.Number{8, 0, 0, ""}, Entry{
			Name:                       "gps with cert",
			TunnelInterface:            "tunnel.1",
			AntiReplay:                 true,
			CopyTos:                    true,
			Type:                       TypeGlobalProtectSatellite,
			GpsPublishConnectedRoutes:  false,
			GpsInterface:               "gps interface",
			GpsInterfaceFloatingIpIpv4: "gps ipv4 floating ip",
			GpsInterfaceFloatingIpIpv6: "gps ipv6 floating ip",
			GpsLocalCertificate:        "my cert",
			GpsCertificateProfile:      "cert prof",
		}},
	}
}
