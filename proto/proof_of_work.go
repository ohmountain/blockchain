package proto

import (
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"
	"runtime"
	"strconv"
)

var (
	maxNonce = math.MaxInt64
	CPUS     = runtime.NumCPU() * 2
	SUCCESS  = false
)

// 难度
const targetBits = 16

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// 对于一个区块来说，一旦生成hash，那么这个数据是固定的
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	return bytes.Join([][]byte{
		pow.block.PrevBlockHash,
		pow.block.Data,
		[]byte(strconv.FormatInt(pow.block.Timestamp, 16)),
		[]byte(strconv.FormatInt(targetBits, 16)),
		[]byte(strconv.FormatInt(int64(nonce), 16)),
	}, []byte{})
}

// 工作量证明
// 即某一块产生的hash一定小于target
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	var hash [32]byte

	data := pow.prepareData(pow.block.Nonce)
	hash = sha256.Sum256(data[:])
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.target) == -1
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hash [32]byte
	nonce := 0

	ch := make(chan int64, CPUS)

	single := maxNonce / CPUS

	for i := 1; i <= CPUS; i++ {
		start := (i-1)*single + 1
		end := single * i

		if i == CPUS {
			end += maxNonce - single*i
		}

		go calcHash(pow, ch, int64(start), int64(end))
	}

	for i := 0; i < CPUS; i++ {
		n := <-ch

		if n > 0 {
			SUCCESS = false
			nonce = int(n)
		}
	}

	hash = sha256.Sum256(pow.prepareData(nonce))

	return nonce, hash[:]
}

func calcHash(pow *ProofOfWork, ch chan int64, start int64, end int64) {
	var hashInt big.Int
	var hash [32]byte
	var nonce int = int(start)
	var final int = int(end)

	for nonce <= final {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			SUCCESS = true
			ch <- int64(nonce)
			break
		} else {
			nonce++
		}

		if SUCCESS && pow.block.Validate() {
			ch <- 0
			break
		}

	}

	ch <- 0
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
	b.PoW = pow

	return pow
}
