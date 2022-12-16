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
+ Prestaciones. En particular, el tiempo que toma la ejecución de las tareas
+ Dificultad de la sintaxis.

Voy a utilizar la información del repositorio [Awesome Go](https://github.com/avelino/awesome-go).

Si nos vamos al apartado de dicho repositorio de Build Automation voy a destacar Task y make:

+ Una opción sería [Task](https://github.com/go-task/task), surge como una alternativa más simple que Make. Task ejecuta tareas registradas en un Taskfile con formato .yml o .yaml. La sintaxis de los Taskfile es bastante sencilla, si bien requiere más líneas que los Makefile.

    + Task recibe mantenimiento asiduamente, lo cual se puede apreciar en la página de github.

    + Task es más rápido en la ejecución de tareas que Make, para ello se han ejecutado tareas de compilación e instalación de dependencias y se han demostrado que los tiempos obtenidos por Task han sido mejores que los obtenidos por Make. Se adjunta una tabla con los tiempos obtenidos por cada gestor de tareas para distintas tareas.

        | Tiempos       | Task  | Make  |
        |---------------|-------|-------|
        | default       |  0.028s     | 0.078s      |  
        | build         |  0.044s     | 0.051s      |
        | check         |  0.021s     | 0.039s      |
        | installdeps   |  0.038s     | 0.057s      |    
        | test          |  0.015s     | 0.034s      |
        | clean         |  0.017s     | 0.027s      |

    + La sintaxis de los Taskfile es más simple que la necesitada por los Makefile, en particular, para proyectos grandes. Requiere más líneas que un Makefile y la indentación puede ser problemática, pero permite generar tareas fáciles de entender para el programador.

+ La última opción considerada es [Make](https://www.gnu.org/software/make/). Es una herramienta que se ha utilizado ampliamente durante los años, pero su sintaxis y los errores derivados de esta hacen que sea compleja de usar para proyectos de gran tamaño.

    + Make recibe mantenimiento asiduamente, como se puede apreciar en su documentación.

    + Make es más lento en la ejecución de tareas que Task.

    + La sintaxis de los Makefile se vuelve compleja de entender cuando se tratan de proyectos grandes.


Hay otros *task runners*, aunque estos llevan bastante tiempo (meses o incluso años) sin recibir soporte, luego los he descartado automáticamente.

Task y Make son dos buenos candidatos para usar como gestor de tareas, pero he decidido utilizar Task. Task y Make reciben ambos soporte asiduamente, pero Task obtiene unos mejores tiempos en la ejecución de las tareas, además de usar una sintaxis más simple que Make. Por estas razones he elegido Task.

## Elección Gestor de Dependencias
La gestión de dependencias en GO se hace por medio la herramienta de línea de órdenes de GO. Es más, buscando gestores de dependencias para GO he visto que están obsoletos, y que los repositorios de los mismo han sido archivados.

Luego realmente no se requiere de un gestor de dependencias concreto. En este caso, como hemos usado Task como gestor de tareas simplemente añadimos al Taskfile las claves necesarias para la gestión de dependencias.

## Órdenes Taskfile
En este apartado se incluyen las distintas órdenes que acepta el Taskfile del proyecto:
+ go-task check: comprueba la sintaxis de todos los archivos y devuelve los errores.
+ go-task build: compila todos los archivos del proyecto actual. Almacena los compilados en ./bin.
+ go-task installdeps: instala las dependencias de todos los archivos go que cuelgan del directorio raíz.
+ go-task clean: elimina los archivos compilados almacenados en ./bin.
+ go-task test: no implementado todavía. En un futuro ejecutará los tests.
+ go-task list: muestra las posible órdenes aceptadas por Task y una descripción de que hace cada cosa.
+ go-task: ejecuta check, build y clean.

## Elección de *test runner*
Para la elacción del *test runner* voy a priorizar:
+ Que ofrezca funciones de aserción.
+ Que reciba mantenimiento con asiduidad.
+ Customización de la salida.

Entre posibles *test runners* para GO he encontrado:
+ [Testity](https://github.com/stretchr/testify): se trata de un framework para el testeo y es uno de los más populares. Testify ofrece funciones de aserción para la comparación y verificación de valores. También propociona funcionalidades de *mock* que se pueden utilizar en el código de los test. Fijándonos en su repositorio de github recibe mantenimiento de forma asidua. Respecto a la salida, no es customizable y es bastante estándar.
+ [Testing](https://pkg.go.dev/testing): se trata del paquete por defecto proporcionado por GO, permite el testeo de paquetes y utiliza la herramienta de comando en línea de go, `go test`. Por ser un paquete proporcionado por GO recibe mantenimiento con asiduidad. El problema de esta opción es que no soporta aserciones. La salida proporcionada es completamente estándar. 
+ [GoConvey](https://github.com/smartystreets/goconvey): es un framework de testing BDD que, al igual que todos los anteriores, es compatible con el comando `go test` de la herramienta de comandos en línea de go. Sin embargo, fijándonos en su repositorio de github lleva un cierto tiempo sin recibir mantenimiento (más de un año). Permite la salida por terminal o bien por interfaz web, en ambos casos la salida es similar, coloreada, detallada y fácil de leer.
+ [Ginkgo](https://github.com/onsi/ginkgo): se trata de un framefork de testeo de tipy BDD. Fijándonos en su repositorio vemos que recibe mantenimiento con asiduidad. Proporciona *container nodes* que asisten a la organización de especificaciones y a la creación de aserciones. La salida es fácil de leer y se puede obtener en varios formatos. Además se puede customizar como se recoge dicha salida.

Atendiendo a las opciones consideradas voy a elegir Ginkgo como *test runner*.