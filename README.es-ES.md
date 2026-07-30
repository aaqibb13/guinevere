

![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/aaqibb13/guinevere) ![GitHub last commit](https://img.shields.io/github/last-commit/aaqibb13/guinevere)
# AranGO-base-setup
Este repositorio contiene la configuración de la base de datos para ArangoDB para tu proyecto en Go, junto con base-crud, creación de Colecciones, Grafos (próximamente), Analizadores Personalizados, Vistas, Ejecutor de Consultas y Transacciones de BD.

## Dependencias Directas (Minimalistas)
1. **go-driver**: Se utiliza go-driver para proporcionar la funcionalidad que pueda necesitar un usuario.
2. **viper**: se utiliza para leer la configuración desde variables de entorno. 
3. **logrus**: se utiliza para el registro de logs, lo que te permite aprovechar la funcionalidad extendida del paquete logrus en caso de que desees establecer niveles de formato o hacer tus registros más verbosos. 
   
    **Nota:** logrus se puede omitir si lo consideras necesario.

   
## Funcionalidades incluidas
Las funcionalidades incorporadas por defecto son:
- [x] Creación de bases de datos y usuarios correspondientes
- [x] Creación de **Colecciones**, **Analizador Personalizado** y **Vistas** al inicio (preferiblemente)
- [x] Ejecutor de consultas para la ejecución de consultas AQL sin procesar
- [x] QueryExecutor con conteo completo del cursor al ejecutar consultas AQL sin procesar
- [x] Inicialización de una **Transacción** de BD

# TODO
* Pospuesto
- [ ] Agregar pruebas para cada función (pruebas unitarias)
- [ ] Agregar creación de Grafos 
- [ ] Agregar diferentes tipos de analizadores (como norm, ngram, pipeline, etc.)
- [ ] Incluir funcionalidad de edición de Vistas (ArangoDB no soporta la edición de vistas de forma predeterminada)
- [ ] Mejorar la documentación y la estructura del código
- [ ] Agregar scripts para respaldar datos

## Cómo usar:
- Clona el repositorio en la ubicación que prefieras:

          git clone https://github.com/aaqibb13/guinevere.git
      
- Ejecuta:

          go mod tidy
- Ya estás listo para usar la configuración base.

## Ejecutar arangodb localmente
- Asegúrate de tener Docker instalado en tu sistema (dependiendo de tu máquina: `Linux`, `Mac` o `Windows`)
- Dirígete a `guinevere/deployments` y ejecuta:
  
        docker-compose -f arango-docker-compose.yml up --build -d db
- Deberías poder ver el servidor ArangoDB ejecutándose en `127.0.0.1:8529`, al cual puedes acceder utilizando las credenciales especificadas en el archivo `arango-docker-compose.yml`.
