import { renderError } from "./dom.js"

export async function fetchArtists() {
  try {
    const res = await fetch('http://localhost:8080/api/artists')
    // At this point, we are sure it's a valid response
    const data = await res.json()
    return data
  } catch (err) {
    renderError(err)
    return []
  }
}
