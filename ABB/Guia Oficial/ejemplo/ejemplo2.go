
func (abb *abb[string]) Mayores(clave string, mayores Lista[string]) {
	if abb == nil {
		return
	}

	if abb.clave > clave {
		abb.izq.Mayores(clave, mayores)
		mayores.InsertarUltimo(abb.clave)
	}
	abb.der.Mayores(clave, mayores)
}

/*
	Complejidad del algoritmo:

	En el peor de los casos la clave es menor a la mínima del árbol.
	Como la llamada a abb.der.Mayores(clave, mayores) no está dentro de un if,
	siempre se realiza, por lo que debo recorrer tanto el subárbol izquierdo
	como el derecho. En consecuencia, en este caso visito todos los nodos.
	La complejidad en el peor caso es O(n).

	En el mejor caso la clave es la mayor del árbol.
	La condición (abb.clave > clave) nunca se cumple, por lo que no se
	procesa el subárbol izquierdo. Sin embargo, siempre se desciende por el
	subárbol derecho, cuya altura es log n para un ABB balanceado.
	Por lo tanto, la complejidad en el mejor caso es Ω(log n).

*/





func (abb *abb[string]) Mayores(clave string, mayores Lista[string]) {
	// caso BASE SIEMPRE
	if abb == nil {
		return
	}

	if strings.strcmp(abb.clave, clave) > 0 {
		abb.izq.Mayores(clave, mayores)
		mayores.InsertarUltimo(abb.clave)
	}

	abb.der.Mayores(clave, mayores)
}

