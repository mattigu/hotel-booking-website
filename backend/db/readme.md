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

