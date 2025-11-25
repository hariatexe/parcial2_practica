func (abb *abb[int]) suma() int {
	if abb == nil {
		return 0
	}

	abbIzq := abb.izq.suma() // es un int, no una estructura con campo .clave
	abbDer := abb.der.suma() // tambien es un int

	// si no pusiera abb.clave estaria perdiendo el valor del nodo actual
	return abb.clave + abbIzq + abbDer
}

/*
	Complejidad del algoritmo:
	Este algoritmo tiene una complejidad de O(n), siendo n la cantidad de nodos
	visitados, ya que para esta primitiva recorre todos los nodos sumandos
	sus respectivos valores (claves)
	
*/

