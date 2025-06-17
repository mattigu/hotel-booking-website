--- Vacancy Rate (%) = (Total Available Room Nights - Total Occupied Room Nights) / Total Available Room Nights * 100
CREATE OR REPLACE FUNCTION calculate_hotel_vacancy(
    hotel_id_input INTEGER,
    start_date_input DATE,
    end_date_input DATE
) RETURNS NUMERIC AS $$
DECLARE
    total_available_room_nights INTEGER;
    total_occupied_room_nights INTEGER;
    vacancy_rate NUMERIC;
BEGIN
    SELECT COUNT(*) * (end_date_input - start_date_input + 1)
    INTO total_available_room_nights
    FROM rooms
    WHERE hotel_id = hotel_id_input;

    SELECT COUNT(*)
    INTO total_occupied_room_nights
    FROM reservation_rooms rr
    JOIN reservations r ON rr.reservation_id = r.id
    WHERE r.hotel_id = hotel_id_input
      AND r.start_date <= end_date_input
      AND r.end_date >= start_date_input;

    IF total_available_room_nights = 0 THEN
        RETURN 0;
    END IF;

    vacancy_rate := (total_available_room_nights - total_occupied_room_nights) * 100.0 / total_available_room_nights;

    RETURN vacancy_rate;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE function add_new_user(
    name text,
    surname text,
    phone_number text
) RETURNS boolean as $$
DECLARE
    found_customer INTEGER;
BEGIN
    SELECT COUNT(*)
    INTO found_customer
    FROM customers c
    WHERE c.name == name
    AND c.surname == surname
    AND c.phone_number == phone_number;
        IF found_customer = 0 THEN
        INSERT INTO customers VALUES (
            name, surname, phone_number
        );
        RETURN 1;
    END IF;
    RETURN 0;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION is_room_available(
    room_id_input INTEGER,
    start_date_input DATE,
    end_date_input DATE
) RETURNS BOOLEAN AS $$
DECLARE
    conflicting_reservations INTEGER;
BEGIN
    SELECT COUNT(*)
    INTO conflicting_reservations
    FROM reservation_rooms rr
    JOIN reservations r ON rr.reservation_id = r.id
    WHERE rr.room_id = room_id_input
      AND r.start_date < end_date_input
      AND r.end_date > start_date_input;
    
    IF conflicting_reservations > 0 THEN
        RETURN FALSE;
    ELSE
        RETURN TRUE;
    END IF;
END;
$$ LANGUAGE plpgsql;
