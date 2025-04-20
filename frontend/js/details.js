import { renderArtistDetails, renderError } from "./dom.js";


export function artistDetails() {
  const buttons = document.body.querySelectorAll(".details-btn");
  // console.log(buttons);

  buttons.forEach((btn) => {
    btn.addEventListener("click", (e) => {
      e.preventDefault();
      const artistId = btn.dataset.id;

      fetch(`http://localhost:8080/api/artists/details?id=${artistId}`)
        .then(res => {
          if (!res.ok) {
            error = true
            const error = new Error(res.statusText);
            error.status = res.status;
            throw error;
          }
          return res.json();
        })
        .then(data => {
          renderArtistDetails(data);
        })
        .catch(err => {
          renderError(err);
        });
    });
  });
}
