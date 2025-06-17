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
  INSERT INTO "hotel_ratings" ("hotel_id", "current_rating", "review_count")
  VALUES (
    NEW."hotel_id",
    (SELECT ROUND(AVG("rating")::numeric, 2)
     FROM "reviews"
     WHERE "hotel_id" = NEW."hotel_id"),
    (SELECT COUNT(*)
     FROM "reviews"
     WHERE "hotel_id" = NEW."hotel_id")
  )
  ON CONFLICT ("hotel_id")
  DO UPDATE SET 
    "current_rating" = (
      SELECT ROUND(AVG("rating")::numeric, 2)
      FROM "reviews"
      WHERE "hotel_id" = NEW."hotel_id"
    ),
    "review_count" = (
      SELECT COUNT(*)
      FROM "reviews"
      WHERE "hotel_id" = NEW."hotel_id"
    );
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER after_review_insert
AFTER INSERT ON "reviews"
FOR EACH ROW
EXECUTE FUNCTION update_hotel_rating_after_review();