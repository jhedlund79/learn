import axios from 'axios';

const API_KEY = 'bae3c76b4d0c4a0fce93c5b275d9fb92';
const ROOT_URL = `http://api.openweathermap.org/data/2.5/forecast?appid=${API_KEY}`;

export const FETCH_WEATHER = 'FETCH_WEATHER';

export function fetchWeather(city) {
  const url = `${ROOT_URL}&q=${city},us`;
  const request = axios.get(url);
  console.log('Request:', request);
  return {
    type: FETCH_WEATHER,
    payload: request,
  };
}
