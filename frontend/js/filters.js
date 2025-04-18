import { renderArtists, touchedInputs } from './dom.js'
import { renderError } from './error.js'


export function setupFilters() {
  const form = document.getElementById('filter-form')
  
  form.addEventListener('submit', async (e) => {
    e.preventDefault()

    const filters = {
      creationDateFrom: touchedInputs.has("creation-from")
        ? form.querySelector('input[name="creation-from"]').value
        : "0",

      creationDateTo: touchedInputs.has("creation-to")
        ? form.querySelector('input[name="creation-to"]').value
        : "0",

      firstAlbumFrom: touchedInputs.has("album-from")
        ? form.querySelector('input[name="album-from"]').value
        : "0",

      firstAlbumTo: touchedInputs.has("album-to")
        ? form.querySelector('input[name="album-to"]').value
        : "0",

      members: form.querySelector('input[name="members"]:checked')?.value || "0",

      concertDates: form.querySelector('#location-input')?.value.trim() || "",
    };

    try {
      const res = await fetch('http://localhost:8080/api/artists/filter', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(filters)
      })
      const data = await res.json()
      console.log(res.status);
      
      console.log(data);


      if (!res.ok) {
        throw {
          status: res.status,
          message: data.message || "Server error",
          details: data.details || ""
        }

      }
      renderArtists(filteredArtists)
    } catch (err) {
      // network failure or invalid JSON
      renderError({
        status: err.status || "Network Error",
        message: err.message || "Something went wrong",
        details: err.details || ""
      })
    }
  })
}
