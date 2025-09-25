DROP SCHEMA public CASCADE;
CREATE SCHEMA public;

CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL
);

CREATE TABLE modulos (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    status BOOLEAN DEFAULT TRUE,
    icon VARCHAR(50) NULL
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
    matricula VARCHAR(50) UNIQUE NOT NULL,
    correo VARCHAR(100) UNIQUE NOT NULL,
    contrasena TEXT NOT NULL,
    rol INT REFERENCES roles(id)
);

CREATE TABLE estadoNotificacion (
    id SERIAL PRIMARY KEY,
    descripcion VARCHAR(100) NOT NULL
);

-- Tabla: mensajes
CREATE TABLE mensajes (
    id SERIAL PRIMARY KEY,
    fecha TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    usuario INT NOT NULL REFERENCES usuarios(id),
    concepto VARCHAR(100) NOT NULL,
    descripcion TEXT NOT NULL,
    status VARCHAR(50) NOT NULL
);

-- Tabla: notificaciones
CREATE TABLE notificaciones (
    id SERIAL PRIMARY KEY,
    mensaje INT NOT NULL REFERENCES mensajes(id) ON DELETE CASCADE,
    usuario INT NOT NULL REFERENCES usuarios(id),
    estado INT NOT NULL REFERENCES estadoNotificacion(id)
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
    nombre VARCHAR(100),
    extension VARCHAR(10) NOT NULL
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
    archivo INT REFERENCES archivos(id) NULL
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


-- ===== ROLES =====
INSERT INTO roles (nombre) VALUES
('Administrador'),
('Gerente'),
('Auditor');

-- ===== PERMISOS =====
INSERT INTO modulos (nombre, status) VALUES
('Roles', true),
('Modulos', true),
('Modulos de Rol', true),
('Usuarios', true),
('Tipos de Actividades', true),
('Estado de las Actividades', true),
('Actividades', true),
('Actividades por Usuario', true),
('Tipos de Bitacora', true),
('Bitacoras', true),
('Industrias', true),
('Areas Industriales', true),
('Tipos de Medidores', true),
('Fuentes de Lecturas', true),
('Unidades de Medidas', true),
('Lecturas', true),
('Medidores', true),
('Lineas de Agua', true),
('Metas', true),
('Tipos de Archivos', true),
('Archivos', true),
('Reportes', true),
('Tipo de Alertas', true),
('Alertas', true),
('Alertas de los Medidores', true),
('Mensajes', true),
('Notificaciones', true);

INSERT INTO modulosRol (rol, modulo) VALUES
(1, 1),
(1, 2),
(1, 3),
(1, 4),
(1, 5),
(1, 6),
(1, 7),
(1, 8),
(1, 9),
(1, 10),
(1, 11),
(1, 12),
(1, 13),
(1, 14),
(1, 15),
(1, 16),
(1, 17),
(1, 18),
(1, 19),
(1, 20),
(1, 21),
(1, 22),
(1, 23),
(1, 24),
(1, 25),
(1, 26),
(1, 27),
(2, 1),
(2, 2),
(2, 3),
(2, 4),
(2, 5),
(2, 6),
(2, 7),
(2, 8),
(2, 9),
(2, 10),
(2, 11),
(2, 12),
(2, 13),
(2, 14),
(2, 15),
(2, 16),
(2, 17),
(2, 18),
(2, 19),
(2, 20),
(2, 21),
(2, 22),
(2, 23),
(2, 24),
(2, 25),
(2, 26),
(2, 27),
(3, 1),
(3, 2),
(3, 3),
(3, 4),
(3, 5),
(3, 6),
(3, 7),
(3, 8),
(3, 9),
(3, 10),
(3, 11),
(3, 12),
(3, 13),
(3, 14),
(3, 15),
(3, 16),
(3, 17),
(3, 18),
(3, 19),
(3, 20),
(3, 21),
(3, 22),
(3, 23),
(3, 24),
(3, 25),
(3, 26),
(3, 27);

-- ===== USUARIOS =====

INSERT INTO usuarios (primerNombre, segundoNombre, primerApellido, segundoApellido, matricula, correo, contrasena, rol)
VALUES
('Andrea', 'Guadalupe', 'Quintana', 'Zepeda', 'A003', '0322103793@ut-tijuana.edu.mx', 'passAndrea', 1),
('Juan', 'Antonio', 'Avalos', 'Garcia', 'A004', '0322103675@ut-tijuana.edu.mx', 'passJuan', 1),
('Jonathan', NULL, 'Martinez', 'Zavala', 'A005', '0322103758@ut-tijuana.edu.mx', 'passJonathan', 1),
('Jose', 'de Jesus', 'Ponce', 'Duarte', 'A006', '0322103790@ut-tijuana.edu.mx', 'passJose', 1),
('Miguel', 'Isaac', 'Garcia', 'Lopez', 'A007', '0322103717@ut-tijuana.edu.mx', 'passMiguel', 1),
('Cesia', 'Nuemi', 'Ochoa', 'Huerta', 'A008', '0322104047@ut-tijuana.edu.mx', 'passCesia', 1);

INSERT INTO tipoActividad (nombre) VALUES ('Inspección');
INSERT INTO tipoActividad (nombre) VALUES ('Mantenimiento');

-- ===== ESTADO ACTIVIDAD =====
INSERT INTO estadoActividad (nombre) VALUES ('En Progreso');
INSERT INTO estadoActividad (nombre) VALUES ('Finalizado');

INSERT INTO estadoActividad (nombre) VALUES ('En Pausa');
INSERT INTO estadoActividad (nombre) VALUES ('Archivado');

INSERT INTO tipoBitacora (nombre) VALUES ('Bitácora de Producción');
INSERT INTO tipoBitacora (nombre) VALUES ('Bitácora de Mantenimiento');

-- ===== TIPO MEDIDOR =====
INSERT INTO tipoMedidor (nombre, descripcion) VALUES ('Flujo', 'Medidor de flujo de agua');
INSERT INTO tipoMedidor (nombre, descripcion) VALUES ('Temperatura', 'Medidor de temperatura del flujo de agua');
INSERT INTO tipoMedidor (nombre, descripcion) VALUES ('Presión', 'Medidor de presión hidráulica');

-- ===== FUENTE LECTURA =====
INSERT INTO fuenteLectura (nombre, descripcion) VALUES ('Manual', 'Registro por operador');
INSERT INTO fuenteLectura (nombre, descripcion) VALUES ('Automática', 'Sensor digital');

-- ===== UNIDAD MEDIDA =====
INSERT INTO unidadMedida (nombre) VALUES ('Litros');
INSERT INTO unidadMedida (nombre) VALUES ('°C');
INSERT INTO unidadMedida (nombre) VALUES ('PSI');

-- ===== TIPO ARCHIVO =====
INSERT INTO tipoArchivo (nombre, extension) VALUES ('PDF', '.pdf');
INSERT INTO tipoArchivo (nombre, extension) VALUES ('CSV', '.csv');
INSERT INTO tipoArchivo (nombre, extension) VALUES ('XLSX', '.xlsx');
INSERT INTO tipoArchivo (nombre, extension) VALUES ('XLS', '.xls');

INSERT INTO tipomedidor (id, nombre, descripcion) VALUES
(4, 'Caudal', 'Medidor de caudal de agua'),
(5, 'Turbidez', 'Medidor de turbidez del agua'),
(6, 'Conductividad', 'Medidor de conductividad eléctrica'),
(7, 'pH', 'Medidor de pH del agua'),
(8, 'Nivel', 'Medidor de nivel de agua'),
(9, 'Caudal Secundario', 'Medidor adicional de caudal')
ON CONFLICT (id) DO NOTHING;

-- ========================================
-- 2. Inserción de UNIDADES DE MEDIDA (unidadmedida)
-- ========================================
INSERT INTO unidadmedida (id, nombre) VALUES
(4, 'Conductividad'),
(5, 'pH'),
(6, 'Nivel'),
(7, 'Caudal'),
(8, 'Turbidez')
ON CONFLICT (id) DO NOTHING;

-- ===== TIPO ALERTA =====
INSERT INTO tipoAlerta (nombre) VALUES
('Sobrepresión'),
('Sobretemperatura'),
('Fuga detectada'),
('Turbidez alta'),
('Conductividad elevada');