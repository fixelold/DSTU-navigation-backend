-- INSERT INTO sector_link (number_sector, link, id_sector, id_link)
-- VALUES (unnest(array[131, 132, 132, 133]), 
--         unnest(array[132, 131, 133, 132]),
--         unnest(array[3,    7,   7,   11]),
--         unnest(array[7,    3,   11,   7]));

INSERT INTO sector_link (number_sector, link, id_sector, id_link)
VALUES
        (131, 132, 3, 7),
        (132, 131, 7, 3),
        (132, 133, 7, 11),
        (133, 132, 11, 7),
        (132, 134, 7, 15),
        (134, 132, 15, 7),
        (134, 135, 15, 19),
        (135, 134, 19, 15),
        (135, 136, 19, 23),
        (136, 135, 23, 19);