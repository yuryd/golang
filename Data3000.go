// Data3000
package main

import (
	"fmt"
	"hash/fnv"
	"math"

	"github.com/govalues/decimal"
)

const maxPrimes = 25000

var primes []uint64 = append(make([]uint64, 0, maxPrimes), []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}...)

func init() {
	fmt.Print("Locating primes.. ")
	//	primes = make([]uint64, maxPrimes)
	const maxPrimes uint64 = 25000
	var foundPrimes uint64
	for i := 30; uint64(i) <= maxPrimes; i++ {
		prime := true
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				prime = false
				foundPrimes++
				break
			}
		}
		if prime {
			primes = append(primes, uint64(i))
		}
	}
	fmt.Println(foundPrimes, "Completed")
	for _, p := range primes {
		if p == 0 {
			continue
		}
		fmt.Print(p, " ")
	}
	fmt.Println()
}

func NextPrime(value uint64) uint64 {
	var result uint64
	for _, p := range primes {
		if p >= value {
			result = p
			break
		}
	}
	return result
}

func hashToUint64(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}

type MasterTable[T MasterTableAccessor] struct {
	data     []*T
	freemap  []uint64
	name     string
	capacity uint64
}

func (mt *MasterTable[T]) GetRecId(keyVal any) uint64 {
	if str, ok := keyVal.(string); ok {
		/* act on str */
		return hashToUint64(str) % mt.capacity
	} else {
		/* not string */
		return 0
	}

}

func (mt *MasterTable[T]) Used() uint64 {
	var result uint64
	for _, byte := range mt.freemap {
		for i := range 64 {
			result += byte & (1 << (63 - i))
		}
	}
	return result
}

func (mt *MasterTable[T]) Room() uint64 {
	return mt.capacity - mt.Used()
}

func MakeFreeMap(size uint64) []uint64 {
	var bytes uint64
	if size%64 == 0 {
		bytes = size / 64
	} else {
		bytes = size/64 + 1
	}
	return make([]uint64, bytes, bytes)
}

func (mt *MasterTable[T]) IsPosFree(pos uint64) bool {
	_, _, _, free := mt.GetFreeMap(pos)
	return free
}

func (mt *MasterTable[T]) freeMap(pos uint64) (uint64, uint8, uint64) {
	byte := pos / 64
	bit := pos % 64
	mask := 1 << (63 - bit)
	return uint64(byte), uint8(bit), uint64(mask)
}

func (mt *MasterTable[T]) GetFreeMap(pos uint64) (uint64, uint8, uint64, bool) {
	byte, bit, mask := mt.freeMap(pos)
	return byte, bit, mask, mt.freemap[byte]&mask == 0
}

func (mt *MasterTable[T]) SetFreeMap(pos uint64, value bool) {
	byte, _, mask := mt.freeMap(pos)
	if value {
		mt.freemap[byte] |= mask
	} else {
		mt.freemap[byte] ^= mask
	}
}

func (mt *MasterTable[T]) Put(record MasterTableAccessor) bool {
	var result bool
	recPos := mt.GetRecId(record.KeyValue())
	_, _, _, free := mt.GetFreeMap(recPos)
	if free {
		if rec, ok := record.(MasterTableAccessor); ok {
			mt.data[recPos] = &rec
			mt.SetFreeMap(recPos, true)
			result = true
		}
	} else {
	}
	return result
}

func MakeMasterTable[T MasterTableAccessor](name string, size uint64) *MasterTable[T] {
	cap := NextPrime(size)
	res := MasterTable[T]{data: make([]T, cap, cap), name: name, capacity: cap, freemap: MakeFreeMap(cap)}
	return &res
}

type MasterTableAccessor interface {
	KeyValue() any
}

type ItemMaster struct {
	itemId    string
	itemName  string
	itemCost  decimal.Decimal
	itemPrice decimal.Decimal
}

func (mta *ItemMaster) KeyValue() any {
	return mta.itemId
}

func (mta *ItemMaster) ItemName() string {
	return mta.itemName
}

func (mta *ItemMaster) ItemCost() decimal.Decimal {
	return mta.itemCost
}

func (mta *ItemMaster) ItemPrice() decimal.Decimal {
	return mta.itemPrice
}

func main() {
	const capacity = 25
	fmt.Println("data3000!")
	strs := []string{"hello world", "goodbye forever", "goodbye forever!"}
	for _, s := range strs {
		hash := hashToUint64(s)
		fmt.Println(s, hash, hash%capacity)
	}
	mt := MakeMasterTable[*MasterTableAccessor]("Items", capacity)
	fmt.Printf("%+v\n", mt)
	byte, bit, mask, free := mt.GetFreeMap(64)
	fmt.Printf("%d %d %064b %t\n", byte, bit, mask, free)
	mt.SetFreeMap(64, true)
	byte, bit, mask, free = mt.GetFreeMap(64)
	fmt.Printf("%d %d %064b %t\n", byte, bit, mask, free)
	mt.SetFreeMap(64, false)
	byte, bit, mask, free = mt.GetFreeMap(64)
	fmt.Printf("%d %d %064b %t\n", byte, bit, mask, free)
}
