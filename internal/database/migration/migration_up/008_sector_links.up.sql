INSERT INTO sector_link (number_sector, link, id_sector, id_link)
VALUES (unnest(array[131, 132, 132, 132, 133, 134, 134, 135, 135, 135, 136, 137]), 
        unnest(array[132, 131, 133, 134, 132, 132, 135, 134, 136, 137, 135, 135]),
        unnest(array[17, 18, 18, 18, 19, 20, 20, 21, 21, 21, 22, 23]),
        unnest(array[18, 17, 19, 20, 18, 18, 21, 20, 22, 23, 21, 21]));

INSERT INTO sector_link (number_sector, link, id_sector, id_link)
VALUES (unnest(array[141, 142, 142, 142, 143, 144, 144, 145, 145, 145, 146, 147]), 
        unnest(array[142, 141, 143, 144, 142, 142, 145, 144, 146, 147, 145, 145]),
        unnest(array[25,  26,  26,  26,  27,  28,  28,  29,  29,  29,  30,  31]),
        unnest(array[26,  25,  27,  28,  26,  26,  29,  28,  30,  31,  29,  29]));