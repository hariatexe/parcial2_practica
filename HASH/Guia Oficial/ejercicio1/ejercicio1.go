/*
	Respuesta:

	 A. En el peor de los casos ya que la implementacion no considera los borrados en el factor de carga
	 el hash parece estar poco lleno de lo que en realidad esta, es decir si solo se guia de los 
	 elementos o celdas ocupadas el error esta en que en verdad en las celdas borradas siguen
	 habiendo elementos mugre, por lo que el factor de carga verdaderamente es mucho mayor de lo que
	 realmente se cree, por lo que si bien el hash puede decir que tiene un factor de carga muy
	 chico que no amerita redimension, la verdad es que necesita una redimension por tener
	 muchos lugares o celdas ocupadas
	 Siendo el factor de carga = a las celdas ocupadas sobre el tamano del hash
	 Es decir que el peor escenario es que tenga muchas celdas borradas que contradigan
	 lo que el factor de carga indica, como no se redimensiona, las operaciones de insercion y busqueda
	 se vuelven MUY LENTAS casi O(n) por la gran cantidad de celdas "sucias".
	 Si la implementacion no reutiliza las celdas borradas empeora el escenario ya que busca celdas
	 vacias


	 B. hash => capacidad inicial 100
	 hash_obtener(), cantidad actual de elementos en el hash es 1
	 Observacion: no se realizo ninguna redimension, pero si se insertaron y borraron elementos
	 Como dice que la cantidad actual de elementos en el hash es 1, quiere decir que todos los demas
	 elementos que se insertaron se borraron, es decir que la cantidad de elementos en el hash
	 contando las celdas ocupadas mas las celdas borradas, tenemos un total de 1 + k (elementos borrados)
	 por lo tanto si la capacidad inicial es de 100 y no hay redimension ya que el factor de carga
	 cuenta que hay 1 cocmo cantidad de ocupados y 100 como cantidad total de tamano, por lo que el factor
	 de carga es 1/100, pero en realidad el factor contando los borrados es 1+k / 100, siendo k la cantidad de borrados.
	 Podria decir entonces que el numero maximo de casillas
	 que podria llegar a visitar hash_obtener() es de un total de k+1 => k (casillas borradas)
	 con k siendo n-1 en el peor de los casos, n como el total de elementos, en el peor de los 
	 casos hash_obtener() va a visitar las 100 casillas

*/