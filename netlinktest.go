package main

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"github.com/vishvananda/netlink"
	"net"
)

func main() {
	tp0, err := createTAPAdapter("nltest_tp0")
	assert(err)

	tp1, err := createTAPAdapter("nltest_tp1")
	assert(err)

	br0, err := createBridge("nltest_br0")
	assert(err)

	log.Warn("Attach first adapter to bridge")
	attachToBridge(tp0, br0)

	log.Warn("Attach second adapter to bridge")
	attachToBridge(tp1, br0)
}

func attachToBridge(link netlink.Link, bridge *netlink.Bridge) {
	la := getMAC("Pre-LinkSetMaster", link)
	ba := getMAC("Pre-LinkSetMaster", bridge)
	checkDuplicate(la, ba)

	log.Infof("LinkSetMaster: %q -> %q", link.Attrs().Name, bridge.Name)
	assert(netlink.LinkSetMaster(link, bridge))

	lb := getMAC("Post-LinkSetMaster", link)
	bb := getMAC("Post-LinkSetMaster", bridge)

	// The bridge is assumed to change its MAC, see:
	// - https://backreference.org/2010/07/28/linux-bridge-mac-addresses-and-dynamic-ports/
	// - https://lists.linuxfoundation.org/pipermail/bridge/2010-May/007204.html

	// Check if the MAC of the given link has changed
	checkChanged(link.Attrs().Name, la, lb)

	checkDuplicate(lb, bb)
}

func createTAPAdapter(tapName string) (*netlink.Tuntap, error) {
	la := netlink.NewLinkAttrs()
	la.Name = tapName
	tuntap := &netlink.Tuntap{
		LinkAttrs: la,
		Mode:      netlink.TUNTAP_MODE_TAP,
	}
	return tuntap, addLink(tuntap)
}

func createBridge(bridgeName string) (*netlink.Bridge, error) {
	la := netlink.NewLinkAttrs()
	la.Name = bridgeName
	bridge := &netlink.Bridge{LinkAttrs: la}
	return bridge, addLink(bridge)
}

func addLink(link netlink.Link) (err error) {
	if err = netlink.LinkAdd(link); err == nil {
		err = netlink.LinkSetUp(link)
	}

	return
}

// Utility functions

func getMAC(context string, l netlink.Link) net.HardwareAddr {
	// Fetch the *net.Interface for the given Link,
	// this is done to retrieve the populated HardwareAddr
	i, err := net.InterfaceByIndex(l.Attrs().Index)
	assert(err)

	log.Infof("%s %q MAC address: %s", context, i.Name, i.HardwareAddr)
	return i.HardwareAddr
}

func checkChanged(name string, a, b net.HardwareAddr) {
	if !bytes.Equal(a, b) {
		log.Errorf("%q changed MAC address: %s -> %s", name, a, b)
	}
}

func checkDuplicate(a, b net.HardwareAddr) {
	if bytes.Equal(a, b) {
		log.Errorf("Duplicate MAC address: %s", a)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
