package keys

import (
	"io"

	"github.com/FiiLabs/metaos/wallet/keyring"
)

type KeybaseGenerator func(buf io.Reader) (keyring.Keystore, error)
