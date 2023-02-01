INSERT INTO sector_link (number_sector, link, id_sector, id_link)
VALUES (unnest(array[131, 132, 132, 132, 133, 134, 134, 135, 135, 135, 136, 137]), 
        unnest(array[132, 131, 133, 134, 132, 132, 135, 134, 136, 137, 135, 135]),
        unnest(array[15, 16, 16, 16, 17, 18, 18, 19, 19, 19, 20, 21]),
        unnest(array[16, 15, 17, 18, 16, 16, 19, 18, 20, 21, 19, 19]))