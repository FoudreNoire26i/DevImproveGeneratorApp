CREATE OR REPLACE FUNCTION check_number_of_row_chats()
RETURNS TRIGGER AS
$body$
BEGIN
	IF (SELECT count(*) FROM chats) > 10000
	THEN
		RAISE EXCEPTION 'INSERT statement exceeding maximum number of rows for this table';
END IF;
	RETURN NEW;
END;
$body$
LANGUAGE plpgsql;

CREATE TRIGGER tr_check_number_of_row_chats
BEFORE INSERT ON chats
FOR EACH ROW EXECUTE PROCEDURE check_number_of_row_chats();

CREATE OR REPLACE FUNCTION check_number_of_row_placeportids()
RETURNS TRIGGER AS
$body$
BEGIN
	IF (SELECT count(*) FROM placeportids) > 10000
	THEN
		RAISE EXCEPTION 'INSERT statement exceeding maximum number of rows for this table';
END IF;
	RETURN NEW;
END;
$body$
LANGUAGE plpgsql;

CREATE TRIGGER tr_check_number_of_row_placeportids
BEFORE INSERT ON placeportids
FOR EACH ROW EXECUTE PROCEDURE check_number_of_row_placeportids();

CREATE OR REPLACE FUNCTION check_number_of_row_interviews()
RETURNS TRIGGER AS
$body$
BEGIN
	IF (SELECT count(*) FROM interviews) > 10000
	THEN
		RAISE EXCEPTION 'INSERT statement exceeding maximum number of rows for this table';
END IF;
	RETURN NEW;
END;
$body$
LANGUAGE plpgsql;

CREATE TRIGGER tr_check_number_of_row_interviews
BEFORE INSERT ON interviews
FOR EACH ROW EXECUTE PROCEDURE check_number_of_row_interviews();

CREATE OR REPLACE FUNCTION check_number_of_row_bateauids()
RETURNS TRIGGER AS
$body$
BEGIN
	IF (SELECT count(*) FROM bateauids) > 10000
	THEN
		RAISE EXCEPTION 'INSERT statement exceeding maximum number of rows for this table';
END IF;
	RETURN NEW;
END;
$body$
LANGUAGE plpgsql;

CREATE TRIGGER tr_check_number_of_row_bateauids
BEFORE INSERT ON bateauids
FOR EACH ROW EXECUTE PROCEDURE check_number_of_row_bateauids();

CREATE OR REPLACE FUNCTION check_number_of_row_notifications()
RETURNS TRIGGER AS
$body$
BEGIN
	IF (SELECT count(*) FROM notifications) > 10000
	THEN
		RAISE EXCEPTION 'INSERT statement exceeding maximum number of rows for this table';
END IF;
	RETURN NEW;
END;
$body$
LANGUAGE plpgsql;

CREATE TRIGGER tr_check_number_of_row_notifications
BEFORE INSERT ON notifications
FOR EACH ROW EXECUTE PROCEDURE check_number_of_row_notifications();

CREATE OR REPLACE FUNCTION check_number_of_row_comments()
RETURNS TRIGGER AS
$body$
BEGIN
	IF (SELECT count(*) FROM comments) > 10000
	THEN
		RAISE EXCEPTION 'INSERT statement exceeding maximum number of rows for this table';
END IF;
	RETURN NEW;
END;
$body$
LANGUAGE plpgsql;

CREATE TRIGGER tr_check_number_of_row_comments
BEFORE INSERT ON comments
FOR EACH ROW EXECUTE PROCEDURE check_number_of_row_comments();

