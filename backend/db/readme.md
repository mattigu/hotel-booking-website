# Projekt BD2

### Model ER
![ER model](/backend/db/images/BD2%20-%20ER%20diagram.png)

### Opis modelu ER

**Encja:** `Hotel`

| *Atrybut*     | *Opis*                                   |
| ------------- | ---------------------------------------- |
| Name          | Nazwa hotelu.                            |
| Star Standard | Poziom jakości hotelu, np. 3-gwiazdkowy. |
| Description   | Krótki opis hotelu.                      |

**Encja:** `Address`
| *Atrybut*    | *Opis*                               |
| ------------ | ------------------------------------ |
| Street       | Ulica, na której znajduje się hotel. |
| House Number | Numer budynku.                       |
| Zip Code     | Kod pocztowy.                        |
| City         | Miasto.                              |
| Country      | Kraj.                                |

**Encja:** `Promotion`
| *Atrybut*  | *Opis*                             |
| ---------- | ---------------------------------- |
| Discount   | Wysokość zniżki w ramach promocji. |
| Date Start | Data rozpoczęcia promocji.         |
| Date End   | Data zakończenia promocji.         |

**Encja:** `Review`
| *Atrybut*   | *Opis*                        |
| ----------- | ----------------------------- |
| Author Name | Imię autora recenzji.         |
| Rating      | Ocena przyznana przez autora. |
| Review Text | Treść recenzji.               |

**Encja:** `Room`
| *Atrybut*         | *Opis*                              |
| ----------------- | ----------------------------------- |
| Room Number       | Numer pokoju.                       |
| Single Bed Amount | Liczba pojedynczych łóżek w pokoju. |
| Double Bed Amount | Liczba podwójnych łóżek w pokoju.   |

**Encja:** `Amenity`
| *Atrybut*   | *Opis*              |
| ----------- | ------------------- |
| Name        | Nazwa udogodnienia. |
| Description | Opis udogodnienia.  |
| Price       | Cena udogodnienia.  |

**Encja:** `Reservation`
| *Atrybut*  | *Opis*                       |
| ---------- | ---------------------------- |
| Start Date | Data rozpoczęcia rezerwacji. |
| End Date   | Data zakończenia rezerwacji. |

**Encja:** `Payment`
| *Atrybut* | *Opis*                                     |
| --------- | ------------------------------------------ |
| Amount    | Kwota płatności.                           |
| Due Date  | Termin zapłaty.                            |
| Type      | Typ płatności (np. karta, przelew).        |
| Fulfilled | Informacja, czy płatność została wykonana. |

**Encja:** `Customer`
| *Atrybut*    | *Opis*                  |
| ------------ | ----------------------- |
| Name         | Imię klienta.           |
| Surname      | Nazwisko klienta.       |
| Email        | Adres e-mail klienta.   |
| Phone Number | Numer telefonu klienta. |

### Model logiczny
![Logical model](/backend/db/images/BD2_model_relacyjny.png)

**Tabela:** `hotels` \
Zawiera podstawowe informacje dotyczące kżdego hotelu.

| *Kolumna*    | *Typ danych*            | *Ograniczenia* | *Opis* |
| ------------ | ----------------------- | -------------- | ------ |
| id         | SERIAL        | not null, unique | Unikalny identyfikator hotelu. |
| address_id         | integer        | not null | Klucz obcy odnoszący się do adresu hotelu. |
| name         | text        | not null | Nazwa hotelu.|
| description         | text        |  | Opis hotelu. |
| star\_standard         | integer        | not null, integer | Standard opisujący hotel, np. 4-gwiazdkowy. |

**Tabela:** `addresses` \
Zawiera wszystkie informacje dotyczące adresu danej lokalizacji.

| *Kolumna*     | *Typ danych* | *Ograniczenia*   | *Opis*                             |
| ------------- | ------------ | ---------------- | ---------------------------------- |
| id            | SERIAL       | not null, unique | Unikalny identyfikator adresu.     |
| city          | text         | not null         | Miasto.                            |
| street        | text         | not null         | Ulica.                             |
| zip\_code     | text         | not null         | Kod pocztowy.                      |
| house\_number | integer      | not null         | Numer domu.                        |
| country\_id   | integer      | not null         | Klucz obcy odnoszący się do kraju. |

**Tabela:** `countries` \
Tabela słownikowa, zawiera nazwy krajów użyte w adresach.

| *Kolumna* | *Typ danych* | *Ograniczenia*   | *Opis*                        |
| --------- | ------------ | ---------------- | ----------------------------- |
| id        | SERIAL       | not null, unique | Unikalny identyfikator kraju. |
| name      | text         | not null         | Nazwa kraju.                  |

**Tabela:** `rooms` \
Opisuje kluczowe cechy pokoju, wykorzystywane w złożonoych zapytaniach o np. dostępność miejsca w czasie. 

