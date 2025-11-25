/*
	Ejercicio resuelto
	a. Para un hash cerrado, implementar una primitiva func (hash hashCerrado[K, V]) Claves() 
	Lista[K] que reciba un hash y devuelva una lista con sus claves, sin utilizar el 
	iterador interno.

	b. Repetir lo mismo para un hash abierto.
	
	c. Implementar de nuevo la operación, en este caso como una función.
*/

/*
	Notaciones partiendo del punto a), es importante notar que no podemos
	utilizar el iterador externo, ya que se trata de una primitiva para el
	hash (y es una mala practica que, dado que el iterador dependa del hash,
	ahora hagamos que el hash dependa del iterador, lo que se llama
	dependencia circular).
*/

/*	Solucion Hash Cerrado */
// Para este punto, solo es necesario iterar campo por campo
// considerando unicamente aquellos que esten ocupados

// Es una lista de K - Keys / Claves
// V - Values / Valores
func (hash hashCerrado[K, V]) Claves() Lista[K] {
	claves := CrearListaEnlazada[K]()
	for _, celda := range hash.tabla {
		if celda.estado == OCUPADO {
			claves.InsertarUltimo(celda.clave)
		}
	}
	return claves
}
// devuelve Lista[K], que es la lista con las claves
 
/*
	En este caso se evalua que sabemos cuales son los campos de la estructura

	Sabemos que la tabla es de tipo []campo[K,V] o bien []celda[K,V]  y no
	[]*celda[K,V], puesto que es completamente innecesario otro grado de 
	indireccion (como se analiza en la respectiva clase practica)

	Entendemos que el estado correcto a considerar es el de OCUPADO, que es un
	enumerativo. Cuanto mucho una constante. Definitivamente no un "ocupado"
*/


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




