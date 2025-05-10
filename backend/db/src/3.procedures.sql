CREATE OR REPLACE PROCEDURE delete_past_due_reservations()
LANGUAGE plpgsql
AS $$
BEGIN
  DELETE FROM "reservations"
  WHERE "payment_info_id" IN (
    SELECT "id"
    FROM "payments"
    WHERE "due_date" < CURRENT_DATE AND "fulfilled" = false
  );
END;
$$;

CREATE OR REPLACE PROCEDURE calculate_hotel_ratings()
LANGUAGE plpgsql
AS $$
BEGIN
  UPDATE "hotel_ratings"
  SET "current_rating" = subquery.avg_rating
  FROM (
    SELECT "hotel_id", ROUND(AVG("rating")::numeric, 2) AS avg_rating
    FROM "reviews"
    GROUP BY "hotel_id"
  ) AS subquery
  WHERE "hotel_ratings"."hotel_id" = subquery."hotel_id";
END;
$$;

CREATE OR REPLACE PROCEDURE daily_update()
LANGUAGE plpgsql
AS $$
BEGIN
  PERFORM delete_past_due_reservations();
  PERFORM calculate_hotel_ratings();
END;
$$;
