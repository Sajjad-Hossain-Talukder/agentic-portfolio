// src/services/apiService.ts

interface PostOptions {
  [key: string]: any;
}

export async function post<T>(endpoint: string, body: PostOptions): Promise<T> {
  const BASE_URL = import.meta.env.VITE_BASE_URL;
  const API_KEY =  import.meta.env.VITE_API_KEY;

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
