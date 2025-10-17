-- Crear schema si no existe
CREATE SCHEMA IF NOT EXISTS zfut;

-- =========================
-- TABLAS BASE
-- =========================
CREATE TABLE IF NOT EXISTS zfut.roles (
  rol_id BIGSERIAL PRIMARY KEY,
  nombre VARCHAR(80) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS zfut.usuarios (
  user_id BIGSERIAL PRIMARY KEY,
  nombre VARCHAR(120) NOT NULL,
  correo VARCHAR(160) NOT NULL UNIQUE,
  password TEXT NOT NULL,
  fecha_registro TIMESTAMP NOT NULL DEFAULT NOW(),
  estado VARCHAR(20) NOT NULL DEFAULT 'activo',
  rol_id BIGINT REFERENCES zfut.roles(rol_id)
);

-- =========================
-- CLUBES, SUSCRIPCIONES Y PAGOS
-- =========================
CREATE TABLE IF NOT EXISTS zfut.clubs (
  club_id BIGSERIAL PRIMARY KEY,
  nombre VARCHAR(120) NOT NULL,
  pais VARCHAR(80),
  fecha_registro TIMESTAMP NOT NULL DEFAULT NOW(),
  owner_id BIGINT REFERENCES zfut.usuarios(user_id)
);

CREATE TABLE IF NOT EXISTS zfut.suscripciones (
  suscripcion_id BIGSERIAL PRIMARY KEY,
  nombre VARCHAR(100) NOT NULL UNIQUE,
  precio NUMERIC(12,2) NOT NULL DEFAULT 0,
  cant_players INT NOT NULL
);

CREATE TABLE IF NOT EXISTS zfut.club_suscripcion (
  club_subs_id BIGSERIAL PRIMARY KEY,
  club_id BIGINT NOT NULL REFERENCES zfut.clubs(club_id),
  suscripcion_id BIGINT NOT NULL REFERENCES zfut.suscripciones(suscripcion_id),
  owner_id BIGINT REFERENCES zfut.usuarios(user_id),
  fecha_inicio DATE NOT NULL,
  fecha_fin DATE,
  status VARCHAR(30) NOT NULL DEFAULT 'activa',
  jugadores_act INT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS zfut.pagos (
  pago_id BIGSERIAL PRIMARY KEY,
  club_subs_id BIGINT NOT NULL REFERENCES zfut.club_suscripcion(club_subs_id),
  total_pagado NUMERIC(12,2) NOT NULL,
  fecha_pago TIMESTAMP NOT NULL DEFAULT NOW()
);

-- =========================
-- MIEMBROS / JUGADORES
-- =========================
CREATE TABLE IF NOT EXISTS zfut.miembros (
  miembro_id BIGSERIAL PRIMARY KEY,
  posicion VARCHAR(40),
  weareable_id VARCHAR(80),
  fecha_registro TIMESTAMP NOT NULL DEFAULT NOW(),
  altura NUMERIC(5,2),
  peso NUMERIC(5,2),
  user_id BIGINT REFERENCES zfut.usuarios(user_id)
);

-- =========================
-- EQUIPOS Y EVENTOS
-- =========================
CREATE TABLE IF NOT EXISTS zfut.equipos (
  equipo_id BIGSERIAL PRIMARY KEY,
  club_id BIGINT NOT NULL REFERENCES zfut.clubs(club_id),
  nombre VARCHAR(120) NOT NULL,
  categoria VARCHAR(80),
  fecha_registro TIMESTAMP NOT NULL DEFAULT NOW(),
  CONSTRAINT uq_club_nombre UNIQUE (club_id, nombre)
);

CREATE TABLE IF NOT EXISTS zfut.eventos (
  evento_id BIGSERIAL PRIMARY KEY,
  equipo_id BIGINT NOT NULL REFERENCES zfut.equipos(equipo_id),
  creado_por_id BIGINT REFERENCES zfut.miembros(miembro_id),
  titulo_evento VARCHAR(160) NOT NULL,
  fecha_inicio TIMESTAMP NOT NULL,
  fecha_fin TIMESTAMP,
  tipo VARCHAR(60),
  comentarios TEXT,
  lugar VARCHAR(160)
);

-- =========================
-- RUTINAS Y EJERCICIOS
-- =========================
CREATE TABLE IF NOT EXISTS zfut.rutinas (
  rutina_id BIGSERIAL PRIMARY KEY,
  creado_por_id BIGINT REFERENCES zfut.miembros(miembro_id),
  equipo_id BIGINT REFERENCES zfut.equipos(equipo_id),
  nombre VARCHAR(120) NOT NULL,
  objetivo TEXT,
  nivel_dificultad VARCHAR(40),
  fecha_registro TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS zfut.ejercicios (
  ejercicio_id BIGSERIAL PRIMARY KEY,
  rutina_id BIGINT REFERENCES zfut.rutinas(rutina_id),
  nombre VARCHAR(120) NOT NULL,
  series INT,
  repeticiones INT,
  duracion_segs INT,
  intensidad VARCHAR(40)
);

CREATE TABLE IF NOT EXISTS zfut.asignacion_rutinas (
  asignacion_id BIGSERIAL PRIMARY KEY,
  rutina_id BIGINT REFERENCES zfut.rutinas(rutina_id),
  jugador_id BIGINT REFERENCES zfut.miembros(miembro_id),
  equipo_id BIGINT REFERENCES zfut.equipos(equipo_id),
  fecha_inicio DATE NOT NULL,
  fecha_fin DATE,
  frecuencia VARCHAR(80),
  comentarios TEXT,
  rol_id BIGINT REFERENCES zfut.roles(rol_id)
);

-- =========================
-- ACTIVIDADES Y MÉTRICAS
-- =========================
CREATE TABLE IF NOT EXISTS zfut.actividades (
  actividad_id BIGSERIAL PRIMARY KEY,
  jugador_id BIGINT REFERENCES zfut.miembros(miembro_id),
  miembro_id BIGINT REFERENCES zfut.miembros(miembro_id),
  tipo VARCHAR(80) NOT NULL,
  fecha_inicio TIMESTAMP NOT NULL,
  fecha_fin TIMESTAMP,
  comentarios TEXT
);

CREATE TABLE IF NOT EXISTS zfut.signos_vitales (
  signos_id BIGSERIAL PRIMARY KEY,
  actividad_id BIGINT REFERENCES zfut.actividades(actividad_id),
  fecha_registro TIMESTAMP NOT NULL DEFAULT NOW(),
  metrica VARCHAR(80),
  valor NUMERIC(12,4),
  unidad_medida VARCHAR(40),
  miembro_id BIGINT REFERENCES zfut.miembros(miembro_id)
);

CREATE TABLE IF NOT EXISTS zfut.alertas (
  alertas_id BIGSERIAL PRIMARY KEY,
  actividad_id BIGINT REFERENCES zfut.actividades(actividad_id),
  fecha_registro TIMESTAMP NOT NULL DEFAULT NOW(),
  tipo VARCHAR(80),
  descripcion TEXT,
  gravedad VARCHAR(40),
  atendido_si_no BOOLEAN DEFAULT FALSE,
  miembro_id BIGINT REFERENCES zfut.miembros(miembro_id)
);

CREATE TABLE IF NOT EXISTS zfut.reportes (
  reporte_id BIGSERIAL PRIMARY KEY,
  jugador_id BIGINT REFERENCES zfut.miembros(miembro_id),
  fecha_registro TIMESTAMP NOT NULL DEFAULT NOW(),
  tipo VARCHAR(80),
  comentarios TEXT
);

-- =========================
-- DATOS DE PRUEBA
-- =========================
INSERT INTO zfut.roles (nombre)
VALUES ('owner'), ('coach'), ('analyst'), ('assistant'), ('athlete')
ON CONFLICT DO NOTHING;

INSERT INTO zfut.usuarios (nombre, correo, password, estado, rol_id)
VALUES 
('Carla López', 'owner@zfut.test', 'hashed_pw', 'activo', 1),
('Diego Ramírez', 'coach@zfut.test', 'hashed_pw', 'activo', 2),
('Ana Pérez', 'analyst@zfut.test', 'hashed_pw', 'activo', 3),
('Luis Mora', 'assistant@zfut.test', 'hashed_pw', 'activo', 4),
('Jugador Uno', 'p1@zfut.test', 'hashed_pw', 'activo', 5),
('Jugador Dos', 'p2@zfut.test', 'hashed_pw', 'activo', 5)
ON CONFLICT DO NOTHING;

INSERT INTO zfut.clubs (nombre, pais, owner_id)
VALUES ('ZFut FC', 'México', 1)
ON CONFLICT DO NOTHING;

INSERT INTO zfut.suscripciones (nombre, precio, cant_players)
VALUES ('Basic', 0.00, 25), ('Pro', 49.99, 100)
ON CONFLICT DO NOTHING;

INSERT INTO zfut.club_suscripcion (club_id, suscripcion_id, owner_id, fecha_inicio, status, jugadores_act)
VALUES (1, 2, 1, NOW(), 'activa', 6)
ON CONFLICT DO NOTHING;
