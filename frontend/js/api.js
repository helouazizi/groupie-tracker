import { renderArtists, renderError } from "./dom.js"
import { artistDetails } from "./details.js";


export async function fetchArtists() {
  try {
    const res = await fetch('http://localhost:8080/api/artists')
    if (!res.ok) {
      const error = new Error(res.statusText);
      error.status = res.status;
      throw error;
    }
    const data = await res.json()
    renderArtists(data)
    artistDetails()
  } catch (err) {
    renderError(err)
  }
}
