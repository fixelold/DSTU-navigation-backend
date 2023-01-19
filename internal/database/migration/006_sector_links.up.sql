INSERT INTO sector_link (number_sector, link, id_sector, id_link)
VALUES (unnest(array[31, 32, 32, 32, 33, 34, 34, 35, 35, 35, 36, 37]), 
        unnest(array[32, 31, 33, 34, 32, 32, 35, 34, 36, 37, 35, 35]),
        unnest(array[15, 16, 16, 16, 17, 18, 18, 19, 19, 19, 20, 21]),
        unnest(array[16, 15, 17, 18, 16, 16, 19, 18, 20, 21, 19, 19]))