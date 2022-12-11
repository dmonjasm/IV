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
+ Deuda técnica. En particular vamos a buscar que la herramienta reciba mantenimiento con asiduidad.


Voy a utilizar la información del repositorio [Awesome Go](https://github.com/avelino/awesome-go).

Si nos vamos al apartado de dicho repositorio de Build Automation voy a destacar Task y make:

+ Una opción sería [Task](https://github.com/go-task/task). 
    + Task recibe mantenimiento asiduamente, lo cual se puede apreciar en la página de github.

+ La última opción considerada es [Make](https://www.gnu.org/software/make/). También requiere la instalación de código externo, aunque suele venir como parte de muchas distribuciones Linux.

    + Make recibe mantenimiento asiduamente, como se puede apreciar en su documentación.


Hay otros *task runners*, aunque estos llevan bastante tiempo (meses o incluso años) sin recibir soporte, luego los he descartado automáticamente.

Task y Make son dos buenos candidatos para usar como gestor de tareas, pero he decidido utilizar Make. Make recibe soporte de forma regular, y gracias a su gran popularidad y la gran cantidad de años que lleva en desarrollo, no se prevee que en un futuro deje de recibirlo.

## Elección Gestor de Dependencias
La gestión de dependencias en GO se hace por medio la herramienta de línea de órdenes de GO. Es más, buscando gestores de dependencias para GO he visto que están obsoletos, y que los repositorios de los mismo han sido archivados.

Luego realmente no se requiere de un gestor de dependencias concreto. En este caso, como hemos usado Make como gestor de tareas simplemente añadimos al Makefile las claves necesarias para la gestión de dependencias.

## Órdenes Makefile
En este apartado se incluyen las distintas órdenes que acepta el Makefile del proyecto:
+ make check: comprueba la sintaxis de todos los archivos y devuelve los errores.
+ make build: compila todos los archivos del proyecto actual. Almacena los compilados en ./bin.
+ make installdeps: instala las dependencias de todos los archivos go que cuelgan del directorio raíz.
+ make clean: elimina los archivos compilados almacenados en ./bin.
+ make test: no implementado todavía. En un futuro ejecutará los tests.
+ make help: muestra las posible órdenes aceptadas por el make.
+ make: ejecuta check, build y clean.