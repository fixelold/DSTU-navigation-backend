INSERT INTO sector_link (number_sector, link, id_sector, id_link)
VALUES (unnest(array[131, 132, 132, 133]), 
        unnest(array[132, 131, 133, 132]),
        unnest(array[3,    7,   7,   11]),
        unnest(array[7,    3,   11,   7]));

-- INSERT INTO sector_link (number_sector, link, id_sector, id_link)
-- VALUES (unnest(array[141, 142, 142, 142, 143, 144, 144, 145, 145, 145, 146, 147]), 
--         unnest(array[142, 141, 143, 144, 142, 142, 145, 144, 146, 147, 145, 145]),
--         unnest(array[25,  26,  26,  26,  27,  28,  28,  29,  29,  29,  30,  31]),
--         unnest(array[26,  25,  27,  28,  26,  26,  29,  28,  30,  31,  29,  29]));