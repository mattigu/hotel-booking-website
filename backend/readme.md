# Backend Queries


### 1. GET: /get/hotels?city=[city_name]&startdate=[yyyy-mm-dd]&enddate=[yyyy-mm-dd]&guests=[number]

(chwilowo sprawdza tylko miasto, i ilość gości bez innych danych)

przykładowe zpaytanie:

`localhost:3000/get/hotels?city=Warszawa&startdate=2025-05-24&enddate=2025-05-24&guests=3`

zwraca:

```json
[
    {
        "id":1,
        "name":"Hotel Royal",
        "price":"100.99",
        "star_standard":5,
        "sigle_beds":1,
        "double_beds":1,
        "photo_url":"https://photo_site.net/photo.jpg"
    }
]
```
------------------
### 2. GET: /get/hotel/[id]

przykładowe zapytanie:

`localhost:3000/get/hotel/1`

zwraca:
```json
{
    "id":1,
    "name":"Hotel Royal",
    "photo_url":"https://photo_site.net/photo.jpg",
    "star_standard":5,
    "avg_rating":4.2,
    "description":"Luksusowy hotel w centrum Warszawy.",
    "address":
    {
        "city":"Warszawa",
        "street":"Marszałkowska",
        "house_number":"10",
        "country":"Polska"
    },
    "amenities":
    [
        {
            "id":1,
            "name":"Basen",
            "description":"Basen dostępny dla gości."
        },
        {
            "id":2,
            "name":"Siłownia",
            "description":"W pełni wyposażona siłownia."
        }
    ],
    "reviews":
    [
        {
            "username":"jan_kowalski",
            "review_text":"Niesamowite doświadczenie!",
            "rating":5
        }
    ]
}

```
-------------------
### 3. POST: /reserve/room  

przykładowe dane w zapytaniu:
```json
{
    "room_id": 2, 
    "hotel_id":1, 
    "start_date":"2025-05-24", 
    "end_date":"2025-05-26", 
    "customer":
    {
        "name": "Andrzej",
        "surname": "Andrzejowski",
        "phone_number": "999999999"
    }, 
    "payment_info":
    {
        "payment_type": "przelew",
        "payment_data": "20923423409234234234",
        "amount": "330.72"
    }
}
```

----------------------

### 4. POST: /post/opinion
in data: [hotel_id], [username], [star_rating], ["content"]

przykładowe dane w zapytaniu:

```json
{
    "hotel_id": 1,
    "username": "user1",
    "review_text": "super hotel polecam gorąco.",
    "rating": 4,
}
```

------------------------

### 5. GET: /get/reservations?name=[name]&surname=[surname]&phone_number=[phone_number]

przykładowe zapytanie:
`localhost:3000/get/reservations?name=Jan&surname=Kowalski&phone_number=123456789`

odpowiedź:

```json
[
    {
        "hotel_id":1,
        "room_id":1,
        "start_date":"2025-05-01",
        "end_date":"2025-05-05"
    }
]

```

--------------------

### 6. /get/configuration?hotel_id=[id]&guests=[guests]

przykładowe zapytanie:
`localhost:3000/get/configuration?hotel_id=1&guests=2`

odpowiedź:

```json
[
    {
        "single_beds":1,
        "double_beds":1,
        "price":350
    },
    {
        "single_beds":2,
        "double_beds":0,
        "price":280
    }
]
```

--------------------------

### 7. /get/addons?hotel_id=[id]

przykładowe zapytanie:
`localhost:3000/get/addons?hotel_id=1`

odpowiedź:
```json
[
    {
        "id":1,
        "name":"addon1",
        "price":100
    }
]
```