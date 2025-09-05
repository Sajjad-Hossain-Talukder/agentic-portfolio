// src/services/apiService.ts
const BASE_URL = " http://160.191.162.38:8082/api/v1"; // replace with your backend IP
const API_KEY = "sht-RX-AP-G7k2PqX9Lm3Z"; // ideally, store in .env and use Vite/CRA env variables

interface PostOptions {
  [key: string]: any;
}

export async function post<T>(endpoint: string, body: PostOptions): Promise<T> {
  const response = await fetch(`${BASE_URL}/${endpoint}`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "X-API-Key": API_KEY
    },
    body: JSON.stringify(body)
  });

  // console.log("res >> ", response)

  if (!response.ok) {
    const error = await response.json();
    throw new Error(error.error || "API request failed");
  }

  return response.json() as Promise<T>;
}
