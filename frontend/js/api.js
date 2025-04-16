
import { renderArtists, renderError } from "./dom.js"
export async function fetchArtists() {
  try {
    const res = await fetch('http://localhost:8080/api/artists')
    if (!res.ok) throw new Error('Failed to fetch')
    const data = await res.json()

    if (!res.ok) {
      // render backend error with status
      renderError({
        status: res.status,
        message: data.message || "Unknown error",
        details: data.details || ""
      });
      return;
    }
    renderArtists(data);
  } catch (err) {
    // network failure or invalid JSON
    renderError({
      status: "Network Error",
      message: "Failed to connect to server.",
      details: err.message
    });

  }
}