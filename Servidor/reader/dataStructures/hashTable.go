package dataStructures

import "fmt"

type Comentario struct {
	Cadena     string
	Persona    int
	Respuestas HashTable
}

type HashTable struct {
	vector       []*NodeHash
	elements     int
	chargeFactor float64
	size         int
}

type NodeHash struct {
	index  int
	tuples []Tupla
}

type Tupla struct {
	Key   string
	Value string
}

func Tupla_(key string, value string) *Tupla {
	return &Tupla{Key: key, Value: value}
}

func NodoHash_(index int) *NodeHash {
	return &NodeHash{index: index}
}

func HashTable_(size int) *HashTable {
	hashTable := &HashTable{elements: 0, chargeFactor: 0, size: size}
	for i := 0; i < size; i++ {
		hashTable.vector = append(hashTable.vector, nil)
	}
	return hashTable
}

func (hash *HashTable) Attributes() {
	fmt.Println(hash.size, hash.elements, hash.chargeFactor)
}

func (hash *HashTable) HashFunction(id int) int {
	posicion := id % (hash.size - 1)
	if posicion > hash.size {
		return posicion - hash.size
	}
	return posicion
}

func (hash *HashTable) rehashing() {
	follow := hash.size
	factor := 0.0

	for factor < 0.3 {
		factor = float64(hash.elements) / float64(follow)
		follow++
	}
	hash_temporal := HashTable_(follow)

	for _, nodo := range hash.vector {
		for _, tupla := range nodo.tuples {
			hash_temporal.Insert(int(stringtoascii(tupla.Key)), tupla.Key, tupla.Value)
		}
	}

	hash.vector = hash_temporal.vector
	hash.elements = hash_temporal.elements
	hash.size = hash_temporal.size
	hash.chargeFactor = hash_temporal.chargeFactor
}

func (hash *HashTable) Insert(id int, key string, value string) {
	position := hash.HashFunction(id)
	if hash.vector[position] != nil {
		nuevo := Tupla_(key, value)
		hash.vector[position].tuples = append(hash.vector[position].tuples, *nuevo)
	} else {
		nuevo := NodoHash_(position)
		nuevo.tuples = append(nuevo.tuples, *Tupla_(key, value))
		hash.vector[position] = nuevo
		hash.elements++
		hash.chargeFactor = float64(hash.elements) / float64(hash.size)
	}
	if hash.chargeFactor > 0.6 {
		hash.rehashing()
	}
}

func (node *NodeHash) Show() {
	for i := 0; i < len(node.tuples); i++ {
		fmt.Println("index ", i, "value ", node.tuples[i].Value)
	}
}

func (hash *HashTable) Show() {
	for i := 0; i < hash.size; i++ {
		if hash.vector[i] == nil {
			fmt.Println("position ", i, "vacia")
		} else {
			fmt.Println("position ", i)
			hash.vector[i].Show()
		}
	}
}

func stringtoascii(entrada string) int32 {
	cod := []rune(entrada)
	var temp int32
	temp = 0
	for i := 0; i < len(cod); i++ {
		temp = cod[i] + temp
	}
	return temp
}
