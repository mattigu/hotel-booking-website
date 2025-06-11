# Backend Queries
1. GET: /get/hotels?city=[city_name]&startdate=[yyyy-mm-dd]&enddate=[yyyy-mm-dd]&guests=[number]

2. GET: /get/hotel/[id]

3. POST: /reserve/room  
in data: [room_id], [customer_id], [start_date], [end_date], [payment_id]

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

4. POST: /post/opinion
in data: [hotel_id], [username], [star_rating], ["content"]