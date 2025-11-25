// SI EXIGE EFICIENCIA EN O(1) USAMOS OFFSET 
package addall

type dicAddAll[K comparable] struct {
    tabla Diccionario[K, int]
    delta int // offset global
}
type iterAddAll[K comparable] struct {
    interno IterDiccionario[K, int]
    delta   int
}


func CrearDicAddAll[K comparable]() DiccionarioAddAll[K] {
    return &dicAddAll[K]{tabla: CrearHash[K,int](), delta: 0 }
}

func (dic *dicAddAll[K]) Guardar(clave K, valor int) {
    hash := dic.tabla
    hash.Guardar(clave, valor - dic.delta)
}

func (dic *dicAddAll[K]) Pertenece(clave K) bool {
    return dic.tabla.Pertenece(clave)
}

func (dic *dicAddAll[K]) Borrar(clave K) int {
    hash := dic.tabla
    valor := hash.Obtener(clave)
    hash.Borrar(clave)
    return valor
}

func (dic *dicAddAll[K]) Obtener(clave K) int {
    hash := dic.tabla
    valor := hash.Obtener(clave) + dic.delta
    return valor
}

func (dic *dicAddAll[K]) Cantidad() int {
    return dic.tabla.Cantidad()
}

func (dic *dicAddAll[K]) Add(clave K, valor int) {
    valor_guardado := dic.tabla.Obtener(clave)
    valor_real := valor_guardado + dic.delta
    nuevo_valor_real := valor_real + valor
    dic.tabla.Guardar(clave, nuevo_valor_real - dic.delta)
}

func (dic *dicAddAll[K]) AddAll(valor int) {
    dic.delta += valor
}

func (dic *dicAddAll[K]) Iterar(visitar func(clave K, dato int) bool) {
    dic.tabla.Iterar(func (clave K, dato int) bool {
        return visitar(clave, dato + dic.delta)
    })
}

func (dic *dicAddAll[K]) Iterador() IterDiccionario[K, int] {
    return &iterAddAll[K]{
        interno: dic.tabla.Iterador(),
        delta:   dic.delta,
    }
}

func (it *iterAddAll[K]) HaySiguiente() bool {
    return it.interno.HaySiguiente()
}

func (it *iterAddAll[K]) VerActual() (K, int) {
    clave, valorGuardado := it.interno.VerActual()
    return clave, valorGuardado + it.delta
}

func (it *iterAddAll[K]) Siguiente() K {
    // Siguiente devuelve SOLO la clave real
    clave, valorGuardado := it.interno.VerActual()
    it.interno.Siguiente() // avanza el iterador interno
    _ = valorGuardado      // no lo usamos, pero está disponible
    return clave           // la clave no cambia por delta
}


/*
	Comentarios finales

	Insertar: ajusta por delta para mantener consistencia con AddAll.

	Obtener: devuelve el valor “real” sumando delta.

	Add: modifica solo la clave indicada.

	AddAll: solo incrementa el offset, sin recorrer la tabla.

	Borrar: delega al hash base.

	📌 Complejidad
	Operación	Complejidad	Justificación
	Insertar	O(1)	Inserción en hash + ajuste delta
	Obtener	O(1)	Lookup en hash + suma delta
	Borrar	O(1)	Delete en hash
	Add	O(1)	Lookup y actualización de un solo valor
	AddAll	O(1)	Solo incrementa delta
*/






// SI NO EXIGE EFICIENCIA 

// FORMA 1

package addall

type dicAddAll[K comparable] struct {
	diccionario Diccionario[K,int]
}

func CrearDicAddAll[K comparable]() DiccionarioAddAll[K] {
	diccionario := CrearHash[K,int]()
	//return &dicAddAll[K]{diccionario:diccionario}//
	// o mas simple
	return &dicAddAll[K]{diccionario}
}

func (dic *dicAddAll[K]) Guardar(clave K, valor int) {
	dic.diccionario.Guardar(clave, valor)
}

func (dic *dicAddAll[K]) Pertenece(clave K) bool {
	return dic.diccionario.Pertenece(clave)
}

func (dic *dicAddAll[K]) Borrar(clave K) int {
	return dic.diccionario.Borrar(clave)
}

func (dic *dicAddAll[K]) Obtener(clave K) int {
	return dic.diccionario.Obtener(clave)
}

func (dic *dicAddAll[K]) Cantidad() int {
	return dic.diccionario.Cantidad()
}

func (dic *dicAddAll[K]) Add(clave K, valor int) {
	if !dic.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}

	valor += dic.Obtener(clave)
	dic.Guardar(clave, valor)
}

func (dic *dicAddAll[K]) AddAll(valor int) {
	dic.diccionario.Iterar(func (clave K, dato int) bool {
		dic.Guardar(clave, dato+valor)
		return true
	})
	// Pierde la complejidad O(1), ya que itera por todos los elementos del
	// hash lo que converge a O(n) donde n es la cantidad de elementos del hash
}

func (dic *dicAddAll[K]) Iterar(visitar func(clave K, dato int) bool) {
	dic.diccionario.Iterar(visitar)
}

func (dic *dicAddAll[K]) Iterador() IterDiccionario[K,int] {
	return dic.diccionario.Iterador()
}