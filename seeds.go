package dvtutil

import "github.com/martinboehm/btcutil/chaincfg"

var MainnetDNSSeeds = []chaincfg.DNSSeed{
	{"seed.exploredvt.com", true},
	{"seed.minedvt.com", true},
	{"seed.devault.online", true},
	{"seed.proteanx.com", true},
	{"seed.devault.cc", true},
}

var TestnetDNSSeeds = []chaincfg.DNSSeed{
	{"testnet-seed.exploredvt.com", true},
}

func GetDNSSeed(params *chaincfg.Params) []chaincfg.DNSSeed {
	if params.Name == chaincfg.MainNetParams.Name {
		return MainnetDNSSeeds
	}
	return TestnetDNSSeeds
}
