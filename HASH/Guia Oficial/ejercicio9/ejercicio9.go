/*
	9. (★★★) El Ing. Musumeci quiere implementar un hash abierto, pero en el que las listas de cada 
	posición se encuentren ordenadas por clave (se le pasa por parmámetro la función de comparación, 
	por ejemplo strcmp). Explicar cómo mejora o empeora respecto a la versión que vemos en clase 
	para el caso de inserciones, borrados, búsquedas con éxito (el elemento se encuentra en el hash)
	y sin éxito (no se encuentra).

	RESPUESTA:
	Al estar ordenado, podemos anticipar dentro de una lista en que seccion de la misma se va a encontrar
	una determinada clave, 
	si por ejemplo es una lista de claves ordenadas por un hash que devuelve dependiendo de la primera
	inicial de la clave, entonces sabemos que al tener una lista ordenada, y dependiendo de
	la ordenacion interna de la lista supongamos que ordena por la segunda letra entonces si queremos
	buscar una clave que es ab, y vemos en la lista lo siguiente: [aa,ac,...,az] entonces quiere decir
	que en las primeras dos iteraciones lo encontramos y por ende en el resto de la lista no se encontrara
	la clave buscada, para el caso de borrados pasa exactamente lo mismo.
	En el caso de inserciciones seria mas rapido en el sentido que no tiene que llegar al final de la lista
	para introducir el nuevo elemento si no que al saber quien es el anterior y quien es el posterior
	puedo meterlo en el medio

	CORREGIDO:
	* Busqueda: 
	El hecho de que las listas esten ordenadas permite cortar antes la busqueda. 
	Si al recorrer una lista ordenada encuentro una clave mayor a la busqueda, puedo detenerme por que se
	que ya no aparecera mas adelante

	. Busqueda con exito (el elemento se encuentra en el hash): en el peor de los casos seria
	O(L) siendo L la cantidad de elementos en la lista, pero puede ser mas rapida en promedio
	. Busqueda sin exito (el elemento no se encuentra en el hash): mejora por que se puede cortar
	antes la iteracion (promedio mejor que O(L))

	* Inserciones:
	En la version no ordenada, la insercion es O(1), simplemente se agrega al inicio o al final de la lista
	mientras que en la version ordenada, primero hay que buscar la posicion correcta lo que llevara a O(l)
	Entonces podria decirse que empeora el tiempo de insercion

	* Borrar: 
	Coincide con la busqueda, hay que recorrer la lista hasta encontrar la clave O(L), y luego borrarla
	O(1) adicional, hay una leve mejora por el corte anticipado, si no se encuentra la clave

	en el peor caso el borrar y busqueda es igual en la lista ordenada como no ordenada
	por que tienen que recorrer todo, pero en el caso promedio la lista ordenada es O(L) mientras
	que en el caso promedio de la no ordenada es O(N): 
	En la lista no ordenada, el caso promedio es O(N), siendo N la cantidad total de elementos en el bucket, ya que en general se recorre casi toda la lista.
	En cambio, en la lista ordenada, el recorrido puede cortarse antes, siendo en promedio O(L), 
	donde L es la cantidad de elementos recorridos hasta el corte (L < N).
*/
