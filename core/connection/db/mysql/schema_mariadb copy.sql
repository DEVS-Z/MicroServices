-- ========================================
-- CREAR BASE DE DATOS
-- ========================================
CREATE DATABASE IF NOT EXISTS zfut;
USE zfut;

-- ========================================
-- TABLAS BASE
-- ========================================
CREATE TABLE IF NOT EXISTS roles (
  rol_id BIGINT AUTO_INCREMENT PRIMARY KEY,
  nombre VARCHAR(80) NOT NULL UNIQUE
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS usuarios (
  user_id BIGINT AUTO_INCREMENT PRIMARY KEY,
  nombre VARCHAR(120) NOT NULL,
  correo VARCHAR(160) NOT NULL UNIQUE,
  password TEXT NOT NULL,
  fecha_registro TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  estado VARCHAR(20) NOT NULL DEFAULT 'activo',
  rol_id BIGINT,
  FOREIGN KEY (rol_id) REFERENCES roles(rol_id)
) ENGINE=InnoDB;

-- ========================================
-- CLUBES, SUSCRIPCIONES Y PAGOS
-- ========================================
CREATE TABLE IF NOT EXISTS clubs (
  club_id BIGINT AUTO_INCREMENT PRIMARY KEY,
  nombre VARCHAR(120) NOT NULL,
  pais VARCHAR(80),
  fecha_registro TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  owner_id BIGINT,
  FOREIGN KEY (owner_id) REFERENCES usuarios(user_id)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS suscripciones (
  suscripcion_id BIGINT AUTO_INCREMENT PRIMARY KEY,
  nombre VARCHAR(100) NOT NULL UNIQUE,
  precio DECIMAL(12,2) NOT NULL DEFAULT 0,
  cant_players INT NOT NULL
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS club_suscripcion (
  club_subs_id BIGINT AUTO_INCREMENT PRIMARY KEY,
  club_id BIGINT NOT NULL,
  suscripcion_id BIGINT NOT NULL,
  owner_id BIGINT,
  fecha_inicio DATE NOT NULL,
  fecha_fin DATE,
  status VARCHAR(30) NOT NULL DEFAULT 'activa',
  jugadores_act INT NOT NULL DEFAULT 0,
  FOREIGN KEY (club_id) REFERENCES clubs(club_id),
  FOREIGN KEY (suscripcion_id) REFERENCES suscripciones(suscripcion_id),
  FOREIGN KEY (owner_id) REFERENCES usuarios(user_id)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS pagos (
  pago_id BIGINT AUTO_INCREMENT PRIMARY KEY,
  club_subs_id BIGINT NOT NULL,
  total_pagado DECIMAL(12,2) NOT NULL,
  fecha_pago TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (club_subs_id) REFERENCES club_suscripcion(club_subs_id)
) ENGINE=InnoDB;

-- ========================================
-- MIEMBROS / JUGADORES
-- ========================================
CREATE TABLE IF NOT EXISTS miembros (
  miembro_id BIGINT AUTO_INCREMENT PRIMARY KEY,
  posicion VARCHAR(40),
  weareable_id VARCHAR(80),
  fecha_registro TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  altura DECIMAL(5,2),
  peso DECIMAL(5,2),
  user_id BIGINT,
  FOREIGN KEY (user_id) REFERENCES usuarios(user_id)
) ENGINE=InnoDB;

-- ========================================
-- EQUIPOS Y EVENTOS
-- ========================================
CREATE TABLE IF NOT EXISTS equipos (
  equipo_id BIGINT AUTO_INCREMENT PRIMARY KEY,
  club_id BIGINT NOT NULL,
  nombre VARCHAR(120) NOT NULL,
  categoria VARCHAR(80),
  fecha_registro TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY uq_club_nombre (club_id, nombre),
  FOREIGN KEY (club_id) REFERENCES clubs(club_id)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS eventos (
  evento_id BIGINT AUTO_INCREMENT PRIMARY KEY,
  equipo_id BIGINT NOT NULL,
  creado_por_id BIGINT,
  titulo_evento VARCHAR(160) NOT NULL,
  fecha_inicio TIMESTAMP NOT NULL,
  fecha_fin TIMESTAMP,
  tipo VARCHAR(60),
  comentarios TEXT,
  lugar VARCHAR(160),
  FOREIGN KEY (equipo_id) REFERENCES equipos(equipo_id),
  FOREIGN KEY (creado_por_id) REFERENCES miembros(miembro_id)
) ENGINE=InnoDB;

-- ========================================
-- RUTINAS Y EJERCICIOS
-- ========================================
CREATE TABLE IF NOT EXISTS rutinas (
  rutina_id BIGINT AUTO_INCREMENT PRIMARY KEY,
  creado_por_id BIGINT,
  equipo_id BIGINT,
  nombre VARCHAR(120) NOT NULL,
  objetivo TEXT,
  nivel_dificultad VARCHAR(40),
  fecha_registro TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (creado_por_id) REFERENCES miembros(miembro_id),
  FOREIGN KEY (equipo_id) REFERENCES equipos(equipo_id)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS ejercicios (
  ejercicio_id BIGINT AUTO_INCREMENT PRIMARY KEY,
  rutina_id BIGINT,
  nombre VARCHAR(120) NOT NULL,
  series INT,
  repeticiones INT,
  duracion_segs INT,
  intensidad VARCHAR(40),
  FOREIGN KEY (rutina_id) REFERENCES rutinas(rutina_id)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS asignacion_rutinas (
  asignacion_id BIGINT AUTO_INCREMENT PRIMARY KEY,
  rutina_id BIGINT,
  jugador_id BIGINT,
  equipo_id BIGINT,
  fecha_inicio DATE NOT NULL,
  fecha_fin DATE,
  frecuencia VARCHAR(80),
  comentarios TEXT,
  rol_id BIGINT,
  FOREIGN KEY (rutina_id) REFERENCES rutinas(rutina_id),
  FOREIGN KEY (jugador_id) REFERENCES miembros(miembro_id),
  FOREIGN KEY (equipo_id) REFERENCES equipos(equipo_id),
  FOREIGN KEY (rol_id) REFERENCES roles(rol_id)
) ENGINE=InnoDB;

-- ========================================
-- ACTIVIDADES Y MÉTRICAS
-- ========================================
CREATE TABLE IF NOT EXISTS actividades (
  actividad_id BIGINT AUTO_INCREMENT PRIMARY KEY,
  jugador_id BIGINT,
  miembro_id BIGINT,
  tipo VARCHAR(80) NOT NULL,
  fecha_inicio TIMESTAMP NOT NULL,
  fecha_fin TIMESTAMP,
  comentarios TEXT,
  FOREIGN KEY (jugador_id) REFERENCES miembros(miembro_id),
  FOREIGN KEY (miembro_id) REFERENCES miembros(miembro_id)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS signos_vitales (
  signos_id BIGINT AUTO_INCREMENT PRIMARY KEY,
  actividad_id BIGINT,
  fecha_registro TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  metrica VARCHAR(80),
  valor DECIMAL(12,4),
  unidad_medida VARCHAR(40),
  miembro_id BIGINT,
  FOREIGN KEY (actividad_id) REFERENCES actividades(actividad_id),
  FOREIGN KEY (miembro_id) REFERENCES miembros(miembro_id)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS alertas (
  alertas_id BIGINT AUTO_INCREMENT PRIMARY KEY,
  actividad_id BIGINT,
  fecha_registro TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  tipo VARCHAR(80),
  descripcion TEXT,
  gravedad VARCHAR(40),
  atendido_si_no TINYINT(1) DEFAULT 0,
  miembro_id BIGINT,
  FOREIGN KEY (actividad_id) REFERENCES actividades(actividad_id),
  FOREIGN KEY (miembro_id) REFERENCES miembros(miembro_id)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS reportes (
  reporte_id BIGINT AUTO_INCREMENT PRIMARY KEY,
  jugador_id BIGINT,
  fecha_registro TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  tipo VARCHAR(80),
  comentarios TEXT,
  FOREIGN KEY (jugador_id) REFERENCES miembros(miembro_id)
) ENGINE=InnoDB;

-- ========================================
-- DATOS DE PRUEBA
-- ========================================
INSERT INTO roles (nombre) VALUES 
('owner'), ('coach'), ('analyst'), ('assistant'), ('athlete')
ON DUPLICATE KEY UPDATE nombre = nombre;

INSERT INTO usuarios (nombre, correo, password, estado, rol_id) VALUES 
('Carla López', 'owner@zfut.test', 'hashed_pw', 'activo', 1),
('Diego Ramírez', 'coach@zfut.test', 'hashed_pw', 'activo', 2),
('Ana Pérez', 'analyst@zfut.test', 'hashed_pw', 'activo', 3),
('Luis Mora', 'assistant@zfut.test', 'hashed_pw', 'activo', 4),
('Jugador Uno', 'p1@zfut.test', 'hashed_pw', 'activo', 5),
('Jugador Dos', 'p2@zfut.test', 'hashed_pw', 'activo', 5)
ON DUPLICATE KEY UPDATE correo = correo;

INSERT INTO clubs (nombre, pais, owner_id) VALUES 
('ZFut FC', 'México', 1)
ON DUPLICATE KEY UPDATE nombre = nombre;

INSERT INTO suscripciones (nombre, precio, cant_players) VALUES
('Basic', 0.00, 25), ('Pro', 49.99, 100)
ON DUPLICATE KEY UPDATE nombre = nombre;

INSERT INTO club_suscripcion (club_id, suscripcion_id, owner_id, fecha_inicio, status, jugadores_act)
VALUES (1, 2, 1, NOW(), 'activa', 6)
ON DUPLICATE KEY UPDATE status = status;
