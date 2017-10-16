package bridge

import "github.com/docker/libnetwork/iptables"

func (n *bridgeNetwork) setupFirewalld(config *networkConfiguration, i *bridgeInterface) error {
	d := n.driver
	d.Lock()
	driverConfig := d.config
	d.Unlock()

	// Sanity check.
	if driverConfig.EnableIPTables == false {
		return IPTableCfgError(config.BridgeName)
	}

	iptables.OnReloaded(func() { n.setupIPTables(config, i) })
	iptables.OnReloaded(n.portMapper.ReMapAll)

	if driverConfig.EnableIPv6 == true {
		iptables.OnReloaded(n.portMapperV6.ReMapAll)
	}

	return nil
}
