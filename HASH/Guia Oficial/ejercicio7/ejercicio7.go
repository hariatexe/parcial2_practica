type MultiConjunto[K comparable] interface {
	Guardar(elem K)
	Pertenece(elem K) bool
	Borrar(elem K) bool
}

type multiConjunto[K comparable] struct {
	hash Diccionario[K, int] // K es el elemento, int es la cantidad de apariciones
	// K = clave = elemento, valor = int = cantidad de apariciones
}

func CrearMultiConj[K comparable]() MultiConjunto[K] {
	return &multiConjunto[K]{hash: CrearHash[K, int]()}
}

func (conjunto *multiConjunto[K]) Guardar(elem K) {
	hash := conjunto.hash
	veces := 1
	if hash.Pertenece(elem) {
		veces += hash.Obtener(elem)
		hash.Guardar(elem, veces)
	} else {
		hash.Guardar(elem, veces)
	}
}	

func (conjunto *multiConjunto[K]) Pertenece(elem K) bool {
	return conjunto.hash.Pertenece(elem)
}

func (conjunto *multiConjunto[K]) Borrar(elem K) bool {
	hash := conjunto.hash
	if !hash.Pertenece(elem) {
		return false
	}
	contador := hash.Obtener(elem)
	if contador > 1 {
		hash.Guardar(elem, contador-1)
	} else {
		hash.Borrar(elem) // Cantidad de apariciones es 0, por lo tanto borro
	}
	return true
}




// ESCRIBO OTRA VEZ EL MULTICONJUNTO

type MultiConjunto[K comparable] interface {
	Guardar(elem K)
	Pertenece(elem K) bool
	Borrar(elem K) bool 
}

type multiConjunto[K comparable] struct {
	hash Diccionario[K, int]
}

func CrearMultiConj[K comparable]() MultiConjunto[K] {
	return &multiConjunto[K]{ hash: CrearHash[K,int]()}
}

func (conjunto *multiConjunto[K]) Guardar(elem K) {
	// creo el hash
	hash := conjunto.hash
	veces := 1
	if hash.Pertenece(elem) {
		veces += hash.Obtener(elem)
		hash.Guardar(elem, veces)
	}else {
		hash.Guardar(elem, veces)
	}
}

func (conjunto *multiConjunto[K]) Pertenece(elem K) bool {
	return conjunto.hash.Pertenece(elem)
}

func (conjunto *multiConjunto[K]) Borrar(elem K) bool {
	hash := conjunto.hash
	if !hash.Pertenece(elem) {
		return false
	}
	veces := hash.Obtener(elem)
	if veces > 1 {
		hash.Guardar(elem, veces-1)
	}else {
		hash.Borrar(elem)
	}
	return true
}


// ESCRIBO OTRA VEZ

// PRIMERO LA INTERFAZ

type MultiConjunto[K comparable] interface {
	Guardar(elem K)
	Pertenece(elem K) bool
	Borrar(elem K) bool
}

// defino el struct
type multiConjunto[K comparable] struct {
	hash Diccionario[K, int]
}

// defino la funcion que crea el multiconjunto

func CrearMultiConj[K comparable]() MultiConjunto[K] {
	return &multiConjunto[K]{hash: CrearHash[K,int]()}
} 

// creo las demas primitivas

func (conjunto *multiConjunto[K]) Guardar(elem K) {
	hash := conjunto.hash
	veces := 1
	if hash.Pertenece(elem) {
		veces += hash.Obtener(elem)
		hash.Guardar(elem, veces)
	} else {
		hash.Guardar(elem, veces)
	}
}


func (conjunto *multiConjunto[K]) Pertenece(elem K) bool{
	return conjunto.hash.Pertenece(elem)
}

func (conjunto *multiConjunto[K]) Borrar(elem K) bool{
	hash := conjunto.hash
	if !hash.Pertenece(elem) {
		return false
	}
	veces := hash.Obtener(elem)
	if veces > 1 {
		hash.Guardar(elem, veces-1)
	}else {
		hash.Borrar(elem)	
	}

	return true
}