| *Kolumna*        | *Typ danych* | *Ograniczenia*   | *Opis*                         |
| ---------------- | ------------ | ---------------- | ------------------------------ |
| id               | SERIAL       | not null, unique | Unikalny identyfikator pokoju. |
| hotel\_id        | integer      | not null         | Klucz obcy do hotelu.          |
| room\_number     | integer      | not null         | Numer pokoju w hotelu.         |
| single\_bed\_num | integer      | not null         | Liczba pojedynczych łóżek.     |
| double\_bed\_num | integer      | not null         | Liczba podwójnych łóżek.       |
| base\_price      | integer      | not null         | Cena wynajęcia pokoju na 1 noc.    |

**Tabela:** `reservations` \
Zawiera dane o rezerwacjach unikatowe dla tej jednej rezerwacji oraz potrzebne w złożonych zapytaniach bazodanowych.

| *Kolumna*         | *Typ danych* | *Ograniczenia*   | *Opis*                             |
| ----------------- | ------------ | ---------------- | ---------------------------------- |
| id                | SERIAL       | not null, unique | Unikalny identyfikator rezerwacji. |
| customer\_id      | integer      | not null         | Klucz obcy do danych klienta.             |
| hotel\_id         | integer      | not null         | Klucz obcy do hotelu.              |
| start\_date       | date         | not null         | Data rozpoczęcia rezerwacji.       |
| end\_date         | date         | not null         | Data zakończenia rezerwacji.       |
| payment\_info\_id | integer      | not null         | Klucz obcy do płatności.           |

**Tabela:** `room_amount` \
TODO: ???

| *Kolumna* | *Typ danych* | *Ograniczenia* | *Opis*                                              |
| --------- | ------------ | -------------- | --------------------------------------------------- |
| hotel\_id | integer      | not null       | Klucz obcy do tabeli `hotels`.                      |
| room\_id  | integer      | not null       | Klucz obcy do tabeli `rooms`.                       |
| amount    | integer      | not null       | Liczba dostępnych pokoi danego typu w danym hotelu. |

**Tabela:** `promotions` \
Zawiera dane o promocjach oferowanych przez hotele – daty trwania i wartości rabatu.

| *Kolumna*      | *Typ danych* | *Ograniczenia*   | *Opis*                                          |
| -------------- | ------------ | ---------------- | ----------------------------------------------- |
| id             | SERIAL       | not null, unique | Unikalny identyfikator promocji.                |
| hotel\_id      | integer      | not null         | Klucz obcy do hotelu, którego dotyczy promocja. |
| start\_date    | date         | not null         | Data rozpoczęcia promocji.                      |
| end\_date      | date         | not null         | Data zakończenia promocji.                      |
| discount\_flat | integer      |                  | Kwotowy rabat (np. 50 zł).                      |
| discount\_pct  | integer      |                  | Procentowy rabat (np. 10 %).                    |


**Tabela:** `hotel_ratings` \
Tabela agregująca oceny hoteli, używana do liczenia średniej oceny.

| *Kolumna*       | *Typ danych* | *Ograniczenia* | *Opis*    |
| --------------- | ------------ | -------------- | ---------------- |
| hotel\_id       | integer      | not null       | Klucz główny i jednocześnie obcy do `hotels`.                    |
| current\_rating | numeric(3,2) | not null       | Średnia ocena (np. 4,25).  |
| review\_count   | integer      | not null       | Łączna liczba recenzji wziętych pod uwagę przy obliczaniu oceny. |


**Tabela:** `reviews` \
Zawiera recenzje klientów dotyczące konkretnych hoteli. 

| *Kolumna*    | *Typ danych* | *Ograniczenia*   | *Opis*     |
| ------------ | ------------ | ---------------- | ------------- |
| id           | SERIAL       | not null, unique | Unikalny identyfikator recenzji.       |
| username     | text         | not null         | Nazwa użytkownika dodającego recenzję. |
| hotel\_id    | integer      | not null         | Klucz obcy do ocenianego hotelu.       |
| rating       | integer      | not null         | Ocena liczbowo (w skali 1–5).              |
| review\_text | text         | not null         | Treść recenzji.                        |
| upload\_date | date         | not null         | Data dodania recenzji.                 |


**Tabela:** `avg_price_history` \
Tabela analityczna, rejestruje średnią cenę pokoi hotelowych w określonych przedziałach czasowych.

| *Kolumna*     | *Typ danych* | *Ograniczenia*   | *Opis*                                             |
| ------------- | ------------ | ---------------- | -------------------------------------------------- |
| id            | SERIAL       | not null, unique | Unikalny identyfikator rekordu.                    |
| hotel\_id     | integer      | not null         | Klucz obcy do hotelu.                              |
| period\_start | date         | not null         | Początek analizowanego okresu.                     |
| period\_end   | date         | not null         | Koniec analizowanego okresu.                       |
| avg\_price    | integer      | not null         | Średnia cena pokoju w tym okresie (np. w złotych). |

