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
    "payment_info_id": 1
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
    "rating": 4
}
```