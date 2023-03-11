-- INSERT INTO auditorium_position (id_auditorium, x, y, widht, height) 
-- VALUES (
--     unnest(array[30,    31,     32,     33,     34,     35]),
--     unnest(array[611,   667,    36,     791,    1015,   1017]),
--     unnest(array[2255,  2639,   3033,   3029,   2629,   2099]),
--     unnest(array[165,   111,    745,    220,    185,    179]),
--     unnest(array[370,   253,    174,    178,    580,    522])
-- );

-- INSERT INTO auditorium_position (id_auditorium, x, y, widht, height) 
-- VALUES (
--     unnest(array[1,      2,      3,      4,      5,     6,      7,      8,      9,      10,     11]),
--     unnest(array[50,   409,    806,     1009,    52,    49,     572,    694,    1011,   1012,   603]),
--     unnest(array[48,    46,     49,     205,     271,   457,    457,    455,    446,    703,    658]),
--     unnest(array[353,   387,    192,    179,     207,   512,    106,    72,     177,    181,    169]),
--     unnest(array[211,   213,    208,    113,    175,   192,    192,     63,     248,    626,    242])
-- );

INSERT INTO auditorium_position (id_auditorium, x, y, widht, height) 
VALUES (
    unnest(array[30,   31,    32,    33,     34,     35]),
    unnest(array[229,  230,   1,     297,    438,   438]),
    unnest(array[659,  734,   914,   913,   757,   635]),
    unnest(array[59,   58,    295,   139,    62,    62]),
    unnest(array[74,   50,    85,    86,    243,    121])
);

INSERT INTO auditorium_position (id_auditorium, x, y, widht, height) 
VALUES (
    unnest(array[1,    2,     3,     4,      5,    6,     7,     8,      9,     10,     11]),
    unnest(array[1,    161,   290,   434,    1,    1,     230,   267,    438,   438,   230]),
    unnest(array[1,    1,     1,     40,     87,   217,   216,   217,    218,   297,    269]),
    unnest(array[159,  129,   105,   61,     78,   228,   37,    23,     63,    63,    60]),
    unnest(array[85,   86,    88,    106,    128,  50,    53,    26,     78,    136,    50])
);

INSERT INTO auditorium_position (id_auditorium, x, y, widht, height) 
VALUES (
    unnest(array[12,   13,   14,    15,    16,   17,   18,    19,     20]),
    unnest(array[230,  230,  230,   230,   230,  230,  230,   230,    438]),
    unnest(array[319,  344,  368,   400,   430,  533,  580,   616,    565]),
    unnest(array[60,   61,   160,   60,    59,   60,   60,    59,     62]),
    unnest(array[25,   23,   32,    30,    103,  47,   36,    42,     69])
);