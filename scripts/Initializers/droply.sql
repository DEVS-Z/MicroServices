CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL
);

CREATE TABLE modulos (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    status BOOLEAN DEFAULT TRUE
);

CREATE TABLE modulosRol (
    id SERIAL PRIMARY KEY,
    rol INT REFERENCES roles(id),
    modulo INT REFERENCES modulos(id)
);

CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    primerNombre VARCHAR(100),
    segundoNombre VARCHAR(100),
    primerApellido VARCHAR(100),
    segundoApellido VARCHAR(100),
    matricula VARCHAR(50),
    correo VARCHAR(100) UNIQUE NOT NULL,
    contrasena TEXT NOT NULL,
    rol INT REFERENCES roles(id)
);

CREATE TABLE tipoActividad (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL
);

CREATE TABLE estadoActividad (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL
);

CREATE TABLE actividades (
    id SERIAL PRIMARY KEY,
    titulo VARCHAR(150) NOT NULL,
    descripcion TEXT,
    fechaInicio DATE,
    fechaFin DATE,
    tipo INT REFERENCES tipoActividad(id),
    estado INT REFERENCES estadoActividad(id)
);

CREATE TABLE actividadesUsuario (
    id SERIAL PRIMARY KEY,
    usuario INT REFERENCES usuarios(id),
    actividad INT REFERENCES actividades(id)
);

CREATE TABLE tipoBitacora (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL
);

CREATE TABLE bitacoras (
    id SERIAL PRIMARY KEY,
    titulo VARCHAR(150),
    descripcion TEXT,
    creado TIMESTAMP,
    actualizado TIMESTAMP,
    tipo INT REFERENCES tipoBitacora(id),
    usuario INT REFERENCES usuarios(id)
);

CREATE TABLE industria (
    id SERIAL PRIMARY KEY,
    nombreOficial VARCHAR(100) NOT NULL,
    alias VARCHAR(100),
    status BOOLEAN,
    admin INT REFERENCES usuarios(id)
);

CREATE TABLE areasIndustriales (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    descripcion TEXT,
    industria INT REFERENCES industria(id)
);

CREATE TABLE tipoMedidor (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    descripcion TEXT
);

CREATE TABLE fuenteLectura (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    descripcion TEXT
);

CREATE TABLE unidadMedida (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(50) NOT NULL
);

CREATE TABLE medidores (
    id SERIAL PRIMARY KEY,
    codigo VARCHAR(100),
    enlace VARCHAR(100),
    tipo INT REFERENCES tipoMedidor(id),
--     fuenteLectura INT REFERENCES fuenteLectura(id),
    unidadMedida INT REFERENCES unidadMedida(id)
--     lecturas INT REFERENCES lecturas(id)
);

CREATE TABLE lecturas (
    id SERIAL PRIMARY KEY,
    valor DECIMAL(10, 2) NOT NULL,
    fechaRegistro TIMESTAMP NOT NULL,
    fuenteLectura INT REFERENCES fuenteLectura(id),
--     medidor INT REFERENCES medidores(id)
    medidor INT REFERENCES medidores(id)
);

CREATE TABLE lineasAgua (
    id SERIAL PRIMARY KEY,
    codigo VARCHAR(50),
    area INT REFERENCES areasIndustriales(id),
    medidor INT REFERENCES medidores(id)
);

CREATE TABLE metas (
    id SERIAL PRIMARY KEY,
    titulo VARCHAR(150),
    descripcion TEXT,
    industria INT REFERENCES industria(id)
);

CREATE TABLE tipoArchivo (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100)
);

CREATE TABLE archivos (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100),
    url TEXT,
    creado TIMESTAMP,
    tipo INT REFERENCES tipoArchivo(id)
);

CREATE TABLE reportes (
    id SERIAL PRIMARY KEY,
    titulo VARCHAR(150),
    descripcion TEXT,
    creado TIMESTAMP,
    industria INT REFERENCES industria(id),
    usuario INT REFERENCES usuarios(id),
    archivo INT REFERENCES archivos(id)
);

CREATE TABLE tipoAlerta (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100)
);

CREATE TABLE alertas (
    id SERIAL PRIMARY KEY,
    descripcion TEXT,
    tipo INT REFERENCES tipoAlerta(id)
);

CREATE TABLE alertaMedidores (
    id SERIAL PRIMARY KEY,
    medidor INT REFERENCES medidores(id),
    alerta INT REFERENCES alertas(id),
    registro TIMESTAMP
);