# Backend Queries


1. GET: /get/hotels?city=[city_name]&startdate=[yyyy-mm-dd]&enddate=[yyyy-mm-dd]&guests=[number]

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
2. GET: /get/hotel/[id]

przykładowe zapytanie:

`localhost:3000/get/hotel/1`

zwraca:
```json
{
    "id":1,
    "name":"Hotel Royal",
    "adress_id":1,
    "description":"Luksusowy hotel w centrum Warszawy.",
    "star_standard":5
}
```
-------------------
3. POST: /reserve/room  

przykładowe dane w zapytaniu:
```json
{
    "hotel_id":1, 
    "room_id": 1, 
    "start_date":"2025-05-24", 
    "end_date":"2025-05-26", 
    "customer_id": 1, 
    "payment_info_id": 1
}
```

----------------------
(jeszcze niezaimplementowane)
4. POST: /post/opinion
in data: [hotel_id], [username], [star_rating], ["content"]