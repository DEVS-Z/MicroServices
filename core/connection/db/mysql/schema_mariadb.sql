-- ========================================
-- 1. BORRAR LA BASE DE DATOS (SI EXISTE)
-- ========================================
DROP DATABASE IF EXISTS zfut;

-- ========================================
-- 2. CREAR BASE DE DATOS Y USARLA
-- ========================================
CREATE DATABASE IF NOT EXISTS zfut;
USE zfut;

-- ========================================
-- 3. TABLAS DE USUARIOS Y ACCESO
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
-- 4. PERFIL DEL JUGADOR
-- ========================================

CREATE TABLE IF NOT EXISTS jugadores (
  jugador_id BIGINT AUTO_INCREMENT PRIMARY KEY,
  user_id BIGINT NOT NULL UNIQUE,
  posicion VARCHAR(40),
  weareable_id VARCHAR(80) UNIQUE,
  altura DECIMAL(5,2),
  peso DECIMAL(5,2),
  fecha_registro TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES usuarios(user_id)
) ENGINE=InnoDB;

-- ========================================
-- 5. RUTINAS Y EJERCICIOS
-- ========================================

CREATE TABLE IF NOT EXISTS rutinas (
  rutina_id BIGINT AUTO_INCREMENT PRIMARY KEY,
  creado_por_id BIGINT,
  nombre VARCHAR(120) NOT NULL,
  objetivo TEXT,
  tipo VARCHAR(80),
  nivel_dificultad VARCHAR(40),
  fecha_registro TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (creado_por_id) REFERENCES usuarios(user_id)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS ejercicios (
  ejercicio_id BIGINT AUTO_INCREMENT PRIMARY KEY,
  rutina_id BIGINT NOT NULL,
  nombre VARCHAR(120) NOT NULL,
  series INT,
  repeticiones INT,
  duracion_segs INT,
  intensidad VARCHAR(40),
  FOREIGN KEY (rutina_id) REFERENCES rutinas(rutina_id) ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS jugador_rutinas (
  asignacion_id BIGINT AUTO_INCREMENT PRIMARY KEY,
  rutina_id BIGINT NOT NULL,
  jugador_id BIGINT NOT NULL,
  fecha_asignacion DATE NOT NULL DEFAULT (CURRENT_DATE),
  comentarios TEXT,
  FOREIGN KEY (rutina_id) REFERENCES rutinas(rutina_id),
  FOREIGN KEY (jugador_id) REFERENCES jugadores(jugador_id)
) ENGINE=InnoDB;

-- ========================================
-- 6. ACTIVIDADES Y MÉTRICAS
-- ========================================

CREATE TABLE IF NOT EXISTS actividades (
  actividad_id BIGINT AUTO_INCREMENT PRIMARY KEY,
  jugador_id BIGINT NOT NULL,
  rutina_id BIGINT NULL,
  tipo VARCHAR(80) NOT NULL,
  fecha_inicio TIMESTAMP NOT NULL,
  fecha_fin TIMESTAMP,
  descripcion TEXT,
  FOREIGN KEY (jugador_id) REFERENCES jugadores(jugador_id),
  FOREIGN KEY (rutina_id) REFERENCES rutinas(rutina_id)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS metricas (
  metrica_id BIGINT AUTO_INCREMENT PRIMARY KEY,
  actividad_id BIGINT NOT NULL,
  fecha_registro TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  metrica VARCHAR(80),
  valor VARCHAR(100),
  unidad_medida VARCHAR(40),
  FOREIGN KEY (actividad_id) REFERENCES actividades(actividad_id) ON DELETE CASCADE
) ENGINE=InnoDB;

-- ========================================
-- 7. DATOS DE PRUEBA
-- ========================================

-- Insertar Roles
INSERT INTO roles (nombre) VALUES 
('owner'), ('coach'), ('analyst'), ('assistant'), ('athlete')
ON DUPLICATE KEY UPDATE nombre = nombre;

-- Insertar Usuarios
INSERT INTO usuarios (nombre, correo, password, estado, rol_id) VALUES 
('Diego Ramírez', 'coach@zfut.test', 'hashed_pw', 'activo', 2),
('Jugador Uno', 'p1@zfut.test', 'hashed_pw', 'activo', 5),
('Jugador Dos', 'p2@zfut.test', 'hashed_pw', 'activo', 5)
ON DUPLICATE KEY UPDATE correo = correo;

-- Insertar Perfiles de Jugador (vinculados a usuarios)
-- NOTA: Asegúrate que los user_id (2 y 3) coincidan con los IDs que se generaron arriba.
INSERT INTO jugadores (user_id, posicion, altura, peso, weareable_id) VALUES
(2, 'Delantero', 1.80, 75.5, 'WEAR-P1-001'),
(3, 'Defensa', 1.85, 80.1, 'WEAR-P2-002')
ON DUPLICATE KEY UPDATE user_id = user_id;

-- Insertar una Rutina (creada por el coach, user_id = 1)
-- NOTA: Asegúrate que el creado_por_id (1) coincida con el ID del coach.
INSERT INTO rutinas (creado_por_id, nombre, objetivo, nivel_dificultad) VALUES
(1, 'Fuerza Explosiva - Tren Inferior', 'Aumentar potencia de salto y sprint', 'Intermedio')
ON DUPLICATE KEY UPDATE nombre = nombre;

-- Insertar Ejercicios para esa Rutina (asumiendo que la rutina_id es 1)
INSERT INTO ejercicios (rutina_id, nombre, series, repeticiones) VALUES
(1, 'Sentadilla con Salto', 4, 10),
(1, 'Desplantes pliométricos', 3, 12),
(1, 'Box Jumps', 4, 8);

-- Asignar la Rutina 1 al Jugador 1 (asumiendo IDs 1 y 1)
INSERT INTO jugador_rutinas (rutina_id, jugador_id, comentarios) VALUES
(1, 1, 'Lunes y Jueves, enfocar en técnica');

-- Registrar una Actividad para el Jugador 1
INSERT INTO actividades (jugador_id, tipo, fecha_inicio, fecha_fin, descripcion) VALUES
(1, 'Entrenamiento', '2025-11-10 09:00:00', '2025-11-10 10:30:00', 'Sesión de campo');

-- Registrar Métricas para esa Actividad (asumiendo actividad_id es 1)
INSERT INTO metricas (actividad_id, metrica, valor, unidad_medida) VALUES
(1, 'Ritmo Cardíaco Promedio', '145', 'bpm'),
(1, 'Distancia Total', '7.8', 'km'),
(1, 'Sprints > 20km/h', '12', 'conteo'),
(1, 'Alerta', 'Pico de fatiga detectado', 'status');