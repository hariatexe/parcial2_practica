func DiferenciaSimetricaDict[K comparable, V any](d1, d2 Diccionario[K,V]) Diccionario[K,V] {
	// La diferencia simetrica que se necesita son todos los elementos de d1 que no estan
	// en d2 y todos los elementos de d2 que no estan en d1
	hashResultado := CrearHash[K,V]()

	// ultilizo el iterador externo ya que no se la implementacion al ser una funcion
	for iter := d1.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, valor := iter.VerActual()
		if !d2.Pertenece(clave){
			hashResultado.Guardar(clave, valor)
		}
	}

	for iter := d2.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, valor := iter.VerActual()
		if !d1.Pertenece(clave) {
			hashResultado.Guardar(clave, valor)
		}
	}

	return hashResultado // seria la diferencia simetrica
}

/*
	Analisis de complejidad:
	El algoritmo recorre todos los elementos de d1 (hash 1) y luego todos los elementos
	de d2 (hash 2).
	Supongamos que d1 tiene n elementos y d2 tiene m elementos.
	Cada iteracion del bucle realiza:
	Una operacion Pertenece, que es O(1) promedio en un hash
	Una operacion opcional Guardar, tambien O(1) promedio

	Por lo tanto:
	El primer bucle cuesta O(n)
	El segundo bucle cuesta O(m)

	Sumando ambos costos:
	Complejidad total: O(n+m)
*/