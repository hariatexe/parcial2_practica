// Dos hashes son iguales si tienen la misma cantidad de elementos
// Si tienen las mismas claves con los mismos valores


// Se pueden comparar los valores por lo tanto tiene que ser V comparable
func compararDosHashes[K comparable, V comparable](hash1, hash2 Diccionario[K,V]) bool {
	if hash1.Cantidad() != hash2.Cantidad() {
		return false // diferentes cantidades
	}
	// Al asegurarme que la cantidad es igual tanto en el hash1 como en el hash2
	// me ahorro el hecho de ver si las claves de hash2 estan en hash1
	// ya que si una clave esta en hash1 y no en hash2, las cantidades serian diferentes

	encontrado := true

	hash1.Iterar(func(clave K, valor V) bool {
		if !hash2.Pertenece(clave) || hash2.Obtener(clave) != valor {
			encontrado = false
			return false
		}
		return true // seguir iterando 
	})

	return encontrado
}


func compararDosHashes[K comparable, V comparable](hash1, hash2 Diccionario[K,V]) bool {
	if hash1.Cantidad() != hash2.Cantidad() {
		return false
	}

	encontrado := true

	hash1.Iterar(func(clave K, valor V) bool {
		if !hash2.Pertenece(clave) || hash2.Obtener(clave) != valor {
			encontrado = false
			return false // dejar de iterar
		}
		return true // seguir iterando
	})

	return encontrado
}


func compararDosHashes[K comparable, V comparable](hash1, hash2 Diccionario[K,V]) bool {
	// verifico que los tamanos sean iguales
	if hash1.Cantidad() != hash2.Cantidad() {
		return false
	}
	// verifico que las claves sean iguales y que los valores tambien
	// establezco una variable que va a determinar el resultado
	encontrado := true

	// uso el iterador interno del hash
	hash1.Iterar(func(clave K, valor V) bool {
		if !hash2.Pertenece(clave) || hash2.Obtener(clave) != valor {
			encontrado = false
			return false
		}
		return true

	})
	return encontrado
}

func compararDosHashes[K comparable, V comparable](hash1, hash2 Diccionario[K,V]) bool {
	if hash1.Cantidad() != hash2.Cantidad() {
		return false
	}

	encontrado := true
	hash1.Iterar(func(clave K, valor V) bool {
		if !hash2.Pertenece(clave) || hash2.Obtener(clave) != valor {
			encontrado = false
			return false
		}
		return true
	})
	return encontrado
}


// PLANTEADO CON ITERADOR EXTERNO

func compararDosHashes[K comparable, V comparable](hash1,hash2 Diccionario[K,V]) bool {
	if hash1.Cantidad() != hash2.Cantidad() {
		return false
	}

	encontrado := true
	for iter := hash1.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, valor := iter.VerActual()
		if !hash2.Pertenece(clave) || hash2.Obtener(clave) != valor {
			encontrado = false
			return false
		}
	}
	return encontrado
}



/*
	Analisis de complejidad:
	La complejidad temporal de este algoritmo es O(n), donde en el peor de los casos
	se recorren n elementos, con n = hash1.Cantidad().

	Dentro del recorrido, para cada una de las n claves se realizan operaciones de hash
	(PERTENECE, OBTENER) que tienen costo O(1) en el caso promedio.
	Por lo tanto, se ejecutan n veces operaciones O(1), y el tiempo total resulta:
			O(n * 1) = O(n)
	Por otro lado, en cuanto a complejidad espacial, el algoritmo no utiliza estructuras
	adicionales de gran tamano. Unicamente se mantiene la variable encontrado (un booleano)
	mas algunas variables temporales del iterador, lo que determina que el uso de memoria
	adicional es O(1)
	
*/