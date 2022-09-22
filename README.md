# RouteCheck

## Descripción del problema
A la hora de realizar ciertos viajes en automóvil debemos pasar por un peaje o varios, y con el precio actual de la gasolina no sabemos si es realmente rentable pagar dicho peaje, o ahorrarnos el precio de ese peaje y realizar unos cuantos kilómetros extra. A parte de esto, tampoco sabemos si la diferencia en precio va a justificar el ahorro de tiempo que nos puede conllevar ir por peaje.

## Solución
Teniendo en cuenta nuestro automóvil y el precio actual de la gasolina, conocer de antemano cuánto dinero aproximadamente nos va a costar el ir por peaje, o evitar el peaje. En el caso de que nuestra ruta pueda atravesar varios peajes probar las distintas combinaciones que haya y ver cual es la más rentable. Teniendo en cuenta el precio calculado y el tiempo empleado en cada una de las rutas seleccionar una que sea la más favorable en relación precio/tiempo.

## Configuración git
Para acceder a la configuración pinchar este [enlace](https://github.com/dmonjasm/IV/blob/Objetivo-0/docs/config.md)


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
