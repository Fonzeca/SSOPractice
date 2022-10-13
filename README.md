# SSOPractice

### Objetivo
El desarrollo de un micro que integre un SSO contra cuentas de gmail y facebook, capture los datos básicos del user, solicite autorización y mantenga una sesión abierta. Deberá de realizar desconexión automática pasados 15 min sin actividad el usuario o si inicia sesión en otro browser o dispositivo.

### Proyecto
Lo que se hizo fue un microsevicio en golang, y un sitio web en Angular 13 para consumir este.

### Configuracion
En la carpeta "configs" estan los archivos de configuracion.
 - **config.json**:  Configuracion del microservicio. 
 - **environment.prod.ts**: configuracion del sitio web, solo se usa para setear la url del microservicio.
 - **nginx.conf y mime.types**: son archivos config de nginx, por si quieren hacerle deploy con docker.

Hay ejemplos de las archivos de configuracion.

### Deploy
Les deje un docker-compose-yaml para que puedan hacerle deploy facilmente.


### Servidor live
Les dejo una url donde esta montado el microservicio:
http://vps-1791261-x.dattaweb.com:2563/

NOTA: En el servidor solo se puede iniciar sesion con Google, ya que Facebook solo admite callbacks con https.