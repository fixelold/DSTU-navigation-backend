INSERT INTO auditorium (number, id_sector) VALUES
    (unnest(array['1-319а', '1-319', '1-317а', '1-317', '1-319б', '1-320', '1-321', '1-321а', '1-315']), 3);

INSERT INTO auditorium (number, id_sector) VALUES
    (unnest(array['1-313', '1-323', '1-325', '1-326', '1-326а', '1-326в', '1-326б', '1-327а', '1-327', '1-328', '1-330', '1-338а', '1-340', '1-341', '1-343', '1-344', '1-346', '1-347а', '1-305а', '1-309', '1-310']), 7);

INSERT INTO auditorium (number, id_sector) VALUES
    (unnest(array['1-331', '1-333', '1-333а', '1-336']), 11);

INSERT INTO auditorium (number, id_sector) VALUES
    (unnest(array['1-305', '1-347', '1-348', '1-349', '1-350', '1-351', '1-352', '1-353', '1-399', '1-398', '1-397', '1-353а', '1-354', '1-355']), 15);

INSERT INTO auditorium (number, id_sector) VALUES
    (unnest(array['1-396', '1-394', '1-393', '1-392', '1-391', '1-391а', '1-356', '1-358', '1-359', '1-359а', '1-360', '1-361', '1-362', '1-374', '1-375', '1-375а', '1-376', '1-376а', '1-376б', '1-379']), 19);

INSERT INTO auditorium (number, id_sector) VALUES
    (unnest(array['1-363', '1-364', '1-365', '1-367', '1-367а', '1-369']), 23);

-- INSERT INTO auditorium (number, id_sector) VALUES
--     (unnest(array['1-348', '1-349', '1-350', '1-305', '1-351', '1-352', '1-353', '1-353а', '1-354', '1-355', '1-396', '1-397', '1-398', '1-399']), 20);

-- INSERT INTO auditorium (number, id_sector) VALUES
--     (unnest(array['356', '358', '394', '393', '392', '359', '359а', '360', '361', '362', '374', '375', '375а', '376', '376а', '376б', '391', '378']),21);

-- INSERT INTO auditorium (number, id_sector) VALUES
--     (unnest(array['391а', '388', '386', '385', '384', '383']), 22);

-- INSERT INTO auditorium (number, id_sector) VALUES
--     (unnest(array['363', '364', '365', '367', '367а', '369']), 23);

-- INSERT INTO auditorium (number, id_sector) VALUES
--     (unnest(array['1-464']), 4);