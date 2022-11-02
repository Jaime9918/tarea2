# Ejemplo

## Version local o una sola VM

### Como compilar
    1) make combine (dist013)
    2) make datanode1 (dist013)
    3) make namenode (dist014)
    4) make rebeldes (dist014)
    5) make datanode2 (dist015)
    6) make datanode3 (dist016)

### Observaciones

    -A veces hay problemas de puertos usados por la reconexión de la maquina virtual, sugiero usar "sudo lsof -i:port" y luego borrar el proceso con "kill PID"
    -Se asume que el archivo DATA.txt se entregará al comienzo de cada ejecución (ya sea vacio o con datos), en caso de que se necesite que el archivo se cree y resetee en cada ejecución sacar comentarios de la linea  312 y 313.
    -Se asume que los id ingresados son numeros (funciona bien de igual manera con palabras)
    -Se asume que los ids ingresados son numeros menores a 15 digitos.

