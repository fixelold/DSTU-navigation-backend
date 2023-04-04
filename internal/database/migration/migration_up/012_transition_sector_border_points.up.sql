INSERT INTO transition_sector_border_points (id_transition_sector, x, y, widht, height) 
VALUES (
    unnest(array[1,   2]),
    unnest(array[11,  437]),
    unnest(array[784, 145]),
    unnest(array[92,  1]),
    unnest(array[1,   74])
);