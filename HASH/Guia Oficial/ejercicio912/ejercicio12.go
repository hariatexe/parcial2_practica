
/*
	No se puede hacer lo siguiente:
func MergeProveedores[string, int](prov1, prov2 Diccionario[string,int]) Diccionario[string,int]{}
Al no ser genericos string ni int, no hace falta instanciarlos

Es decir en el unico caso que se instancian es cuando se tienen GENERICOS, por ejemplo:

func MergeProveedores[K comparable, V any](prov1, prov2 Diccionario[K,V]) Diccionario[K,V]{}

*/

func MergeProveedores(prov1, prov2 Diccionario[string,int]) Diccionario[string,int] {
	hashResultado := CrearHash[string, int]()
	for iter := prov1.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, dato := iter.VerActual()
		if prov2.Pertenece(clave) {
			dato2 := prov2.Obtener(clave)
			if dato2 < dato {
				dato = dato2
			}
		}
		hashResultado.Guardar(clave, dato)
	}

	for iter := prov2.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, dato := iter.VerActual()
		if !prov1.Pertenece(clave) {
			hashResultado.Guardar(clave, dato)
		}
	}

	return hashResultado
}



func obtenerHash[K comparable, V any](hash1, hash2 Diccionario[K,V]) Diccionario[K,V] {
	hashResultado := CrearHash[K,V]()
	hash1.Iterar(func (clave K, valor V) bool {
		if hash2.Pertenece(clave) {
			valorclave2 := hash2.Obtener(clave)
			if valorclave2 <= valor {
				hashResultado.Guardar(clave, valorclave2)
			} else {
				hashResultado.Guardar(clave, valor)
			}
		}else {
			hashResultado.Guardar(clave, valor)
		}
		return true
	})

	// lo mismo para el hash2
	hash2.Iterar(func (clave K, valor V) bool {
		if !hash1.Pertenece(clave) {
			hashResultado.Guardar(clave, valor)
		}
		return true
	})

	return hashResultado
}



/*

	En este algoritmo tengo que recorrer todos los elementos de d1 (hash1)
	denotemos con la letra n, y ademas lodos los elementos de d2 (hash2), de-
	notemos con la letra m, dentro de los dos bucles hago operaciones O(1) en 
	promedio
	como Pertenece(), Obtener(), Guardar().
	La complejidad del primer bucle es O(n)
	La complejidad del segundo bucle es O(m)
	La complejidad total del algoritmo es O(n+m)

*/

// Si V any es NO ENUMERABLE, entonces las comparaciones no se podran hacer
// al saber que se trata de precios puedo usar float64

func obtenerHash[K comparable](hash1, hash2 Diccionario[K, float64]) Diccionario[K, float64] {
	hashResultado := CrearHash[K,float64]()
	hash1.Iterar(func (clave K, valor float64) bool {
		if hash2.Pertenece(clave){
			valor2 := hash2.Obtener(clave)
			if valor2 < valor {
				valor = valor2
			} 
		}
		hashResultado.Guardar(clave, valor)
		return true
	})

	hash2.Iterar(func (clave K, valor float64) bool {
		if !hash1.Pertenece(clave) {
			hashResultado.Guardar(clave, valor)
		}
		return true
	})

	return hashResultado
}
