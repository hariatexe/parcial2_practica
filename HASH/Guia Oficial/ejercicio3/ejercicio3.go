/*
	RESPUESTA:
	A. Esta no es una funcion de hashing ya que no es determinista. Dos llamadas con la misma clave
	pueden devolver valores distintos cada vez.
	Ejemplo: calcularHash("perro") puede dar 38 y luego 32948 otra.
	Esto rompe totalmente con la logica de un hash: la misma clave debe producir siempre el mismo valor
	hash, si no, no se puede buscar nada. 
	Por consecuencia no depende de la clave.
	El hash deberia usar los caracteres de la clave para calcular el numero, aqui simplemente
	genera un numero aleatorio, no relacionado con la clave.
	Ademas genera colisiones innecesarias, como solo hay 10000 posibles resultados (0-9999), si hay muchas
	claves distintas, las probabilidades de colision crecen rapido.
*/