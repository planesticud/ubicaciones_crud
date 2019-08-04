-- object: ubicaciones | type: SCHEMA --
-- DROP SCHEMA IF EXISTS ubicaciones CASCADE;
CREATE SCHEMA ubicaciones;
-- ddl-end --

SET search_path TO pg_catalog,public,ubicaciones;
-- ddl-end --

-- object: ubicaciones.atributo_lugar | type: TABLE --
-- DROP TABLE IF EXISTS ubicaciones.atributo_lugar CASCADE;
CREATE TABLE ubicaciones.atributo_lugar (
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_atributo_lugar PRIMARY KEY (id),
	CONSTRAINT uq_nombre_atributo_lugar UNIQUE (nombre)

);
-- ddl-end --
COMMENT ON TABLE ubicaciones.atributo_lugar IS 'Tabla que almacena los atributos que puede tener un lugar';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.atributo_lugar.id IS 'Identificador unico de la tabla';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.atributo_lugar.nombre IS 'Indica el nombre del atributo para un lugar';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.atributo_lugar.descripcion IS 'Descripción opcional del atributo del lugar';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.atributo_lugar.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere. ';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.atributo_lugar.activo IS 'Estado del atributo del lugar';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.atributo_lugar.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD.';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.atributo_lugar.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.atributo_lugar.fecha_modificacion IS 'Fecha de la útlima modificación del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_atributo_lugar ON ubicaciones.atributo_lugar  IS 'Llave primaria de la tabla';
-- ddl-end --
COMMENT ON CONSTRAINT uq_nombre_atributo_lugar ON ubicaciones.atributo_lugar  IS 'Constraint unique para que el nombre de atributo_lugar no se repita';
-- ddl-end --

