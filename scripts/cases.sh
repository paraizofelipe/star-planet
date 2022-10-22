#!/bin/bash

BASE_URL="http://localhost:3000/api"

function load_planet {
	curl --silent -X POST "$BASE_URL/planets/load/$1" | jq "."
}

function remove_planet {
	curl --silent -X DELETE "$BASE_URL/planets/id/$1" | jq "."
}

function fetch_planet_by_id {
	curl --silent -X GET "$BASE_URL/planets/id/$1" | jq "."
}

function fetch_planet_by_name {
	curl --silent -X GET "$BASE_URL/planets/name/$1" | jq "."
}

function list_planet {
	curl --silent -X GET "$BASE_URL/planets" | jq "."
}

planets=(1 2 3 4 5 6 7 8 9 10)

time for planet_id in "${planets[@]}"; do
	time load_planet "$planet_id"
done

time list_planet

time for planet_id in "${planets[@]}"; do
	time fetch_planet_by_id "$planet_id"
done

time for planet_id in "${planets[@]}"; do
	time remove_planet "$planet_id"
done
