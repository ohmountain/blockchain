package proto

import (
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"
	"strconv"
)

var (
	maxNonce = math.MaxInt64
)

// 难度
const targetBits = 24

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// 对于一个区块来说，一旦生成hash，那么这个数据是固定的
func (pof *ProofOfWork) prepareData(nonce int) []byte {
	return bytes.Join([][]byte{
		pof.block.PrevBlockHash,
		pof.block.Data,
		[]byte(strconv.FormatInt(pof.block.Timestamp, 16)),
		[]byte(strconv.FormatInt(targetBits, 16)),
		[]byte(strconv.FormatInt(int64(nonce), 16)),
	}, []byte{})
}

// 工作量证明
// 即某一块产生的hash一定小于target
func (pof *ProofOfWork) Validate() bool {
	var hashInt big.Int
	var hash [32]byte

	data := pof.prepareData(pof.block.Nonce)
	hash = sha256.Sum256(data[:])
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pof.target) == -1
}

func (pof *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	for nonce < maxNonce {
		data := pof.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pof.target) == -1 {
			break
		} else {
			nonce++
		}

	}

	return nonce, hash[:]
}

func NewProofOfWork(b *Block) *ProofOfWork {

	target := big.NewInt(1)

	// 左移，左移位数越小，目标越大，越容易产生符合规则的数据，挖矿难度越低
	// 如果左移256位，则target = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF，那么随意产生一个数据都能够小于target
	// 如果左移0位，那么target = 0x1，产生的数据不可能小于target
	target.Lsh(target, uint(256-targetBits))

	pow := new(ProofOfWork)
	pow.target = target
	pow.block = b
	b.PoF = pow

	return pow
}
