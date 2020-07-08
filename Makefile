build:
	CGO_ENABLED=0 go build -o bin/netlinktest

clean-ifaces:
	for i in nltest_tp0 nltest_tp1 nltest_br0; do ip l del $$i; done