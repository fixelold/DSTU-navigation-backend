INSERT INTO transition_sector_border_points (id_transition_sector, x, y, widht, height) 
VALUES (
    unnest(array[1,   2]),
    unnest(array[437, 437]),
    unnest(array[145, 145]),
    unnest(array[1,   1]),
    unnest(array[74,  74])
);

INSERT INTO transition_sector_border_points (id_transition_sector, x, y, widht, height) 
VALUES (
    unnest(array[3,   4]),
    unnest(array[437, 437]),
    unnest(array[145, 145]),
    unnest(array[1,   1]),
    unnest(array[74,  74])
);