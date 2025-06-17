CREATE OR REPLACE FUNCTION update_reservation_rooms()
RETURNS TRIGGER AS $$
BEGIN
  INSERT INTO "reservation_rooms" ("reservation_id", "room_id")
  SELECT NEW."id", NEW."room_id";
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER after_reservation_insert
AFTER INSERT ON "reservations"
FOR EACH ROW
EXECUTE FUNCTION update_reservation_rooms();

CREATE OR REPLACE FUNCTION update_hotel_rating_after_review()
RETURNS TRIGGER AS $$
BEGIN
  UPDATE "hotel_ratings"
  SET "current_rating" = (
    SELECT ROUND(AVG("rating")::numeric, 2)
    FROM "reviews"
    WHERE "hotel_id" = NEW."hotel_id"
  )
  WHERE "hotel_id" = NEW."hotel_id";
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER after_review_insert
AFTER INSERT ON "reviews"
FOR EACH ROW
EXECUTE FUNCTION update_hotel_rating_after_review();