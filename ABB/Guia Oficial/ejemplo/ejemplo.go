/*
	Supongo que la estructura del TDA es:
	type abb struct {
		clave string
		izq *abb
		der *abb
	}

	La aclaracion que nos dice hacer la menor cantidad de operaciones, viene
	de que es necesario aprovechar la propiedad de abb donde todos los nodos a izquierda
	de un nodo raiz, son menores al nodo raiz, y todos losl nodos a derecha de un nodo
	raiz son mayores al nodo raiz

	Este ejercicio es parecido a una busqueda por rango pero sin limite superior
*/

// clave -> cadena donde queremos en este caso los mayores a cadena
/*
	Si encontramos un nodo menor o igual a cadena entonces es innecesario buscar
	a izquierda del nodo ya que sabemos que seran menores, por otro lado podriamos
	empezar buscar a derecha del nodo
*/

// claves son strings
func (abb *abb) Mayores(cadena string) Lista[string] {
	//mayores := TDALista.CrearListaEnlazada[string]()
	mayores := CrearListaEnlazada[string]()
	abb.mayores(cadena, mayores)
	return mayores
}

func (abb *abb[string]) mayores(clave string, mayores Lista[string]) {
	// caso BASE SIEMPRE
	if abb == nil {
		return 
	}

	// si la actual es mayor, llamamos a la izquierda y guardamos la actual
	// el abb.clave es mayor que la clave, por lo tanto es innecesario ir a la derecha
	if strings.strcmp(abb.clave, clave) > 0 {
		abb.izq.mayores(clave, mayores)
		mayores.InsertarUltimo(abb.clave)
	}

	abb.der.mayores(clave, mayores)
	// Va a ir a la izquierda todo lo que pueda, porque ahí pueden estar claves
	// que son mayores que 'clave' pero menores que la del nodo actual.
	//
	// Cuando ya no puede bajar más a la izquierda, la recursión empieza a volver
	// por la pila, y al volver a cada nodo se agrega su clave si es mayor que 'clave'.
	//
	// Después de agregar el nodo actual, sigue recursivamente hacia la derecha,
	// porque todo lo que está en el subárbol derecho es mayor que la clave del nodo
	// actual, y por lo tanto también mayor que 'clave'.

}

/*
		// abb.clave > clave
		// entonces la clave que tengo (clave buscada) es menor a abb.clave, por lo tanto
		// abb.clave debe ser incluida en el resultado.
		//
		// Antes de incluir abb.clave, tengo que revisar el subárbol izquierdo,
		// porque allí puede haber claves que también sean mayores que `clave`,
		// pero menores que abb.clave.
		//
		// Luego de procesar el subárbol izquierdo e insertar abb.clave,
		// sigo siempre hacia el subárbol derecho.
		//
		// ¿Por qué? Porque si abb.clave > clave, entonces **todas** las claves
		// que estén a la derecha de abb.clave también son mayores que `clave`
		// (por propiedad del ABB).
		//
		// En resumen:
		// - Si abb.clave > clave → exploro izquierda, agrego abb.clave y exploro derecha
		// - Si abb.clave ≤ clave → solo tiene sentido explorar derecha



		// abb.clave > clave
		// entonces digo que la clave tengo es menor a la abb.clave en la que estoy
		// por lo tanto tengo que revisar a izquierda para ver si hay mayores
		// al que tengo pero que son menores al que estoy viendo en abb.clave
		// siempre sigo para la derecha
		// una vez que acabe con la izquierda (viendo a la derecha) voy recursivamente
		// tomando todos las claves de los nodos que estan a la derecha de abb.clave
		// ya que si abb.clave es mayor a la clave todos los nodos a derecha tambien
		// lo son



		Supongamos este ABB:

		50
       /  \
     30    70
           / \
         60   80

    Y clave = 40.

	Como 50 > 40, entonces el código hace:

	abb.izq.mayores(40) → explora todo el subárbol izquierdo (30).

	La recursión termina y vuelve al nodo 50.

	Agrega 50 a la lista.

	abb.der.mayores(40) → explora todo el subárbol derecho (70, 60, 80).

	Y esto es exactamente lo que querés.

*/


