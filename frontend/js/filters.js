import { renderArtists, touchedInputs , renderError } from './dom.js'


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
      const res = await fetch('http://localhost:8080/api/artists', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(filters)
      })
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
      renderArtists(filteredArtists)
    } catch (err) {
      // network failure or invalid JSON
      renderError({
        status: "Network Error",
        message: "Failed to connect to server.",
        details: err.message
      });
    }
  })
}
