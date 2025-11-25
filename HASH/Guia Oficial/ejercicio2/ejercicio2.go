/*
	A. hash_obtener(clave) no es un iterador, sino una operacion de busqueda
	que aplica la funcion hash a la clave y obtiene la posicion
	por otro lado la complejidad promedio esperada es O(1) (tiempo constante)


	* Para un HASH ABIERTO
	Cada posicion de la tabla contiene una lista, en el peor de los casos todas los elementos
	colisionan en la misma posicion por lo que la lista se vuelve O(n) y por ende la tabla hash
	se vuelve O(n)
	O no necesariamente se colisionan todos los elementos en la misma posicion, puede ser que
	el hash sea malo o el factor de carga alto (gran nivel de ocupacion en la tabla), lo que ocasiona
	que las listas tiendan a volverse largas es decir pasaron de hacer operacion de tiempo constante
	a tender a operaciones de tiempo lineal

	* Para un HASH CERRADO
	Recordemos que los elementos estan todos dentro del arreglo
	y las colisiones se resuelven buscando otra posicion
	hash_obtener() calcula la posicion hash, 
	si no coincide, sigue probando hasta encontrar la clave o una celda vacia
	El problema aparece cuando el factor de carga es alto (la tabla esta muy llena), o hay muchas
	celdas marcadas como borradas lo que genera un factor de carga que desestima la ocupacion
	real de la tabla, pero tambien que haya muchas colisiones
	es decir en el peor caso, es que al haber un factor carga elevado, se generen mas
	colisiones lo que termine haciendo que la funcion hash que obtiene el valor para x
	en la posicion 1, como hubo varios que colisionaron en la misma posicion,
	por cada salto que va haciendo hash_obtener() va tendiendo a O(n), en el peor
	de los casos hubo n coliciones lo que ocasiona que recorra n elementos la funcion
	hash_obtener()
	entonces la operacion hash_obtener() se puede llegar a volver O(n).

*/