# RouteCheck

## Descripción del problema
A la hora de realizar ciertos viajes en automóvil debemos pasar por un peaje o varios, y con el precio actual de la gasolina no sabemos si es realmente rentable pagar dicho peaje, o ahorrarnos el precio de ese peaje y realizar unos cuantos kilómetros extra. A parte de esto, tampoco sabemos si la diferencia en precio va a justificar el ahorro de tiempo que nos puede conllevar ir por peaje.

## Solución
Teniendo en cuenta nuestro automóvil y el precio actual de la gasolina, conocer de antemano cuánto dinero aproximadamente nos va a costar el ir por peaje, o evitar el peaje. En el caso de que nuestra ruta pueda atravesar varios peajes probar las distintas combinaciones que haya y ver cual es la más rentable. Teniendo en cuenta el precio calculado y el tiempo empleado en cada una de las rutas seleccionar una que sea la más favorable en relación precio/tiempo.

## Configuración git
Para acceder a la configuración pinchar este [enlace](docs/config.md)


## Lista de comprobación
* [x] ¿Se trata de un problema real del que se tenga conocimiento personal?
Sí, durante 4 años he estado trabajando en Madrid y he tenido que realizar la ruta Madrid-Segovia diariamente. En esta ruta hay un peaje en el cual, en función del periodo y del precio que tenga la gasolina dicho día, sale más rentable evitarlo que cogerlo.

* [x] ¿Se trata de un problema que para solucionar requiera el despliegue de una aplicación en la nube?
Sería necesario recoger los precios de los distintos peajes repartidos por España, así como los precios que tiene cada uno para los distintos tipos de vehículos y periodos horarios del día.

* [x] ¿La solución requiere una cierta cantidad de lógica de negocio, en vez
solucionarse sólo almacenando y buscando?
Sí, va a ser necesario hacer Web Scraping y probar varias combinaciones de rutas, así como hacer una cierta heurística para elegir una ruta teniendo en cuenta su precio y tiempo total.

* [x] ¿Se ha incluído la configuración del repositorio y se ha enlazado desde el README?
Sí, todo queda enlazado en el apartado configuración git del README.

## Elección Gestor de Tareas
Para la elección del gestor de tareas se van a tener en cuenta:
+ Estándares de GO.
+ Mejores Prácticas.
+ Deuda técnia. Soporte. Documentación.
+ Prestaciones.

Voy a utilizar la información del repositorio [Awesome Go](https://github.com/avelino/awesome-go).

Si nos vamos al apartado de dicho repositorio de Build Automation voy a destacar Task y el *task runner* implícito de GO.

+ El *task runner* implícito de GO se trata de una herramienta para la gestión de código GO. Como consecuencia de ir implícito en GO tiene un amplio soporte, pues se ve actualizado con el propio lenguaje. Luego cumple el apartado de deuda técnica. Respecto a estándares GO recomienda utilizar el *task runner* implícito para la automatización de tareas (aunque es normal, pues el mismo lo diseña), aunque algunos proyectos combinan el anterior con Makefile.

+ Otra opción sería [Task](https://github.com/go-task/task). Este al igual que el anterior cumple el requisito de la deuda técnica, pues tiene actualizaciones regulares. Si bien cumple la deuda técnica creo que es mejor el *task runner* implícito, pues si bien actualmente Task recibe soporte en un futuro no se sabe si seguirá recibiéndolo, mientras que el anterior recibirá soporte siempre que GO lo reciba. Por otro lado, este *task manager* no es estándar.

Hay otros *task runners*, aunque estos llevan bastante tiempo (meses o incluso años) sin recibir soporte, luego los he descartado automáticamente.

He decidido utilizar el *task runner* implícito de GO. Se trata del estándar que recomienda GO y además tiene un amplio soporte, mientras que Task solo cumple el criterio de deuda técnica.

Por otro lado, en los requisitos del objetivo 3 nos pide que incluyamos una clave fichero en el iv.yaml que apuntará al archivo que se usará para ejecutar las tareas. Debido a esto he decido combinar Make con el *task runner* implícito de GO.

Creo que a pesar de combinar con Make sigue siendo mejor opción, ya que buscando comparaciones entre Make y Task, no he encontrado una clara ventaja para elegir uno sobre otro. Luego la deuda técnica y los estándares es lo que me ha decantado por Make+GO.

## Órdenes Makefile
En este apartado se incluyen las distintas órdenes que acepta el Makefile del proyecto:
+ make check: comprueba la sintaxis de todos los archivos y devuelve los errores.
+ make build: compila todos los archivos del proyecto actual. Almacena los compilados en ./bin.
+ make installdeps: instala las dependencias de todos los archivos go que cuelgan del directorio raíz.
+ make clean: elimina los archivos compilados almacenados en ./bin.
+ make test: no implementado todavía. En un futuro ejecutará los tests.
+ make help: muestra las posible órdenes aceptadas por el make.
+ make: ejecuta check, build y clean.

