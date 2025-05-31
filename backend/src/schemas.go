package main

type testStruct struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type hotelOverview struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Desc string `json:"description"`
	Star int64 `json:"star_standard"`
}

type hotelInfo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Desc string `json:"description"`
	Star int64 `json:"star_standard"`
}