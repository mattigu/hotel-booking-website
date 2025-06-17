-- Dane dla tabeli countries
INSERT INTO "countries" ("name") VALUES
('Polska'),
('Niemcy'),
('Francja'),
('Włochy'),
('Hiszpania');

-- Dane dla tabeli addresses
INSERT INTO "addresses" ("city", "street", "zip_code", "house_number", "country_id") VALUES
('Warszawa', 'Marszałkowska', '00-001', 10, 1),
('Kraków', 'Floriańska', '31-001', 5, 1),
('Berlin', 'Unter den Linden', '10117', 15, 2),
('Paryż', 'Champs-Élysées', '75008', 20, 3),
('Madryt', 'Gran Via', '28013', 25, 5);

-- Dane dla tabeli hotels
INSERT INTO "hotels" ("address_id", "name", "description", "star_standard") VALUES
(1, 'Hotel Royal', 'Luksusowy hotel w centrum Warszawy.', 5),
(2, 'Hotel Prestige', 'Nowoczesny hotel w sercu Krakowa.', 4),
(3, 'Hotel Business', 'Idealny dla podróżujących służbowo.', 4),
(4, 'Hotel Elegance', 'Urokliwy hotel butikowy w Paryżu.', 5),
(5, 'Hotel Panorama', 'Hotel z widokiem na centrum Madrytu.', 5);

-- Dane dla tabeli hotel_amenity_types
INSERT INTO "hotel_amenity_types" ("name", "description") VALUES
('Basen', 'Basen dostępny dla gości.'),
('Siłownia', 'W pełni wyposażona siłownia.'),
('Spa', 'Relaksujące usługi spa.'),
('Parking', 'Parking na miejscu.'),
('WiFi', 'Darmowy szybki internet.');

-- Dane dla tabeli hotel_amenities
INSERT INTO "hotel_amenities" ("hotel_id", "hotel_amenity_type", "price") VALUES
(1, 1, 0),
(1, 2, 0),
(2, 3, 100),
(3, 4, 0),
(4, 5, 0);

-- Dane dla tabeli rooms
INSERT INTO "rooms" ("hotel_id", "room_number", "single_bed_num", "double_bed_num", "base_price") VALUES
(1, 101, 1, 1, 350),
(1, 102, 2, 0, 280),
(2, 201, 1, 1, 220),
(3, 301, 0, 2, 450),
(4, 401, 1, 1, 550);

-- Dane dla tabeli room_amenity_types
INSERT INTO "room_amenity_types" ("name", "description") VALUES
('Telewizor', 'Telewizor'),
('Mini Bar', 'W pełni zaopatrzony mini bar.'),
('Lodówka', 'Lodówka w pokoju.'),
('Klimatyzacja', 'Pokój z klimatyzacją.'),
('Balkon', 'Pokój z prywatnym balkonem.'),
('Ekspres do kawy', 'Ekspres do kawy w pokoju.');

-- Dane dla tabeli room_amenities
INSERT INTO "room_amenities" ("room_id", "room_amenitiy", "price") VALUES
(1, 1, 0),
(1, 2, 20),
(2, 3, 0),
(3, 4, 30),
(4, 5, 10);

-- Dane dla tabeli customers
INSERT INTO "customers" ("name", "surname", "phone_number", "email") VALUES
('Jan', 'Kowalski', '123456789', 'jan.kowalski@example.com'),
('Anna', 'Nowak', '987654321', 'anna.nowak@example.com'),
('Piotr', 'Wiśniewski', '555666777', 'piotr.wisniewski@example.com'),
('Maria', 'Zielińska', '444555666', 'maria.zielinska@example.com'),
('Krzysztof', 'Kamiński', '333444555', 'krzysztof.kaminski@example.com');

-- Dane dla tabeli payments
INSERT INTO "payments" ("id", "payment_type", "due_date", "amount", "fulfilled") VALUES
(1, 'Karta Kredytowa', '2025-05-01', 2000, false),
(2, 'Przelew Bankowy', '2025-05-8', 1500, true),
(3, 'PayPal', '2025-05-20', 1800, false);

-- Dane dla tabeli reservations
INSERT INTO "reservations" ("id", "customer_id", "hotel_id", "room_id", "start_date", "end_date", "payment_info_id") VALUES
(1, 1, 1, 1, '2025-05-01', '2025-05-05', 1),
(2, 2, 2, 2, '2025-05-8', '2025-05-15', 2),
(3, 3, 3, 3, '2025-05-20', '2025-05-25', 3);


-- Dane dla tabeli promotions
INSERT INTO "promotions" ("hotel_id", "start_date", "end_date", "discount_flat", "discount_pct") VALUES
(1, '2025-05-01', '2025-05-8', NULL, 10),
(2, '2025-05-15', '2025-05-20', 100, NULL);

-- Dane dla tabeli reviews
INSERT INTO "reviews" ("username", "hotel_id", "rating", "review_text", "upload_date") VALUES
('jan_kowalski', 1, 5, 'Niesamowite doświadczenie!', '2025-05-01'),
('anna_nowak', 2, 4, 'Bardzo wygodny pobyt.', '2025-05-01'),
('piotr_wisniewski', 3, 3, 'Dobrze, ale mogło być lepiej.', '2025-05-01');

-- Dane dla tabeli avg_price_history
INSERT INTO "avg_price_history" ("hotel_id", "period_start", "period_end", "avg_price") VALUES
(1, '2024-01-01', '2024-06-30', 500),
(2, '2024-01-01', '2024-06-30', 400);

-- Dane dla tabeli vacancy_history
INSERT INTO "vacancy_history" ("hotel_id", "period_start", "period_end", "vacancies") VALUES
(1, '2024-01-01', '2024-06-30', 20),
(2, '2024-01-01', '2024-06-30', 15);

-- Dane dla tabeli reservation_addons
INSERT INTO "reservation_addons" ("id", "name", "price") VALUES
(1, 'Śniadanie', 45),
(2, 'Parking', 25),
(3, 'Dostęp do SPA', 80),
(4, 'Dostęp do siłowni', 30),
(5, 'Wi-Fi Premium', 15);

INSERT INTO reservation_addons VALUES(1, 'addon1', 100, 1);