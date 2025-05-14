CREATE OR REPLACE FUNCTION update_reservation_rooms()
RETURNS TRIGGER AS $$
BEGIN
  INSERT INTO "reservation_rooms" ("reservation_id", "room_id")
  SELECT NEW."id", UNNEST(NEW."room_ids");
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER after_reservation_insert
AFTER INSERT ON "reservations"
FOR EACH ROW
EXECUTE FUNCTION update_reservation_rooms();
