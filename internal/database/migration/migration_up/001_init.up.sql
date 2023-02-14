CREATE TABLE IF NOT EXISTS "user"(
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY (START WITH 1  INCREMENT BY 1),
  email VARCHAR,
  password VARCHAR
);

CREATE TABLE IF NOT EXISTS "building"(
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY (START WITH 1  INCREMENT BY 1),
  number INT
);

CREATE TABLE IF NOT EXISTS "floor"(
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY (START WITH 1  INCREMENT BY 1),
    number INT,
    id_building INT
) ;

CREATE TABLE IF NOT EXISTS "sector"(
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY (START WITH 1  INCREMENT BY 1),
    number INT,
    id_floor INT
) ;

CREATE TABLE IF NOT EXISTS "transition"(
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY (START WITH 1  INCREMENT BY 1),
    id_sectors INT
);

CREATE TABLE IF NOT EXISTS "auditorium"(
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY (START WITH 1  INCREMENT BY 1),
    number VARCHAR,
    id_sector INT
);

CREATE TABLE IF NOT EXISTS "sector_link"(
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY (START WITH 1  INCREMENT BY 1),
  id_sector INT,
  number_sector INT,
  id_link INT,
  link INT
);

CREATE TABLE IF NOT EXISTS "auditorium_position" (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY (START WITH 1  INCREMENT BY 1),
  id_auditorium INT,
  x INT,
  y INT,
  widht INT,
  height INT
);

CREATE TABLE IF NOT EXISTS "border_points" (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY (START WITH 1  INCREMENT BY 1),
  id_auditorium INT,
  x INT,
  y INT,
  widht INT,
  height INT
);

ALTER TABLE "floor" ADD FOREIGN KEY (id_building) REFERENCES "building" (id);

ALTER TABLE "sector" ADD FOREIGN KEY (id_floor) REFERENCES "floor" (id);

ALTER TABLE "transition" ADD FOREIGN KEY (id_sectors) REFERENCES "sector" (id);

ALTER TABLE "auditorium" ADD FOREIGN KEY (id_sector) REFERENCES "sector" (id);

ALTER TABLE "sector_link" ADD FOREIGN KEY (id_sector) REFERENCES "sector" (id);

ALTER TABLE "sector_link" ADD FOREIGN KEY (id_link) REFERENCES "sector" (id);

ALTER TABLE "auditorium_position" ADD FOREIGN KEY (id_auditorium) REFERENCES "auditorium" (id);

ALTER TABLE "border_points" ADD FOREIGN KEY (id_auditorium) REFERENCES "auditorium" (id);