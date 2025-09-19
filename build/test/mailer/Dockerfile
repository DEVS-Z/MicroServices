# Imagen base oficial de Node.js
FROM node:18

# Crear y establecer directorio de trabajo
WORKDIR /app

# Copiar archivos necesarios
COPY index.js .
COPY QubeFlex.png .

# Instalar dependencias directamente
RUN npm install express nodemailer 

# Exponer el puerto de la app
EXPOSE 3000

# Comando para ejecutar la app
CMD ["node", "index.js"]
