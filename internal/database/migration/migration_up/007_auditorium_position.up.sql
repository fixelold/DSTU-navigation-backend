INSERT INTO auditorium_position (id_auditorium, x, y, widht, height) 
VALUES (
    unnest(array[30,    31,     32,     33,     34,     35]),
    unnest(array[611,   667,    36,     791,    1015,   1017]),
    unnest(array[2255,  2639,   3033,   3029,   2629,   2099]),
    unnest(array[165,   111,    745,    220,    185,    179]),
    unnest(array[370,   253,    174,    178,    580,    522])
);

INSERT INTO auditorium_position (id_auditorium, x, y, widht, height) 
VALUES (
    unnest(array[1,      2,      3,      4,      5,     6,      7,      8,      9,      10,     11]),
    unnest(array[50,   409,    806,     1009,    52,    49,     572,    694,    1011,   1012,   603]),
    unnest(array[48,    46,     49,     205,     271,   457,    457,    455,    446,    703,    658]),
    unnest(array[353,   387,    192,    179,     207,   512,    106,    72,     177,    181,    169]),
    unnest(array[211,   213,    208,    113,    175,   192,    192,     63,     248,    626,    242])
);