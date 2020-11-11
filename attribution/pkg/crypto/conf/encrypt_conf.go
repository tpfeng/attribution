package conf

import (
	"flag"
)

var prime = flag.String("prime", "CC5A0467495EC2506D699C9FFCE6ED5885AE09384671EE00B93A28712DA814240F2A471B2C77B120FE70DA02D33B0F85CD737B800942D5A2BF80DCCD290FB6553C9834197DC498F6AC69B5ECEF3FD8B05F231E3632E9AECA2B3F50977CF9033AF3005A9C0A339CFB4922971B3AF05A5955983C12B153BB78A2B1FB14C84A3C662ADDE5BCEBE8779FF9F97C6E73BD29D4044242581455EEB0543E2DB35F43997F46F8596A58080DC053BBB71F9A557185DF80738238713D3EDFD77D47B26977B373FB0D969920B3909CCC24792B5B4E94AD29F6AE6BD73ED5FE6528CDDBEA1560BBCD36E8B25008021A26E9A4E51BBCD8436F38D6A222E2138E7042A73A7877D7", "")

type EncryptConf struct {
	Prime string
}

func (c *EncryptConf) InitConf() *EncryptConf {
	c.Prime = *prime
	return c
}

// 这个方法支持用户直接通过参数传入一个
func NewEncryptConf(prime string) *EncryptConf {
	return &EncryptConf{
		Prime: prime,
	}
}