import { renderArtists } from './dom.js'

export function setupFilters(data) {
  const form = document.getElementById('filter-form')

  form.addEventListener('submit', async (e) => {
    e.preventDefault()

    const formData = new FormData(form)

    // Extract and convert filter values
    const filters = {
      creationDateFrom: parseInt(form.querySelector('input[name="creation-from"]')?.value || 0),
      creationDateTo: parseInt(form.querySelector('input[name="creation-to"]')?.value || 0),
      firstAlbumFrom: parseInt(form.querySelector('input[name="album-from"]')?.value || 0),
      firstAlbumTo: parseInt(form.querySelector('input[name="album-to"]')?.value || 0),
      members: parseInt(form.querySelector('input[name="members"]:checked')?.value || 0),
      concertDates: form.querySelector('#location-input')?.value.trim() || "",
    }
    

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