**Tabela:** `vacancy_history` \
Tabela analityczna, zawiera dane o ilości wolnych pokoi w danym hoteli w określonych przedziałach czasowych.

| *Kolumna*     | *Typ danych* | *Ograniczenia*   | *Opis*                                           |
| ------------- | ------------ | ---------------- | ------------------------------------------------ |
| id            | SERIAL       | not null, unique | Unikalny identyfikator rekordu.                  |
| hotel\_id     | integer      | not null         | Klucz obcy do hotelu.                            |
| period\_start | date         | not null         | Początek analizowanego okresu.                   |
| period\_end   | date         | not null         | Koniec analizowanego okresu.                     |
| vacancies     | integer      | not null         | Średnia liczba wolnych pokoi w danym przedziale. |

**Tabela:** `amenities` \
Zawiera listę możliwych udogodnień hotelu/pokoju.

| *Kolumna*   | *Typ danych* | *Ograniczenia*   | *Opis*                               |
| ----------- | ------------ | ---------------- | ------------------------------------ |
| id          | SERIAL       | not null, unique | Unikalny identyfikator udogodnienia. |
| name        | text         | not null         | Nazwa udogodnienia. |
| description | text         | not null         | Opis udogodnienia.    |


**Tabela:** `available_addons` \
Tabela zawierająca usługi oferowane przez hotel za dodatkową opłatą.

| *Kolumna* | *Typ danych* | *Ograniczenia*   | *Opis*             |
| --------- | ------------ | ---------------- | ------------------- |
| id        | integer      | not null, unique | Unikalny identyfikator dodatku (definiowanego przez hotel). |
| hotel\_id | integer      | not null         | Klucz obcy do hotelu, który oferuje dodatek.        |
| name      | text         | not null         | Nazwa dodatku. |
| price     | integer      | not null         | Cena dodatku.   |

**Tabela:** `customers` \
Przechowuje dane osobowe klientów.

| *Kolumna*     | *Typ danych* | *Ograniczenia*   | *Opis*     |
| ------------- | ------------ | ---------------- | -------- |
| id            | SERIAL       | not null, unique | Unikalny identyfikator klienta. |
| name          | text         | not null         | Imię klienta.                   |
| surname       | text         | not null         | Nazwisko klienta.               |
| phone\_number | text         | not null         | Numer telefonu klienta.                 |
| email         | text         |                  | Adres e‑mail klienta.      |

**Tabela:** `payments` \
Zawiera szczegóły dotyczące płatności.

| *Kolumna*     | *Typ danych* | *Ograniczenia*   | *Opis*                                            |
| ------------- | ------------ | ---------------- | ------------------------------------------------- |
| id            | SERIAL       | not null, unique | Unikalny identyfikator płatności.                 |
| payment\_type | text         | not null         | Rodzaj płatności (np. karta, przelew, gotówka).   |
| payment\_data | text         |                  | Dodatkowe informacje zależne od rodzaju płatności, np. nr karty. |
| due\_date     | date         | not null         | Termin uregulowania płatności.                    |
| amount        | text         | not null         | Kwota do zapłaty.                                 |
| fulfilled     | boolean      | not null          | Informacja o tym, czy płatność została zrealizowana. |

**Tabela:** `hotel_amenities` \
Tabela łącząca hotele z dostępnymi w nich udogodnieniami.

| *Kolumna*   | *Typ danych* | *Ograniczenia* | *Opis*                     |
| ----------- | ------------ | -------------- | -------------------------- |
| hotel\_id   | integer      | not null       | Klucz obcy do `hotels`.    |
| amenity\_id | integer      | not null       | Klucz obcy do `amenities`. |

**Tabela:** `room_amenities` \
Tabela łącząca pokoje z dostępnymi w nich udogodnieniami.

| *Kolumna*   | *Typ danych* | *Ograniczenia* | *Opis*                     |
| ----------- | ------------ | -------------- | -------------------------- |
| room\_id    | integer      | not null       | Klucz obcy do `rooms`.     |
| amenity\_id | integer      | not null       | Klucz obcy do `amenities`. |

**Tabela:** `selected_addons` \
Tabela łącząca rezerwację z wybranymi dodatkowo usługami hotelu.

| *Kolumna*       | *Typ danych* | *Ograniczenia* | *Opis*               |
| --------------- | ------------ | -------------- | ------------------ |
| reservation\_id | integer      | not null       | Klucz obcy do `reservations`.                         |
| addon\_id       | integer      | not null       | Klucz obcy do `reservation_addons`. |


**Tabela:** `reservation_rooms` \
Tabela łącząca rezerwację z zajętymi przez nią pokojami.

| *Kolumna*       | *Typ danych* | *Ograniczenia* | *Opis*        |
| --------------- | ------------ | -------------- | --------- |
| reservation\_id | integer      | not null       | Klucz obcy do `reservations`. |
| room\_id        | integer      | not null       | Klucz obcy do `rooms`.        |

