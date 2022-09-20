# ChangeFormat

## Problema a resolver
Es común que a la hora de entregar documentación y formularios, la entrega se vea rechazada debido a que la plataforma donde se suben los archivos no admite alguno de los formatos subidos. 

## Solución al problema
La solución propuesta es crear un servicio que se encargue de transformar los formatos no reconocidos a formatos aceptados por la plataforma de entrega, de forma que el usuario no se tenga que preocupar por dicha transformación.

## ¿Por qué en la nube?
La idea es que el usuario suba el archivo y el formato de este se adapte según la necesidad del usuario.

## Lógica de negocio
La lógica de negocio consiste en subir el archivo, detectar cual son los formatos aceptados, y si alguno de los archivos no coincide con los mismos se realizar la transformación del mismo a los aceptados. En caso de que se suba un archivo cuyo formato no es contemplado (por ejemplo, se sube una imagen y en la entrega no se contempla la entrega de imágenes) se comunicará al usuario la imposibilidad.
