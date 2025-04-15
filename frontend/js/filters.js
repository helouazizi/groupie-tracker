import { renderArtists ,touchedInputs } from './dom.js'


export function setupFilters() {
  const form = document.getElementById('filter-form')

  form.addEventListener('submit', async (e) => {
    e.preventDefault()

    //const formData = new FormData(form)

    // Extract and convert filter values
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
    


    console.log(filters);

    try {
      const res = await fetch('http://localhost:8080/api/artists', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(filters)
      })

      if (!res.ok) throw new Error('Failed to fetch filtered artists.')

      const filteredArtists = await res.json()
      renderArtists(filteredArtists)
    } catch (err) {
      console.error('Error:', err)
    }
  })
}
