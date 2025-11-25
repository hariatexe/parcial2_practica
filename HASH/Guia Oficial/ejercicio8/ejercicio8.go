/*
	8. ★★★) Se tiene un hash que cuenta con una función de hashing que, recibida una clave, 
devuelve la posición de su inicial en el abecedario. La capacidad inicial del hash es 26. 
Para los puntos B, C y D indicar y justificar si las afirmaciones son verdaderas o falsas. 
Se puede considerar que todas las claves serán palabras (sólo se usan letras para las claves).

a. Mostrar cómo quedan un hash abierto y un hash cerrado (sólo el resultado final) tras guardar 
las siguientes claves: Ambulancia (0), Gato (6), Manzana (12), Ananá (0), Girasol (6), Zapato (25), 
Zapallo (25), Manzana (12), Bolso (1). Aclaración: No es necesario hacer una tabla de 26 posiciones, 
lo importante es que quede claro en cuál posición está cada elemento.

b. En un hash abierto con dicha función de hashing, se decide redimensionar cuando la cantidad alcanza 
la capacidad (factor de carga = 1). El rendimiento de hash_obtener() es mejor en este caso respecto a 
si se redimensionara al alcanzar un factor de carga 2.

c. En un hash cerrado con dicha función de hashing, si se insertan n + 1 claves diferentes 
(considerar que se haya redimensionado acordemente), n con la misma letra inicial, y 1 con otra 
distinta, en el primer caso Obtener() es O(n) y en el segundo siempre O(1).

d. En un hash abierto con dicha función de hashing, si se insertan n + 1 claves diferentes 
(considerar que se haya redimensionado acordemente), n con la misma letra inicial, y 1 con otra 
distinta, en el primer caso Obtener() es O(n) y en el segundo siempre O(1).



1.A
Primer clave: Ambulancia, le aplico el hash: hash(Ambulancia) = 0 => 0%26 = 0
Segunda clave: Gato, le aplico el hash: hash(Gato) = 6 => 6%26 = 6
Tercer clave: Manzana, hash(Manzana) = 12 => 12%26 = 12
Cuarta clave: Anana, hash(Anana) = 0 => 0%26 = 0
Quinta clave: Girasol, hash(Girasol) = 6 => 6%26 = 6
Sexta clave: Zapato, hash(Zapato) = 25 => 25%26 = 25 
Septima clave: Zapallo, hash(Zapallo) = 25 => 25%26 = 25
Octava clave: Manzana, hash(Maznana) = 12 => 12%26 = 12
Novena clave: Bolso, hash(Bolso) = 1 => 1%26 = 1

// Es una hash que tiene una capacidad inicial de 26 posiciones, por lo que para calcular
la posicion de cada elemento que obtuvimos aplico el modulo al hash 


HASH CERRADO
[Ambulancia] 
[Anana]
[Zapallo]
[Bolso]
[]
[]
[Gato]
[Girasol]
[]
[]
[]
[]
[Manzana]
[Manzana]
[]
[]
[]
[]
[]
[]
[]
[]
[]
[]
[]
[Zapato]

HASH ABIERTO

[[Ambulancia,Anana]]
[[Bolso]]
[]
[]
[]
[]
[[Gato,Girasol]]
[]
[]
[]
[]
[]
[[Manzana,Manzana]]
[]
[]
[]
[]
[]
[]
[]
[]
[]
[]
[]
[]
[[Zapato,Zapallo]]


RESPUESTAS VERDADERO FALSO:
1. En un hash abierto con dicha función de hashing, se decide redimensionar cuando la cantidad 
alcanza la capacidad (factor de carga = 1). El rendimiento de hash_obtener() es mejor en este caso 
respecto a si se redimensionara al alcanzar un factor de carga 2.


VERDADERO, EN ESTE CASO AL TENER FACTOR DE CARGA 2 EL ITERADOR EN CADA BUCKET ITERARA EL DOBLE
DE LO QUE ITERA EN UNA HASH DE FACTOR 1. MIENTRAS MAS LLENA ESTE LA TABLA (FACTOR > 1), LAS 
LISTAS EN LOS BUCKETS SERAN MAS LARGAS, ASI QUE PARA BUSCAR UN ELEMENTO TENDRE QUE RECORRER
MAS NODOS.
CON FACTOR DE CARGA = 1, LAS LISTAS ESTAN MAS CORTAS, MENOS ELEMENTOS POR BUCKET, BUSQUEDAS
MAS RAPIDAS

2. En un hash cerrado con dicha función de hashing, si se insertan n + 1 claves diferentes 
(considerar que se haya redimensionado acordemente), n con la misma letra inicial, y 1 con otra distinta, 
en el primer caso Obtener() es O(n) y en el segundo siempre O(1).

FALSO, SI TENGO N CLAVES QUE COLISIONAN (MISMA LETRA INICIAL), TODOS SE ENCUENTRAN EN POSICIONES
CONSECUTIVAS SEGUN EL SONDEO
BUSCAR CUALQUIERA DE ESAS N CLAVES ES O(N) EN EL PEOR CASO, PERO BUSCAR LA QUE TIENE OTRO HASH
DISTINTO NO ES O(1) NECESARIAMENTE: si el bucket donde debería ir está ocupado por otras claves, 
el sondeo tiene que avanzar hasta encontrar un lugar vacío o la clave buscada.
PUEDE TENDER A O(N)

3. En un hash abierto con dicha función de hashing, si se insertan n + 1 claves diferentes 
(considerar que se haya redimensionado acordemente), n con la misma letra inicial, y 1 con otra distinta, 
en el primer caso Obtener() es O(n) y en el segundo siempre O(1).

FALSO, LA BUSQUEDA DE LA CLAVE CON MUCHAS COLISIONES ES O(N) (LISTA LARGA). 
LA BUSQUEDA DE LA CLAVE CON HASH DISTINTO ES O(1) PROMEDIO SI NO HAY COLISION EN SU BUCKET.
PERO PODRIA SER O(K), SI ESE BUCKET TIENE K ELEMENTOS DEBIDO A COLISIONES PREVIAS.
TODAS LAS CLAVES CON IGUAL LETRA INICIAL ESTAN EN UN BUCKET, ASI QUE EL BUCKET SIGUIENTE
POSIBLEMENTE ES EL DISTINTO. POR 
ESO EN LA BUSQUEDA DEL SEGUNDO NO SIEMPRE ES O(1)

EN UN HASHING ABIERTO PARA OBTENER
LA FUNCION HASHING DICE EN QUE BUCKET BUSCAR LA CLAVE
LUEGO EL ITERADOR RECORRE SOLAMENTE LA LISTA DE ESE BUCKET BUSCANDO LA CLAVE, 
POR ESO SI TIENE K COLISIONES PREVIAS SERIA O(K)
OSEA QUE NI VE EL BUCKET DE N ELEMENTOS DE MISMA INICIAL.


EN EL PEOR DE LOS CASOS EN EL ABIERTO TIENDE A O(K) SIENDO K LA CANTIDAD DE COLISIONES PREVIAS

*/