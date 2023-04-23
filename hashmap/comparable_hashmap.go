package hashmap

type Comparable[T1 comparable, T2 any] struct {
	m        *HashMap[*MapKey[T1], T2]
	hashFunc func(T1) uint64
}

func NewComparable[T1 comparable, T2 any](buckets ...int) *Comparable[T1, T2] {
	var m = &Comparable[T1, T2]{
		m: Map[*MapKey[T1], T2](buckets...),
	}
	m.hashFunc = getHashFunc[T1]()
	return m
}

func (m *Comparable[T1, T2]) Set(k T1, v T2) {
	(*HashMap[*MapKey[T1], T2])(m.m).Set(makeHasher(k, m.hashFunc), v)
}

func (m *Comparable[T1, T2]) Get(k T1) (v T2, ok bool) {
	return (*HashMap[*MapKey[T1], T2])(m.m).Get(makeHasher(k, m.hashFunc))
}

func (m *Comparable[T1, T2]) Delete(k T1) (ok bool) {
	return (*HashMap[*MapKey[T1], T2])(m.m).Delete(makeHasher(k, m.hashFunc))
}

func (m *Comparable[T1, T2]) Len() int {
	return (*HashMap[*MapKey[T1], T2])(m.m).Len()
}

func (m *Comparable[T1, T2]) Keys() []T1 {
	var keys = make([]T1, m.Len(), m.Len())
	var i int
	(*HashMap[*MapKey[T1], T2])(m.m).Range(func(k *MapKey[T1], v T2) bool {
		keys[i] = k.v
		i++
		return true
	})
	return keys
}

func (m *Comparable[T1, T2]) Values() []T2 {
	var values = make([]T2, m.Len(), m.Len())
	var i int
	(*HashMap[*MapKey[T1], T2])(m.m).Range(func(k *MapKey[T1], v T2) bool {
		values[i] = v
		i++
		return true
	})
	return values
}

func (m *Comparable[T1, T2]) Range(f func(T1, T2) bool) {
	(*HashMap[*MapKey[T1], T2])(m.m).Range(func(k *MapKey[T1], v T2) bool {
		return f(k.v, v)
	})
}

func (m *Comparable[T1, T2]) Clear() {
	(*HashMap[*MapKey[T1], T2])(m.m).Clear()
}
