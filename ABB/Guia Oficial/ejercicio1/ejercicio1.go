func (abb *abb[string]) Altura() int{
	if abb == nil {
		return 0
	}
	abbIzq := abb.izq.Altura()
	abbDer := abb.der.Altura()
	return 1 + max(abbIzq, abbDer)
} 


/*
	Cada nodo se visita una vez 
*/