import { renderArtists, touchedInputs, renderError } from './dom.js';

export function setupFilters() {
  const form = document.getElementById('filter-form');

  form.addEventListener('submit', (e) => {
    e.preventDefault();

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

    fetch('http://localhost:8080/api/artists/filter', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(filters),
    })
      .then(res => {
        if (!res.ok) {
          console.log("error");
          throw new Error(`HTTP Error ${res.status}: ${res.statusText}`);
        }
        return res.json();
      })
      .then(data => {
        renderArtists(data);
      })
      .catch(err => {
        renderError(err.message);
      });
  });
}
