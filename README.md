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
+ Instalación. Si la requiere, y en caso de así ser, cómo de compleja es la misma.


Voy a utilizar la información del repositorio [Awesome Go](https://github.com/avelino/awesome-go).

Si nos vamos al apartado de dicho repositorio de Build Automation voy a destacar Task y make:

+ Una opción sería [Task](https://github.com/go-task/task). El principal problema de este es que requiere la instalación de código externo, aunque esta es bastante simple.

    + Este al igual que el anterior cumple el requisito de mantenimiento, pues tiene actualizaciones regulares. Si bien cumple el requisito de mantenimiento creo que son mejores el *task runner* implícito o Make, pues si bien actualmente Task recibe soporte, en un futuro no se sabe si seguirá recibiéndolo, mientras que las otras dos opciones debido a que una se actualiza con el propio lenguaje y la otra es tremendamente popular, garantizan su soporte durante los próximos años.

    + Por otro lado, este *task manager* no es estándar, surge como alternativa más simple que Make. Este utiliza archivos .yml o .yaml para la automatización de tareas.

    + Respecto a las funcionalidades ofrecerá todas aquellas que se implementen en el archivo *Taskfile*. Esto incluye todas las funcionalidades del anterior.

+ La última opción considerada es [Make](https://www.gnu.org/software/make/). También requiere la instalación de código externo, aunque suele venir como parte de muchas distribuciones Linux.

    + Make sigue recibiendo soporte, en concreto la última actualización es del 31 de Octubre de 2022, luego cumple el requisito del mantenimiento. Además a diferencia de Task, debido a la gran popularidad de Make es probable que reciba soporte durante muchos más años.

    + Este *task runner* no es el estándar de GO, aunque es ampliamente utilizado.

    + En cuanto a funcionalidades, ocurre como con Task, tendrá todas las funcionalidades que tenga el archivo Makefile implementadas. De nuevo, esto incluye todas las funcionalidades del *task runner* implícito de GO.


Hay otros *task runners*, aunque estos llevan bastante tiempo (meses o incluso años) sin recibir soporte, luego los he descartado automáticamente.

Task y Make son dos buenos candidatos para usar como gestor de tareas, pero he decidido utilizar Make. Buscando comparaciones entre Make y Task, no he encontrado una clara ventaja para elegir uno sobre otro, la elección se debe principalmente a que Task lleva pocos años en desarrollo, lo que hace que las actualizaciones y cambios del mismo son muy frecuentes, y la documentación de los errores es bastantes escasa.  Make permite la utilización de todas las funcionalidades del *task runner* implícito de GO y muchas más, siempre que se implementen en el Makefile, y recibe soporte de forma regular, y gracias a su gran popularidad y la gran cantidad de años que lleva en desarrollo, si bien recibe soporte, las actualizaciones no hacen modificaciones mayores.

En los requisitos del objetivo 3 nos pide que incluyamos una clave fichero en el iv.yaml que apuntará al archivo que se usará para ejecutar las tareas. En este caso el fichero va a ser el Makefile

## Elección Gestor de Dependencias
La gestión de dependencias en GO se hace por medio de un sistema descentralizado. El usuario simplemente tiene que localizar un paquete, para lo que GO provee un [motor de búsqueda](pkg.go.dev). Y una vez localizado el paquete simpelmente se debe realizar la importación del paquete.

En el archivo go.mod se mantiene una lista de las dependencias actuales del módulo. Cuando se añade una dependencia se crea un go.sum que contiene los checksums de los módulos en los que se depende, y que son utilizados para verificar la integridad de los módulos descargados. Las herramientas de GO permiten comandos para el mantenimiento e instalación de dependencias.

Teniendo en cuenta todo lo anterior usaré también Make como gestor de dependencias, añadiendo las claves que sean necesarias para la gestión de las mismas.

## Órdenes Makefile
En este apartado se incluyen las distintas órdenes que acepta el Makefile del proyecto:
+ make check: comprueba la sintaxis de todos los archivos y devuelve los errores.
+ make build: compila todos los archivos del proyecto actual. Almacena los compilados en ./bin.
+ make installdeps: instala las dependencias de todos los archivos go que cuelgan del directorio raíz.
+ make clean: elimina los archivos compilados almacenados en ./bin.
+ make test: no implementado todavía. En un futuro ejecutará los tests.
+ make help: muestra las posible órdenes aceptadas por el make.
+ make: ejecuta check, build y clean.