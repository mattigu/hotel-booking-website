CREATE TABLE "hotels" (
  "id" SERIAL PRIMARY KEY,
  "address_id" integer NOT NULL,
  "name" text NOT NULL,
  "description" text,
  "star_standard" integer NOT NULL
);

CREATE TABLE "room_amount" (
  "hotel_id" integer NOT NULL,
  "room_id" integer NOT NULL,
  "amount" integer NOT NULL
);

CREATE TABLE "rooms" (
  "id" SERIAL PRIMARY KEY,
  "hotel_id" integer NOT NULL,
  "room_number" integer NOT NULL,
  "single_bed_num" integer NOT NULL,
  "double_bed_num" integer NOT NULL,
  "base_price" integer NOT NULL
);

CREATE TABLE "amenities" (
  "id" SERIAL PRIMARY KEY,
  "name" text NOT NULL,
  "description" text NOT NULL
);

CREATE TABLE "room_amenities" (
  "room_id" integer NOT NULL,
  "amenity_id" integer NOT NULL
);

CREATE TABLE "hotel_amenities" (
  "hotel_id" integer NOT NULL,
  "amenity_id" integer NOT NULL
);

CREATE TABLE "customers" (
  "id" SERIAL PRIMARY KEY,
  "name" text NOT NULL,
  "surname" text NOT NULL,
  "phone_number" text NOT NULL,
  "email" text
);

CREATE TABLE "reservations" (
  "id" SERIAL PRIMARY KEY,
  "customer_id" integer NOT NULL,
  "hotel_id" integer NOT NULL,
  "room_id" integer NOT NULL,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "payment_info_id" integer NOT NULL
);

CREATE TABLE "promotions" (
  "id" SERIAL PRIMARY KEY,
  "hotel_id" integer NOT NULL,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "discount_flat" integer,
  "discount_pct" integer
);

CREATE TABLE "addresses" (
  "id" SERIAL PRIMARY KEY,
  "city" text NOT NULL,
  "street" text NOT NULL,
  "zip_code" text NOT NULL,
  "house_number" integer NOT NULL,
  "country_id" integer NOT NULL
);

CREATE TABLE "countries" (
  "id" SERIAL PRIMARY KEY,
  "name" text NOT NULL
);

CREATE TABLE "reviews" (
  "id" SERIAL PRIMARY KEY,
  "username" text NOT NULL,
  "hotel_id" integer NOT NULL,
  "rating" integer NOT NULL,
  "review_text" text NOT NULL,
  "upload_date" date NOT NULL
);

CREATE TABLE "payments" (
  "id" SERIAL PRIMARY KEY,
  "payment_type" text NOT NULL,
  "payment_data" text,
  "due_date" date NOT NULL,
  "amount" text NOT NULL,
  "fulfilled" boolean DEFAULT false
);

CREATE TABLE "avg_price_history" (
  "id" SERIAL PRIMARY KEY,
  "hotel_id" integer NOT NULL,
  "period_start" date NOT NULL,
  "period_end" date NOT NULL,
  "avg_price" integer NOT NULL
);

CREATE TABLE "vacancy_history" (
  "id" SERIAL PRIMARY KEY,
  "hotel_id" integer NOT NULL,
  "period_start" date NOT NULL,
  "period_end" date NOT NULL,
  "vacancies" integer NOT NULL
);

CREATE TABLE "hotel_ratings" (
  "hotel_id" integer PRIMARY KEY,
  "current_rating" numeric(3,2) NOT NULL,
  "review_count" integer NOT NULL DEFAULT 0
);

CREATE TABLE "reservation_addons" (
  "id" integer PRIMARY KEY,
  "hotel_id" integer NOT NULL,
  "name" text NOT NULL,
  "price" integer NOT NULL
);

CREATE TABLE "reservation_rooms" (
  "reservation_id" integer,
  "room_id" integer,
  PRIMARY KEY ("reservation_id", "room_id")
);

CREATE TABLE "reservation_add_ons" (
  "reservation_id" integer NOT NULL,
  "addon_id" integer NOT NULL,
  PRIMARY KEY ("reservation_id", "addon_id")
);

ALTER TABLE "room_amenities" ADD FOREIGN KEY ("amenity_id") REFERENCES "amenities" ("id");

ALTER TABLE "hotel_amenities" ADD FOREIGN KEY ("amenity_id") REFERENCES "amenities" ("id");

ALTER TABLE "hotels" ADD FOREIGN KEY ("address_id") REFERENCES "addresses" ("id") ON DELETE CASCADE;

ALTER TABLE "hotel_amenities" ADD FOREIGN KEY ("hotel_id") REFERENCES "hotels" ("id") ON DELETE CASCADE;

ALTER TABLE "rooms" ADD FOREIGN KEY ("hotel_id") REFERENCES "hotels" ("id") ON DELETE CASCADE;

ALTER TABLE "room_amenities" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id") ON DELETE CASCADE;

ALTER TABLE "reservations" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id") ON DELETE CASCADE;

ALTER TABLE "reservations" ADD FOREIGN KEY ("hotel_id") REFERENCES "hotels" ("id") ON DELETE CASCADE;

ALTER TABLE "reservations" ADD FOREIGN KEY ("payment_info_id") REFERENCES "payments" ("id");

ALTER TABLE "promotions" ADD FOREIGN KEY ("hotel_id") REFERENCES "hotels" ("id") ON DELETE CASCADE;

ALTER TABLE "addresses" ADD FOREIGN KEY ("country_id") REFERENCES "countries" ("id");

ALTER TABLE "reviews" ADD FOREIGN KEY ("hotel_id") REFERENCES "hotels" ("id") ON DELETE CASCADE;

ALTER TABLE "avg_price_history" ADD FOREIGN KEY ("hotel_id") REFERENCES "hotels" ("id") ON DELETE CASCADE;

ALTER TABLE "vacancy_history" ADD FOREIGN KEY ("hotel_id") REFERENCES "hotels" ("id") ON DELETE CASCADE;

ALTER TABLE "hotel_ratings" ADD FOREIGN KEY ("hotel_id") REFERENCES "hotels" ("id") ON DELETE CASCADE;

ALTER TABLE "room_amount" ADD FOREIGN KEY ("hotel_id") REFERENCES "hotels" ("id") ON DELETE CASCADE;

ALTER TABLE "room_amount" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id") ON DELETE CASCADE;

ALTER TABLE "reservation_rooms" ADD FOREIGN KEY ("reservation_id") REFERENCES "reservations" ("id") ON DELETE CASCADE;

ALTER TABLE "reservation_add_ons" ADD FOREIGN KEY ("reservation_id") REFERENCES "reservations" ("id") ON DELETE CASCADE;

ALTER TABLE "reservation_add_ons" ADD FOREIGN KEY ("addon_id") REFERENCES "reservation_addons" ("id") ON DELETE CASCADE;

ALTER TABLE "reservation_addons" ADD FOREIGN KEY ("hotel_id") REFERENCES "hotels" ("id") ON DELETE CASCADE;

ALTER TABLE "reservation_rooms" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id") ON DELETE CASCADE;