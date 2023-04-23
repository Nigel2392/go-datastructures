package hashmap_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/Nigel2392/go-datastructures"
	"github.com/Nigel2392/go-datastructures/hashmap"
)

type stringHasher string // implements Hasher

func (s stringHasher) Hash() uint64 {
	// return the key as a uint64
	return datastructures.FastStrHash(s)
}

func (s stringHasher) Equals(other stringHasher) bool {
	return s == other //datastructures.FastStrCmp(s, other)
}

func TestHashMap(t *testing.T) {
	var (
		stringHasherKeys = [...]stringHasher{
			"a", "b", "c", "d", "e", "f", "g", "ab", "bc",
			"cd", "de", "ef", "fg", "sun", "sunny", "flower", "flow", "flight", "flights", "mango",
		}
	)

	var hashTable = hashmap.Map[stringHasher, string]()
	for _, v := range stringHasherKeys {
		hashTable.Set(v, string(v))
	}

	for _, v := range stringHasherKeys {
		if val, ok := hashTable.Get(v); !ok || val != string(v) {
			t.Fatalf("key: %s, value: %s", v, val)
		}
	}

	t.Logf("%#v", hashTable)

	for _, v := range stringHasherKeys {
		if !hashTable.Delete(v) {
			t.Fatalf("couldn't delete key: %s, value: %s\n%#v", v, v, hashTable)
		}
	}

	if hashTable.Len() != 0 {
		t.Fatalf("Size: %d", hashTable.Len())
	}

	for _, v := range stringHasherKeys {
		if _, ok := hashTable.Get(v); ok {
			t.Fatalf("key: %s, value: %s", v, v)
		}
	}

	hashTable.Range(func(key stringHasher, value string) bool {
		t.Fatalf("key: %s, value: %s", key, value)
		return false
	})

	t.Logf("%#v", hashTable)

	for _, v := range stringHasherKeys {
		hashTable.Set(v, string(v))
	}

	if hashTable.Len() != len(stringHasherKeys) {
		t.Fatalf("Size: %d", hashTable.Len())
	}

	for _, v := range stringHasherKeys {
		var val, found = hashTable.Pop(v)
		if !found || val != string(v) {
			t.Fatalf("key: %s, value: %s", v, val)
		}
	}

	if hashTable.Len() != 0 {
		t.Fatalf("Size: %d", hashTable.Len())
	}

	t.Logf("%#v", hashTable)
}

var (
	SmallArrayKeys = [256]stringHasher{}

	MediumArrayKeys = [256 * 256]stringHasher{}

	LargeArrayKeys = [256 * 256 * 256]stringHasher{}
)

func initArrays() {
	var currentTimeUnix = time.Now().UnixNano()
	for i := 0; i < len(LargeArrayKeys); i++ {
		var key = stringHasher("key" + strconv.FormatInt(int64(i)+currentTimeUnix, 10))
		if i < len(SmallArrayKeys) {
			SmallArrayKeys[i] = key
		}
		if i < len(MediumArrayKeys) {
			MediumArrayKeys[i] = key
		}
		LargeArrayKeys[i] = key
	}
}

func BenchmarkHashMap_SetSmall(b *testing.B) {
	b.StopTimer()
	initArrays()
	var hashTable = hashmap.Map[stringHasher, string](len(SmallArrayKeys))
	for _, v := range SmallArrayKeys {
		hashTable.Set(v, string(v))
	}
	b.Logf("Starting benchmark with %d items", len(SmallArrayKeys))
	var lastKey stringHasher = SmallArrayKeys[len(SmallArrayKeys)-1]
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if val, ok := hashTable.Get(lastKey); !ok || val != string(lastKey) {
			b.Fatalf("key: %s, value: %s", lastKey, val)
		}
	}
}
func BenchmarkHashMap_SetMedium(b *testing.B) {
	b.StopTimer()
	initArrays()
	var hashTable = hashmap.Map[stringHasher, string](len(MediumArrayKeys))
	for _, v := range MediumArrayKeys {
		hashTable.Set(v, string(v))
	}
	b.Logf("Starting benchmark with %d items", len(MediumArrayKeys))
	var lastKey stringHasher = MediumArrayKeys[len(MediumArrayKeys)-1]
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if val, ok := hashTable.Get(lastKey); !ok || val != string(lastKey) {
			b.Fatalf("key: %s, value: %s", lastKey, val)
		}
	}
}

func BenchmarkHashMap_SetLarge(b *testing.B) {
	b.StopTimer()
	initArrays()
	var hashTable = hashmap.Map[stringHasher, string](len(LargeArrayKeys))
	for _, v := range LargeArrayKeys {
		hashTable.Set(v, string(v))
	}
	b.Logf("Starting benchmark with %d items", len(LargeArrayKeys))
	var lastKey stringHasher = LargeArrayKeys[len(LargeArrayKeys)-1]
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if val, ok := hashTable.Get(lastKey); !ok || val != string(lastKey) {
			b.Fatalf("key: %s, value: %s", lastKey, val)
		}
	}
}

func BenchmarkSTDMap_Large(b *testing.B) {
	b.StopTimer()
	var hashTable = make(map[stringHasher]string, len(LargeArrayKeys))
	for _, v := range LargeArrayKeys {
		hashTable[v] = string(v)
	}
	b.Logf("Starting benchmark with %d items", len(LargeArrayKeys))
	var lastKey stringHasher = LargeArrayKeys[len(LargeArrayKeys)-1]
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if val, ok := hashTable[lastKey]; !ok || val != string(lastKey) {
			b.Fatalf("key: %s, value: %s", lastKey, val)
		}
	}
}
