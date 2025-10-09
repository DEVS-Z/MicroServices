IF NOT EXISTS (SELECT * FROM sys.schemas WHERE name = 'zfut')
    EXEC('CREATE SCHEMA zfut;');
GO

-- =========================
-- TABLAS BASE
-- =========================
CREATE TABLE zfut.roles (
  rol_id BIGINT IDENTITY(1,1) PRIMARY KEY,
  nombre NVARCHAR(80) NOT NULL UNIQUE
);

CREATE TABLE zfut.usuarios (
  user_id BIGINT IDENTITY(1,1) PRIMARY KEY,
  nombre NVARCHAR(120) NOT NULL,
  correo NVARCHAR(160) NOT NULL UNIQUE,
  password NVARCHAR(MAX) NOT NULL,
  fecha_registro DATETIME2 NOT NULL DEFAULT SYSDATETIME(),
  estado NVARCHAR(20) NOT NULL DEFAULT 'activo',
  rol_id BIGINT NULL REFERENCES zfut.roles(rol_id)
);

-- =========================
-- CLUBES, SUSCRIPCIONES Y PAGOS
-- =========================
CREATE TABLE zfut.clubs (
  club_id BIGINT IDENTITY(1,1) PRIMARY KEY,
  nombre NVARCHAR(120) NOT NULL,
  pais NVARCHAR(80),
  fecha_registro DATETIME2 NOT NULL DEFAULT SYSDATETIME(),
  owner_id BIGINT NULL REFERENCES zfut.usuarios(user_id)
);

CREATE TABLE zfut.suscripciones (
  suscripcion_id BIGINT IDENTITY(1,1) PRIMARY KEY,
  nombre NVARCHAR(100) NOT NULL UNIQUE,
  precio DECIMAL(12,2) NOT NULL DEFAULT 0,
  cant_players INT NOT NULL
);

CREATE TABLE zfut.club_suscripcion (
  club_subs_id BIGINT IDENTITY(1,1) PRIMARY KEY,
  club_id BIGINT NOT NULL REFERENCES zfut.clubs(club_id),
  suscripcion_id BIGINT NOT NULL REFERENCES zfut.suscripciones(suscripcion_id),
  owner_id BIGINT REFERENCES zfut.usuarios(user_id),
  fecha_inicio DATE NOT NULL,
  fecha_fin DATE,
  status NVARCHAR(30) NOT NULL DEFAULT 'activa',
  jugadores_act INT NOT NULL DEFAULT 0
);

CREATE TABLE zfut.pagos (
  pago_id BIGINT IDENTITY(1,1) PRIMARY KEY,
  club_subs_id BIGINT NOT NULL REFERENCES zfut.club_suscripcion(club_subs_id),
  total_pagado DECIMAL(12,2) NOT NULL,
  fecha_pago DATETIME2 NOT NULL DEFAULT SYSDATETIME()
);

-- =========================
-- MIEMBROS / JUGADORES
-- =========================
CREATE TABLE zfut.miembros (
  miembro_id BIGINT IDENTITY(1,1) PRIMARY KEY,
  posicion NVARCHAR(40),
  weareable_id NVARCHAR(80),
  fecha_registro DATETIME2 NOT NULL DEFAULT SYSDATETIME(),
  altura DECIMAL(5,2),
  peso DECIMAL(5,2),
  user_id BIGINT REFERENCES zfut.usuarios(user_id)
);

-- =========================
-- EQUIPOS Y EVENTOS
-- =========================
CREATE TABLE zfut.equipos (
  equipo_id BIGINT IDENTITY(1,1) PRIMARY KEY,
  club_id BIGINT NOT NULL REFERENCES zfut.clubs(club_id),
  nombre NVARCHAR(120) NOT NULL,
  categoria NVARCHAR(80),
  fecha_registro DATETIME2 NOT NULL DEFAULT SYSDATETIME(),
  CONSTRAINT UQ_club_nombre UNIQUE (club_id, nombre)
);

CREATE TABLE zfut.eventos (
  evento_id BIGINT IDENTITY(1,1) PRIMARY KEY,
  equipo_id BIGINT NOT NULL REFERENCES zfut.equipos(equipo_id),
  creado_por_id BIGINT REFERENCES zfut.miembros(miembro_id),
  titulo_evento NVARCHAR(160) NOT NULL,
  fecha_inicio DATETIME2 NOT NULL,
  fecha_fin DATETIME2,
  tipo NVARCHAR(60),
  comentarios NVARCHAR(MAX),
  lugar NVARCHAR(160)
);

