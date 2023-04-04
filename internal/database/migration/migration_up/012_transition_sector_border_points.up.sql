INSERT INTO aud_border_points (transition_sector_border_points, x, y, widht, height) 
VALUES (
    unnest(array[1,   2]),
    unnest(array[11,  437]),
    unnest(array[784, 145]),
    unnest(array[92,  1]),
    unnest(array[1,   74])
);