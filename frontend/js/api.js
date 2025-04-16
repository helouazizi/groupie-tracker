
import { renderError } from "./error.js"
export async function fetchArtists() {
  try {
    const res = await fetch('http://localhost:8080/api/artists')
    if (!res.ok) throw new Error('Failed to fetch')
    return await res.json()
  } catch (err) {
    // network failure or invalid JSON
    renderError({
      status: "Network Error",
      message: "Failed to connect to server.",
      details: err.message
    });
    return []
  }
}