-- object: ubicaciones.lugar | type: TABLE --
-- DROP TABLE IF EXISTS ubicaciones.lugar CASCADE;
CREATE TABLE ubicaciones.lugar (
	id serial NOT NULL,
	nombre character varying(100) NOT NULL,
	tipo_lugar_id integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_lugar PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE ubicaciones.lugar IS 'Tabla que almacena los lugares vinculados a una ubicación de forma Jerarquica  Pais, Depto, Municipio, Ciudad, Localidades, Barrios, etc.';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.lugar.id IS 'Identificador unico de la tabla';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.lugar.nombre IS 'Corresponde al nombre del lugar';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.lugar.activo IS 'Campo que indica si el lugar es activo o no';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.lugar.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.lugar.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_lugar ON ubicaciones.lugar  IS 'Llave primaria de la tabla';
-- ddl-end --

-- object: ubicaciones.relacion_lugares | type: TABLE --
-- DROP TABLE IF EXISTS ubicaciones.relacion_lugares CASCADE;
CREATE TABLE ubicaciones.relacion_lugares (
	id serial NOT NULL,
	lugar_padre_id integer NOT NULL,
	lugar_hijo_id integer NOT NULL,
	tipo_relacion_lugar_id integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_relacion_lugares PRIMARY KEY (id),
	CONSTRAINT uq_padre_hijo UNIQUE (lugar_padre_id,lugar_hijo_id)

);
-- ddl-end --
COMMENT ON TABLE ubicaciones.relacion_lugares IS 'Tabla que almacena la relacion Jerarquica entre los lugares';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.relacion_lugares.id IS 'Identificador unico de la tabla';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.relacion_lugares.lugar_padre_id IS 'Identificador del lugar padre';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.relacion_lugares.lugar_hijo_id IS 'Identificador del lugar hijo';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.relacion_lugares.activo IS 'Indica si la relación entre los lugares se encuentra activa';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.relacion_lugares.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.relacion_lugares.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_relacion_lugares ON ubicaciones.relacion_lugares  IS 'Llave primaria de la tabla';
-- ddl-end --
COMMENT ON CONSTRAINT uq_padre_hijo ON ubicaciones.relacion_lugares  IS 'Restriccion para que solo pueda existir una unica relacion entre un padre y un hijo';
-- ddl-end --

-- object: ubicaciones.tipo_lugar | type: TABLE --
-- DROP TABLE IF EXISTS ubicaciones.tipo_lugar CASCADE;
CREATE TABLE ubicaciones.tipo_lugar (
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_tipo_lugar PRIMARY KEY (id),
	CONSTRAINT uq_nombre_tipo_lugar UNIQUE (nombre)

);
-- ddl-end --
COMMENT ON TABLE ubicaciones.tipo_lugar IS 'Almacena los parametros de tipo_lugar';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.tipo_lugar.id IS 'Identificador unico de la tabla';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.tipo_lugar.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parametro.';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.tipo_lugar.descripcion IS 'Descripción opcional del parámetro.';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.tipo_lugar.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.tipo_lugar.activo IS 'Campo de tipo boolean que inidica si el parametro esta activo';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.tipo_lugar.numero_orden IS ' En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.tipo_lugar.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.tipo_lugar.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_tipo_lugar ON ubicaciones.tipo_lugar  IS 'Llave primaria de la tabla';
-- ddl-end --
COMMENT ON CONSTRAINT uq_nombre_tipo_lugar ON ubicaciones.tipo_lugar  IS 'Constraint unique para que el nombre de tipo_lugar no se repita';
-- ddl-end --

-- object: ubicaciones.tipo_relacion_lugar | type: TABLE --
-- DROP TABLE IF EXISTS ubicaciones.tipo_relacion_lugar CASCADE;
CREATE TABLE ubicaciones.tipo_relacion_lugar (
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_tipo_relacion_lugar PRIMARY KEY (id),
	CONSTRAINT uq_nombre_tipo_relacion_lugar UNIQUE (nombre)

);
-- ddl-end --
COMMENT ON TABLE ubicaciones.tipo_relacion_lugar IS 'Tabla que contiene los tipos de relacion que pueden existir entre los lugares (Ejemplo: Capital)';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.tipo_relacion_lugar.id IS 'Identificador unico de la tabla';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.tipo_relacion_lugar.nombre IS 'Campo que indica el nombre del parametro.';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.tipo_relacion_lugar.descripcion IS 'Descripción opcional del parámetro.';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.tipo_relacion_lugar.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.tipo_relacion_lugar.activo IS 'Campo de tipo boolean que inidica si el parametro esta activo';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.tipo_relacion_lugar.numero_orden IS ' En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD.';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.tipo_relacion_lugar.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.tipo_relacion_lugar.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_tipo_relacion_lugar ON ubicaciones.tipo_relacion_lugar  IS 'Llave primaria de la tabla';
-- ddl-end --
COMMENT ON CONSTRAINT uq_nombre_tipo_relacion_lugar ON ubicaciones.tipo_relacion_lugar  IS 'Constraint unique para que el nombre de tipo_relacion_lugar no se repita.';
-- ddl-end --

-- object: ubicaciones.valor_atributo_lugar | type: TABLE --
-- DROP TABLE IF EXISTS ubicaciones.valor_atributo_lugar CASCADE;
CREATE TABLE ubicaciones.valor_atributo_lugar (
	id serial NOT NULL,
	valor character varying(150) NOT NULL,
	lugar_id integer NOT NULL,
	atributo_lugar_id integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_valor_atributo_lugar PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE ubicaciones.valor_atributo_lugar IS 'Tabla que almacena el valor de los atributos que tiene un lugar';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.valor_atributo_lugar.id IS 'Identificador unico de la tabla';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.valor_atributo_lugar.valor IS 'Valor para el atributo del lugar';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.valor_atributo_lugar.activo IS 'Indica el estado del registro';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.valor_atributo_lugar.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN ubicaciones.valor_atributo_lugar.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_valor_atributo_lugar ON ubicaciones.valor_atributo_lugar  IS 'Llave primaria de la tabla';
-- ddl-end --

-- object: fk_lugar_tipo_lugar | type: CONSTRAINT --
-- ALTER TABLE ubicaciones.lugar DROP CONSTRAINT IF EXISTS fk_lugar_tipo_lugar CASCADE;
ALTER TABLE ubicaciones.lugar ADD CONSTRAINT fk_lugar_tipo_lugar FOREIGN KEY (tipo_lugar_id)
REFERENCES ubicaciones.tipo_lugar (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_valor_atributo_lugar_atributo_lugar | type: CONSTRAINT --
-- ALTER TABLE ubicaciones.valor_atributo_lugar DROP CONSTRAINT IF EXISTS fk_valor_atributo_lugar_atributo_lugar CASCADE;
ALTER TABLE ubicaciones.valor_atributo_lugar ADD CONSTRAINT fk_valor_atributo_lugar_atributo_lugar FOREIGN KEY (atributo_lugar_id)
REFERENCES ubicaciones.atributo_lugar (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_valor_atributo_lugar_lugar | type: CONSTRAINT --
-- ALTER TABLE ubicaciones.valor_atributo_lugar DROP CONSTRAINT IF EXISTS fk_valor_atributo_lugar_lugar CASCADE;
ALTER TABLE ubicaciones.valor_atributo_lugar ADD CONSTRAINT fk_valor_atributo_lugar_lugar FOREIGN KEY (lugar_id)
REFERENCES ubicaciones.lugar (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_relacion_lugares_tipo_relacion_lugar | type: CONSTRAINT --
-- ALTER TABLE ubicaciones.relacion_lugares DROP CONSTRAINT IF EXISTS fk_relacion_lugares_tipo_relacion_lugar CASCADE;
ALTER TABLE ubicaciones.relacion_lugares ADD CONSTRAINT fk_relacion_lugares_tipo_relacion_lugar FOREIGN KEY (tipo_relacion_lugar_id)
REFERENCES ubicaciones.tipo_relacion_lugar (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_lugar_hijo | type: CONSTRAINT --
-- ALTER TABLE ubicaciones.relacion_lugares DROP CONSTRAINT IF EXISTS fk_lugar_hijo CASCADE;
ALTER TABLE ubicaciones.relacion_lugares ADD CONSTRAINT fk_lugar_hijo FOREIGN KEY (lugar_hijo_id)
REFERENCES ubicaciones.lugar (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_lugar_padre | type: CONSTRAINT --
-- ALTER TABLE ubicaciones.relacion_lugares DROP CONSTRAINT IF EXISTS fk_lugar_padre CASCADE;
ALTER TABLE ubicaciones.relacion_lugares ADD CONSTRAINT fk_lugar_padre FOREIGN KEY (lugar_padre_id)
REFERENCES ubicaciones.lugar (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --