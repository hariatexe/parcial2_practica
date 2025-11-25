// PRIMITIVA HASH CERRADO - OBTENCION DE CLAVES
func (hash hashCerrado[K,V]) Claves() Lista[K] {
	claves := CrearListaEnlazada()[K] // Lista que es de tipo claves
	for _, celda := range hash.tabla {
		if celda.estado == OCUPADO {
			claves.InsertarUltimo(celda.clave)
		}
	}
	return claves
}

func (hash hashCerrado[K,V]) Claves() Lista[K] {
	claves := CrearListaEnlazada[K]()
	for _, celda := range hash.tabla {
		if celda.estado == OCUPADO {
			claves.InsertarUltimo(celda.clave)
		}
	}
}

func (hash hashCerrado[K,V]) Claves() Lista[K] {
	claves := CrearListaEnlazada[K]()
	for _, celda := range hash.tabla {
		if celda.estado == OCUPADO {
			claves.InsertarUltimo(celda.clave)
		}
	}
}


// PRIMITIVA HASH ABIERTO - OBTENCION DE CLAVES

// Utilizo el iterador de la lista
// Aqui ya no tengo celdas con estados

func (hash hashAbierto[K,V]) Claves() Lista[K] {
	claves := CrearListaEnlazada[K]()

	for _, lista := range hash.tabla {
		lista.Iterar(func (par parClaveValor[K,V]) bool {
			claves.InsertarUltimo(par.clave)
			return true
		})
	}

	return claves
}

func (hash hashAbierto[K,V]) Claves() Lista[K] {
	claves := CrearListaEnlazada[K]()
	for _, lista := range hash.tabla {
		lista.Iterar(func (par parClaveValor[K,V]) bool {
			claves.InsertarUltimo(par.clave)
			return true
		})
	}
}


func (hash hashAbierto[K,V]) Claves() Lista[K] {
	claves := CrearListaEnlazada[K]()
	for _, lista := range hash.tabla {
		lista.Iterar(func (par parClaveValor[K,V]) bool {
			claves.InsertarUltimo(par.clave)
			return true
		})
	}
}


// FUNCION HASH - OBTENCION DE CLAVES
// YA QUE NO SE PUEDE ACCEDER A LA ESTRUCTURA INTERNA, ENTONCES HAY QUE USAR
// EL ITERADOR EXTERNO DEL HASH (No nos importa la implementacion, es decir si es abierto o cerrado)

func Claves(dicc Diccionario) Lista[K] {
	claves := CrearListaEnlazada[K]()
	// dicc.Iterador() instancia el iterador externo del hash 	
	for iter := dicc.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, _ := iter.VerActual()
		claves.InsertarUltimo(clave)		
	}
	return claves
} 


func Claves(dicc Diccionario) Lista[K] {
	claves := CrearListaEnlazada[K]
	for iter := dicc.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, _ := iter.VerActual()
		claves.InsertarUltimo(clave)
	}
	return claves
}


func Claves(dicc Diccionario) Lista[K] {
	claves := CrearListaEnlazada[K]
	for iter := dicc.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, _ := iter.VerActual()
		claves.InsertarUltimo(clave)
	}
	return claves
}

func Claves(dicc Diccionario) Lista[K] {
	claves := CrearListaEnlazada[K]
	for iter := dicc.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, _ := iter.VerActual()
		claves.InsertarUltimo(clave)
	}
	return claves
}


// AHORA UTILIZANDO EL ITERADOR INTERNO

// RECORDEMOS LA ANTERIOR IMPLEMENTACION DE LA PRIMITIVA CON HASH ABIERTO

func (hash hashAbierto[K,V]) Claves() Lista[K] {
	claves := CrearListaEnlazada[K]()
	for _, lista := range hash.tabla {
		lista.Iterar(func (par parClaveValor[K,V]) bool {
			claves.InsertarUltimo(par.clave)
			return true
		})
	}
	return claves
}
// AHORA LO HAGO UTILIZANDO EL ITERADOR INTERNO DEL HASH

func (hash hashAbierto[K,V]) Claves() Lista[K] {
	claves := CrearListaEnlazada[K]()
	hash.Iterar(func (clave K, valor V) bool {
		claves.InsertarUltimo(clave)
		return true
	})

	return claves
}

// VALE TAMBIEN PARA EL CERRADO Y COMO FUNCION

// CASO HASH CERRADO - PRIMITIVA CON ITERADOR INTERNO
func (hash hashCerrado[K,V]) Claves() Lista[K] {
	claves := CrearListaEnlazada[K]()
	hash.Iterar(func (clave K, valor V) bool {
		claves.InsertarUltimo(clave)
		return true
	})
	return claves
}

// CASO FUNCION - NO IMPORTA LA IMPLEMENTACION DEL HASH SI ES ABIERTO O CERRADO
func ClavesDiccionario(dicc Diccionario) Lista[K] {
	claves := CrearListaEnlazada[K]()
	dicc.Iterar(func (clave K, valor V) bool {
		claves.InsertarUltimo(clave)
		return true
	})
	return claves
}
