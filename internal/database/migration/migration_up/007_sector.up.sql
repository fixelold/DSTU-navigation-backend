-- INSERT INTO sector(number, id_floor, id_transition)
-- VALUES (unnest(array[111, 112, 113, 114, 115, 116, 117]), 1);

-- INSERT INTO sector(number, id_floor)
-- VALUES (unnest(array[121, 122, 123, 124, 125, 126, 127]), 2);

-- INSERT INTO sector(number, id_floor)
-- VALUES (unnest(array[131, 132, 133, 134, 135, 136, 137]), 3);

-- INSERT INTO sector(number, id_floor)
-- VALUES (unnest(array[141, 142, 143, 144, 145, 146, 147]), 4);

INSERT INTO sector(number, id_floor, id_transition)
VALUES 
    (111, 1, 1),
    (121, 2, 1),
    (131, 3, 1),
    (141, 4, 1),

    (112, 1, 2),
    (122, 2, 2),
    (132, 3, 2),
    (142, 4, 2),

    (113, 1, 3),
    (123, 2, 3),
    (133, 3, 3),
    (143, 4, 3);