import { renderArtistDetails,renderError } from "./dom.js";

export async function artistDetails() {
  // const container = document.getElementById("container"); 
  const buttons = document.body.querySelectorAll(".details-btn");
  // console.log(buttons);
  buttons.forEach((btn) => {
    btn.addEventListener("click", async (e) => {
      e.preventDefault();
      const artistId = btn.dataset.id;
      // console.log(artistId);
      try {
        const res = await fetch(`http://localhost:8080/api/artists/details?id=${artistId}`);
        if (!res.ok){
          const error = new Error(res.statusText);
          error.status = res.status;
          throw error;
        }
        const data = await res.json(); // always parse JSON
        renderArtistDetails(data);
      } catch (err) {
        renderError(err)
      }
    });
  });
}





