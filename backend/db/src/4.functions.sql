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
