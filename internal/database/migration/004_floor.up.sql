INSERT INTO floor(number, id_building) 
VALUES (unnest(array[1, 2, 3, 4]), 2);