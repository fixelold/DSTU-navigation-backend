-- INSERT INTO aud_border_points (id_auditorium, x, y, widht, height) 
-- VALUES (
--     unnest(array[30,    31,    32,     33,     34,     35]),
--     unnest(array[611,   667,    36,     791,    1015,   1017]),
--     unnest(array[2255,  2639,   3033,   3029,   2629,   2099]),
--     unnest(array[1,     1,      745,    220,    1,      1]),
--     unnest(array[370,   253,    1,      1,      580,    522])
-- );

-- INSERT INTO aud_border_points (id_auditorium, x, y, widht, height) 
-- VALUES (
--     unnest(array[1,     2,    3,    4,      5,      6,      7,     8,      9,      10,   11]),
--     unnest(array[269,   409,  806,  1009,   257,    271,    572,   694,    1011,   1012, 603]),
--     unnest(array[263,   254,  256,  273,    271,    444,    457,   455,    446,    703,  658]),
--     unnest(array[134,   387,  192,  1,      1,      287,    106,   72,     1,      1,    1]),
--     unnest(array[1,     1,    1,    58,     186,    1,      1,     1,     248,     626,  242])
-- );

INSERT INTO aud_border_points (id_auditorium, x, y, widht, height) 
VALUES (
     unnest(array[30,  31,    32,    33,     34,     35]),
    unnest(array[287,  287,   1,     297,    438,   438]),
    unnest(array[659,  734,   914,   913,    757,   635]),
    unnest(array[1,    1,     295,   139,    1,    1]),
    unnest(array[74,   50,    1,    1,       156,    121])
);

INSERT INTO aud_border_points (id_auditorium, x, y, widht, height) 
VALUES (
    unnest(array[1,   2,     3,    4,    5,    6,     7,     8,      9,     10,     11]),
    unnest(array[80,  161,   290,  435,  80,   1,     230,   267,    438,   434,   287]),
    unnest(array[87,  87,    86,   88,   88,   217,   216,   217,    218,   297,    269]),
    unnest(array[80,  130,   105,  1,    1,    228,   37,    23,     1,     1,      1]),
    unnest(array[1,   1,     1,    58,   128,  1,     1,     1,      80,    135,    50])
);