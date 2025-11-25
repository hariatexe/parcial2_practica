	func ExisteUnParConSuma(arr []int, K int) {
	hashAuxiliar := CrearHash[int, bool]()
	for _, numero := range arr {
		if hashAuxiliar.Pertenece(K - numero) {
			return true
			// encontramos que el complemento -> complemento + numero = K
			// por lo que podria decir que existe el par complemento + numero = K
		}
		hashAuxiliar.Guardar(numero, true)
	}
	return false
}

/*
	La complejidad del algoritmo es la siguiente:
	Itero por cada elemento del arreglo por lo tanto podriamos llamar n
	a la cantidad de elementos en el arreglo, por cada elemento hago una operacion
	O(1) en promedio ya que verifico si el complemento pertenece, sino lo guardo
	por lo tanto seria n veces algo O(1) promedio, la complejidad resultante del algoritmo
	es O(n)
*/
