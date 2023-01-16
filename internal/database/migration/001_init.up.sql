CREATE TABLE "user" (
  id smallint GENERATED ALWAYS AS IDENTITY (START WITH 1  INCREMENT BY 1),
  email text,
  password text,
  PRIMARY KEY (id)
);

CREATE TABLE "building" (
  id smallint GENERATED ALWAYS AS IDENTITY (START WITH 1  INCREMENT BY 1),
  number smallint,
  PRIMARY KEY (id)
);

CREATE TABLE "floor" (
    id smallint GENERATED ALWAYS AS IDENTITY (START WITH 1  INCREMENT BY 1),
    number smallint,
    id_building smallint,
    PRIMARY KEY (id)
) ;

CREATE TABLE "sector" (
    id smallint GENERATED ALWAYS AS IDENTITY (START WITH 1  INCREMENT BY 1),
    number smallint,
    id_floor smallint,
    PRIMARY KEY (id)
) ;

CREATE TABLE "transition" (
    id smallint GENERATED ALWAYS AS IDENTITY (START WITH 1  INCREMENT BY 1),
    id_sectors smallint,
    PRIMARY KEY (id)
);

CREATE TABLE "auditorium" (
    id smallint GENERATED ALWAYS AS IDENTITY (START WITH 1  INCREMENT BY 1),
    number varchar,
    id_sector smallint,
    PRIMARY KEY (id)
);

ALTER TABLE "floor" ADD FOREIGN KEY (id_building) REFERENCES "building" (id);

ALTER TABLE "sector" ADD FOREIGN KEY (id_floor) REFERENCES "floor" (id);

ALTER TABLE "transition" ADD FOREIGN KEY (id_sectors) REFERENCES "sector" (id);

ALTER TABLE "auditorium" ADD FOREIGN KEY (id_sector) REFERENCES "sector" (id);