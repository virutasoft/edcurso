package animales

type Mascota interface {
	Comunicarse() //para decirle a gato y a perro que implementen la inteface, deben tener el metodo que necesiten ussar y ya está, nada mas, es decir:  ele metodo Comunciarse()

	//LA ÚNICA RESTRICCIÓN COMO TAL DE LAS INTERFACES, ES QUE EN LA FUNCIÓN DE CADA UNO DE LOS MÉTODOS  QUE CONECTA, NO DEBE HABER UN PUNTERO EN EL TIPO DE DATO, EN ESTE CASO: p Perro, donde Perro es el tipo de dato
}
