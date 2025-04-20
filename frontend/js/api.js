import { renderArtists, renderError } from "./dom.js"
import { artistDetails } from "./details.js";

export function fetchArtists() {
  fetch('http://localhost:8080/api/artists')
    .then(res => {
      if (!res.ok) {
        const error = new Error(res.statusText);
        error.status = res.status;
        throw error;
      }
      return res.json();
    })
    .then(data => {
      renderArtists(data);
      artistDetails();
    })
    .catch(err => {
      renderError(err);
    });
}
