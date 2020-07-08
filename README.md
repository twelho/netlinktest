# netlinktest

A small Go binary to test seemingly arbitrary MAC address changes in [netlink](https://github.com/vishvananda/netlink).
See [vishvananda/netlink#553](https://github.com/vishvananda/netlink/issues/553) for more details.

## Usage

`make` to build the binary, then run it with `sudo ./bin/netlinktest`.
Use `sudo make clean-ifaces` to clean up the generated interfaces post-run.