// La implementacion ya sea como funcion o como primitiva
// vale tanto para el hash cerrado y como funcion

func (hash hashAbierto[K,V]) Claves() Lista[K] {
	claves := CrearListaEnlazada[K]()
	// Iterar() recibe un funcion que aplicara sobre los elementos
	hash.Iterar(func (clave K, valor V) bool {
		claves.InsertarUltimo(clave)
		return true
	})
	return claves
}
