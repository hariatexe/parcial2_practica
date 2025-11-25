func (dicc *hashCerrado[K,V]) CantidadValoresDistintos(cmp func(V,V) bool) int {
	hashAuxiliar := CrearHash[V, bool](cmp)

	for _, celda := range dicc.tabla {
		if celda.estado == OCUPADO {
			hashAuxiliar.Guardar(celda.valor, true)
			/* Verifico si el valor que quiero guardar ya se encuentra en el hashAuxiliar
				si la funcion devuelve true el guardar no hace nada
				si devuelve false guarda el elemento porque no esta
				Esto asegura que el hash auxiliar almacene exactamente los valores distintos
				Se llama a la funcion cmp(a,b) para comparar el valor nuevo (a) que estoy introduciendo
				con el valor viejo (b) ya guardado
				Si cmp(a,b) == true -> significa que ya existe un valor igual.
				Entonces el hash no inserta nada en esa posicion
				La operacion Guardar termina aqui (no hay cambios)
				Si cmp(a,b) == false -> significa que son distintos, entonces:
				Al ser distintos los valores quiere decir que en la posicion que le correspondia hay algun
				otro valor, por lo que busca la siguiente celda disponible segun la tecnica de sondeo
				E inserta el valor en la primera celda libre.
			*/
		}
	}
	return hashAuxiliar.Cantidad()
}

/*
	ANALISIS DE COMPLEJIDAD:
	La complejidad de este algoritmo es de O(n), al ser cerrado
	el peor caso es en el cual se tenga que recorrer n celdas ocupadas
	por lo tanto se guardan en el hash auxiliar los n elementos
	por lo tanto la complejidad del algoritmo es O(n)
	En el caso promedio recorreria n elementos o celdas
	siendo n la cantidad de celdas ocupadas del hash original, que-
	dando la complejidad en O(n)
*/


// Lo hacemos otra vez

func (dicc *hashCerrado[K,V]) CantidadValoresDistintos(cmp func(V,V) bool) int {
		hashAuxiliar := CrearHash[V, bool](cmp)
		for _, celda := range dicc.tabla {
			if celda.estado = OCUPADO {
				hashAuxiliar.Guardar(celda.valor, true)
			}
		}
		return hashAuxiliar.Cantidad()
}


func (dicc *hashCerrado[K,V]) CantidadValoresDistintos(cmp func(V,V) bool) int {
	// cuantos valores son distintos
	hashAuxiliar := CrearHash[V, bool](cmp)
	for _, celda := range dicc.tabla {
		if celda.estado == OCUPADO {
			hashAuxiliar.Guardar(celda.valor, true)
		}
	}
	return hashAuxiliar.Cantidad()
}