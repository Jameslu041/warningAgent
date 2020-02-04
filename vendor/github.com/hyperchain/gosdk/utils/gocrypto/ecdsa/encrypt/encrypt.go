package encrypt

import (
	"github.com/hyperchain/gosdk/common"
)

func PrivateToAddress(privateHex string) (string, error) {
	if privateHex == "" {
		return "", nil
	}
	p := ToECDSA(common.FromHex(privateHex))
	addr := PubkeyToAddress(p.PublicKey)
	return common.ToHex(addr[:]), nil
}
