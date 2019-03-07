package address

import (
	"math/rand"
	"testing"
	"time"

	"gx/ipfs/QmPVkJMTeRC6iBByPWdrRkD3BE5UXsj5HPzb4kPqL186mS/testify/require"
	logging "gx/ipfs/QmbkT7eMTyXfpeyB3ZMxxcxg7XH8t6uXp49jqzz4HB7BGF/go-log"

	"github.com/filecoin-project/go-filecoin/bls-signatures"
	"github.com/filecoin-project/go-filecoin/crypto"
)

func init() {
	logging.SetDebugLogging()
	rand.Seed(time.Now().Unix())
}

func TestIDAddress(t *testing.T) {
	require := require.New(t)

	addr := NewIDAddress(uint64(rand.Int()))
	require.Equal(ID, addr.Protocol())

	str, err := encode(Mainnet, addr)
	require.NoError(err)

	maybe, err := decode(str)
	require.NoError(err)
	require.Equal(addr, maybe)

}

func TestSecp256k1Address(t *testing.T) {
	require := require.New(t)

	sk, err := crypto.GenerateKey()
	require.NoError(err)

	addr := NewSecp256k1Address(crypto.PublicKey(sk))
	require.Equal(SECP256K1, addr.Protocol())

	str, err := encode(Mainnet, addr)
	require.NoError(err)

	maybe, err := decode(str)
	require.NoError(err)
	require.Equal(addr, maybe)

}

func TestActorAddress(t *testing.T) {
	require := require.New(t)

	actorMsg := make([]byte, 20)
	rand.Read(actorMsg)

	addr := NewActorAddress(actorMsg)
	require.Equal(Actor, addr.Protocol())

	str, err := encode(Mainnet, addr)
	require.NoError(err)

	maybe, err := decode(str)
	require.NoError(err)
	require.Equal(addr, maybe)

}

func TestBLSAddress(t *testing.T) {
	require := require.New(t)

	pk := bls.PrivateKeyPublicKey(bls.PrivateKeyGenerate())

	addr := NewBLSAddress(pk[:])
	require.Equal(BLS, addr.Protocol())

	str, err := encode(Mainnet, addr)
	require.NoError(err)

	maybe, err := decode(str)
	require.NoError(err)
	require.Equal(addr, maybe)

}

func TestMarshalJSON(t *testing.T) {
	require := require.New(t)

	actorMsg := make([]byte, 20)
	rand.Read(actorMsg)

	addr := NewActorAddress(actorMsg)
	require.Equal(Actor, addr.Protocol())

	b, err := addr.MarshalJSON()
	require.NoError(err)

	var newAddr Address
	newAddr.UnmarshalJSON(b)
	require.Equal(addr, newAddr)

}
