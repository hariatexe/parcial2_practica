func (abb *Arbol[int]) HijosDirectos() {
	if abb == nil {
		return 0
	}

	abbIzq := abb.izq.HijosDirectos()
	abbDer := abb.der.HijosDirectos()

	tieneDos := 0
	if abb.izq != nil && abb.der != nil {
		tieneDos = 1
	}
	return tieneDos + abbIzq + abbDer
}

// La complejidad de este algoritmo es O(n)