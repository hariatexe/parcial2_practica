// DEVOLVER LAS CLAVES EN HASH ABIERTO

/*
	Para este caso, consideramos todas las listas, las cuales podemos ir iterando 
	utilizando el iterador externo o interno. Aquí mostramos una implementación 
	utilizando el iterador interno, no porque sea mejor implementación, sino para que 
	tengan un ejemplo de uso (ya que probablemente la primera idea que tengan ustedes 
	será con el externo).
*/


func (hash hashAbierto[K,V]) Claves() Lista[K] {
	claves := CrearListaEnlazada[K]()
	for _, lista := range hash.tabla {
		// Cada celda del hash contiene un lista con los elementos colisionados
		// Para este ejemplo itero con el iterador interno, que recibe una funcion
		lista.Iterar(func (par parClaveValor[K,V]) bool {
			claves.InsertarUltimo(par.clave)
			return true // Para que siga iterando
		})	
	}

	return claves
}

/*
	Aclaración: Hay quienes deciden implementar el Hash Abierto de tal forma que no 
	tenga listas vacías (esto es, si una posición aún no ha sido utilizada, entonces se 
	guarda nil y se crea la lista cuando sea necesaria). Esto es totalmente válido, 
	y en todo caso con aclararlo en algún lado es suficiente 
 	(pero debe hacerse la validación contra nil en ese caso).
*/

// EN VEZ DE USAR EL ITERADOR INTERNO AHORA USO EL EXTERNO

package claves

func (hash *hashAbierto[K, V]) Claves() Lista[K] {
    resultado := CrearListaEnlazada[K]()
	for _, lista := range hash.tabla {
		if lista.EstaVacia() {
			continue
		}
		iter := lista.Iterador()
		for iter.HaySiguiente() {
			par := iter.VerActual()
			resultado.InsertarUltimo(par.clave)
			// iterador interno de lista 
			// lista.Iterar(func (par parClaveValor[K,V]) bool{})
			// es decir que el iterador de lista (ya sea interno o externo)iter.VerActual() devuelve un par 
			// no como el iterador de hash (en este caso interno)
			// que devuelve hash.Iterar(func (clave K, valor V) bool {}), directamente devuelve clave K y valor V 
			iter.Siguiente()
		}
	}
	return resultado
}
















/// AHORA COMO FUNCION

/*
	En este caso, no solo si esta permitido utilizar el iterador externo del hash,
	sino que no nos queda otra opcion, dado que no es posible acceder a los campos
	internos de la estructura. Ademas, en particular no se nos dice cual es la implemen-
	tacion, cosa que no es necesario conocer.

	Entonces, simplemente iteramos utilizando el iterador externo y guardandolo
	en una lista
*/

func ClavesDiccionario(dicc Diccionario) Lista[K] {
	claves := CrearListaEnlazada[K]()
	for iter := dicc.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, _ := iter.VerActual()
		claves.InsertarUltimo(clave)
	}
	return claves
}