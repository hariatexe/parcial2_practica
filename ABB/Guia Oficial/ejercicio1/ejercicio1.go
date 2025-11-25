func (abb *abb[string]) Altura() int{
	if abb == nil {
		return 0
	}
	abbIzq := abb.izq.Altura()
	abbDer := abb.der.Altura()
	return 1 + max(abbIzq, abbDer)
} 


/*
	La complejidad de este algoritmo es de O(n)
	Recorre todos los nodos del arbol solamente una vez, si tengo n nodos
	entonces tengo complejidad de O(n)
*/