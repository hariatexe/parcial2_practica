// LA IMPLEMENTACION DEL HASH CERRADO SERIA ASI

var estadoParClaveDato int
// ESTABLEZCO UN ENUMERABLE
const (
	VACIO = estadoParClaveDato(iota)
	BORRADO
	OCUPADO
)

type celdaHash[K comparable, V any] struct {
	clave K
	dato V
	estado estadoParClaveDato
}

type hashCerrado[K comparable, V any] struct {
	tabla []celdaHash[K,V]
	cantidad int
	tam int
	borrados int
}

func (hash hashCerrado[K,V]) Claves() Lista[K] {
	claves := CrearListaEnlazada[K]()
	for _, celda := range hash.tabla {
		if celda.estado == OCUPADO {
			claves.InsertarUltimo(celda.clave)
		}
	}
	return claves
}


// LA IMPLEMENTACION DEL HASH ABIERTO SERIA ASI:

type celda[K comparable, V any] struct {
	clave K
	dato V
}

type hashAbierto[K comparable, V any] struct {
	tabla []lista.Lista[celda[K,V]]
	cantidad int
	cmp func(K,K) bool
}

// ITERADOR INTERNO DEL HASH ABIERTO - SU IMPLEMENTACION

type iterHashAbierto[K comparable, V any] struct {
	hashAbierto *hashAbierto[K,V] // el hash sobre el que itera, modificando el hash
	indiceBalde int
	iteradorLista lista.IteradorLista[celda[K,V]] // iterador que itera sobre
	// la lista del balde
}

// UTILIZANDO EL ITERADOR DE LISTA
func (hash hashAbierto[K,V]) Claves() Lista[K] {
	claves := CrearListaEnlazada[K]()
	for _, lista := hash.tabla {
		lista.IteradorLista(func (par parClaveDato[K,V]) bool {
			claves.InsertarUltimo(par.clave)
			return true
		})
	}
	return claves
}

// UTILIZANDO EL ITERADOR INTERNO DEL HASH
func (hash hashAbierto[K,V]) Claves() Lista[K] {
	claves := CrearListaEnlazada[K]()
	// Utilizando el iterador interno
	hash.Iterar(func (clave K, dato V) bool {
		claves.InsertarUltimo(clave)
		return true
	})
	return claves
}



