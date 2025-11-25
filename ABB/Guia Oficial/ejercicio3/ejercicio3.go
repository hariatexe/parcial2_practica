func datoONeutro(ab *arbol[int]) int {
	if abb == nil {
		return 1 // hablamos de multiplicacion el neutro es el 1
	} else {
		return ab.dato
	}
}

func MultiplicarConHijosPre(arbol *Arbol[int]) {
	if arbol == nil {
		return
	}

	valor_izq := datoONeutro(arbol.izq)
	valor_der := datoONeutro(arbol.der)
	// pre-order
	arbol.dato += valor_izq * valor_der
	MultiplicarConHijosPre(arbol.izq)
	MultiplicarConHijosPre(arbol.der)
}

func MultiplicarConHijosIn(arbol *Arbol[int]) {
	if arbol == nil {
		return
	}

	MultiplicarConHijosIn(arbol.izq) // primero a la izquierda
	valor_izq := datoONeutro(arbol.izq)
	valor_der := datoONetruo(arbol.der)
	arbol.dato += valor_izq * valor_der // al medio
	MultiplicarConHijosIn(arbol.der) // luego a la derecha
}


func MultiplicarConHijosPos(arbol *Arbol[int]) {
	if arbol == nil {
		return
	}
	MultiplicarConHijosPos(arbol.izq) // primero a la izquierda
	MultiplicarConHijosPos(arbol.der) // luego a la derecha
	valor_izq := datoONeutro(arbol.izq)
	valor_der := datoONeutro(arbol.der) 
	arbol.dato += valor_izq * valor_der // al medio
}

