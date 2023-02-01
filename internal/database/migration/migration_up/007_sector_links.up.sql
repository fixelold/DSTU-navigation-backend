INSERT INTO sector_link (number_sector, link, id_sector, id_link)
VALUES (unnest(array[131, 132, 132, 132, 133, 134, 134, 135, 135, 135, 136, 137]), 
        unnest(array[132, 131, 133, 134, 132, 132, 135, 134, 136, 137, 135, 135]),
        unnest(array[17, 18, 18, 18, 19, 20, 20, 21, 21, 21, 22, 23]),
        unnest(array[18, 17, 19, 20, 18, 18, 21, 20, 22, 23, 21, 21]))