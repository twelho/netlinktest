# netlinktest

A small Go binary to test seemingly arbitrary TAP adapter MAC address changes in the kernel during bridge attach. Implemented with [netlink](https://github.com/vishvananda/netlink), see [vishvananda/netlink#553](https://github.com/vishvananda/netlink/issues/553) for more details.

## Usage

`make` to build the binary, then run it with `sudo ./bin/netlinktest`. Use `sudo make clean-ifaces` to clean up the generated interfaces post-run.

## License

[MIT](https://opensource.org/license/mit/) ([LICENSE](LICENSE))
