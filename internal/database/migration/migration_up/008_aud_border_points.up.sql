INSERT INTO aud_border_points (id_auditorium, x, y, widht, height) 
VALUES (
    unnest(array[30,    31,    32,     33,     34,     35]),
    unnest(array[611,   667,    36,     791,    1015,   1017]),
    unnest(array[2255,  2639,   3033,   3029,   2629,   2099]),
    unnest(array[1,     1,      745,    220,    1,      1]),
    unnest(array[370,   253,    1,      1,      580,    522])
);