-- =========================
-- RUTINAS Y EJERCICIOS
-- =========================
CREATE TABLE zfut.rutinas (
  rutina_id BIGINT IDENTITY(1,1) PRIMARY KEY,
  creado_por_id BIGINT REFERENCES zfut.miembros(miembro_id),
  equipo_id BIGINT REFERENCES zfut.equipos(equipo_id),
  nombre NVARCHAR(120) NOT NULL,
  objetivo NVARCHAR(MAX),
  nivel_dificultad NVARCHAR(40),
  fecha_registro DATETIME2 NOT NULL DEFAULT SYSDATETIME()
);

CREATE TABLE zfut.ejercicios (
  ejercicio_id BIGINT IDENTITY(1,1) PRIMARY KEY,
  rutina_id BIGINT REFERENCES zfut.rutinas(rutina_id),
  nombre NVARCHAR(120) NOT NULL,
  series INT,
  repeticiones INT,
  duracion_segs INT,
  intensidad NVARCHAR(40)
);

CREATE TABLE zfut.asignacion_rutinas (
  asignacion_id BIGINT IDENTITY(1,1) PRIMARY KEY,
  rutina_id BIGINT REFERENCES zfut.rutinas(rutina_id),
  jugador_id BIGINT REFERENCES zfut.miembros(miembro_id),
  equipo_id BIGINT REFERENCES zfut.equipos(equipo_id),
  fecha_inicio DATE NOT NULL,
  fecha_fin DATE,
  frecuencia NVARCHAR(80),
  comentarios NVARCHAR(MAX),
  rol_id BIGINT REFERENCES zfut.roles(rol_id)
);

-- =========================
-- ACTIVIDADES Y MÉTRICAS
-- =========================
CREATE TABLE zfut.actividades (
  actividad_id BIGINT IDENTITY(1,1) PRIMARY KEY,
  jugador_id BIGINT REFERENCES zfut.miembros(miembro_id),
  miembro_id BIGINT REFERENCES zfut.miembros(miembro_id),
  tipo NVARCHAR(80) NOT NULL,
  fecha_inicio DATETIME2 NOT NULL,
  fecha_fin DATETIME2,
  comentarios NVARCHAR(MAX)
);

CREATE TABLE zfut.signos_vitales (
  signos_id BIGINT IDENTITY(1,1) PRIMARY KEY,
  actividad_id BIGINT REFERENCES zfut.actividades(actividad_id),
  fecha_registro DATETIME2 NOT NULL DEFAULT SYSDATETIME(),
  metrica NVARCHAR(80),
  valor DECIMAL(12,4),
  unidad_medida NVARCHAR(40),
  miembro_id BIGINT REFERENCES zfut.miembros(miembro_id)
);

CREATE TABLE zfut.alertas (
  alertas_id BIGINT IDENTITY(1,1) PRIMARY KEY,
  actividad_id BIGINT REFERENCES zfut.actividades(actividad_id),
  fecha_registro DATETIME2 NOT NULL DEFAULT SYSDATETIME(),
  tipo NVARCHAR(80),
  descripcion NVARCHAR(MAX),
  gravedad NVARCHAR(40),
  atendido_si_no BIT DEFAULT 0,
  miembro_id BIGINT REFERENCES zfut.miembros(miembro_id)
);

CREATE TABLE zfut.reportes (
  reporte_id BIGINT IDENTITY(1,1) PRIMARY KEY,
  jugador_id BIGINT REFERENCES zfut.miembros(miembro_id),
  fecha_registro DATETIME2 NOT NULL DEFAULT SYSDATETIME(),
  tipo NVARCHAR(80),
  comentarios NVARCHAR(MAX)
);

-- =========================
-- DATOS DE PRUEBA
-- =========================
INSERT INTO zfut.roles (nombre)
VALUES ('owner'), ('coach'), ('analyst'), ('assistant'), ('athlete');

INSERT INTO zfut.usuarios (nombre, correo, password, estado, rol_id)
VALUES 
('Carla López', 'owner@zfut.test', 'hashed_pw', 'activo', 1),
('Diego Ramírez', 'coach@zfut.test', 'hashed_pw', 'activo', 2),
('Ana Pérez', 'analyst@zfut.test', 'hashed_pw', 'activo', 3),
('Luis Mora', 'assistant@zfut.test', 'hashed_pw', 'activo', 4),
('Jugador Uno', 'p1@zfut.test', 'hashed_pw', 'activo', 5),
('Jugador Dos', 'p2@zfut.test', 'hashed_pw', 'activo', 5);

INSERT INTO zfut.clubs (nombre, pais, owner_id)
VALUES ('ZFut FC', 'México', 1);

INSERT INTO zfut.suscripciones (nombre, precio, cant_players)
VALUES ('Basic', 0.00, 25), ('Pro', 49.99, 100);

INSERT INTO zfut.club_suscripcion (club_id, suscripcion_id, owner_id, fecha_inicio, status, jugadores_act)
VALUES (1, 2, 1, GETDATE(), 'activa', 6);
GO
