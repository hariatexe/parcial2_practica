func MasDeLaMitad(arr []int) bool {
	// planteo un hash
	diccionario := CrearHash[int,int]()
	if len(arr) == 1 {
		return true
	}
	if len(arr) == 0 {
		return false 
	}
	for _, elem := range arr {
		veces := 1
		if diccionario.Pertenece(elem){
			veces += diccionario.Obtener(elem)
			diccionario.Guardar(elem, veces)
			if veces > len(arr) / 2 { // Si aparece mas de la mitad de veces
				return true
			}
		}else {
			diccionario.Guardar(elem, veces)
		}
	}
	return false 
}

/*
	ORDEN DE LA FUNCION:
	El algoritmo itera por cada elemento del arreglo guardando en el hash
	las apariciones lo cual son operaciones O(1), al hacer operaciones O(1) por
	cada elemento del arreglo eso queda en n por O(1), la complejidad es O(n),
	en el peor de los casos podria llegar a ser O(n^2), supongamos el caso
	donde todos los elementos del hash colisionen en una misma posicion
	entonces se iterara primero por los elementos del arreglo
	y luego a medida que se agregan elementos en la misma posicion (ya sea por
	error de la implementacion o un hash deficiente) cada vez que se quiera
	obtener o actualizar un elemento se tendra que recorrer lo que la lista 
	tenga en el caso de que sea un hashAbierto, converge a O(v), que en el
	peor caso podria llegar a ser n por lo
	tanto la complejidad del algoritmo:
	O(n) -> Caso Promedio
	O(n^2) -> Peor Caso


	✅ Ejemplo de peor caso (hash abierto)

	Supongamos:

	El arreglo tiene n elementos.

	La función hash es pésima, y para toda clave devuelve siempre la misma posición.

	Entonces todos los elementos terminan en el mismo bucket (una única lista interna).

	Consecuencia:

	El primer elemento se inserta en O(1).

	El segundo también, pero la lista ya tiene 1 elemento.

	El k-ésimo elemento se inserta en O(k), porque la lista tiene tamaño k−1.

	Entonces:
					n
	Costo total = SIGMA O(K) = O(n^2)
					K=1

	Y por lo tanto:

	Cada inserción converge a O(n) en el peor caso.

	El algoritmo completo converge a O(n²).
	 
*/