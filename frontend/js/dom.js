import { loadHomePage } from "./main.js";

export function renderArtists(artists) {
  const container = document.getElementById("container");
  container.innerHTML = "";

  if (!artists || artists.length === 0) {
    container.innerHTML = "No Artist Found";
    return;
  }

  artists.forEach((artist) => {
    const card = document.createElement("div");
    card.className = "card";

    card.innerHTML = `
      <img src="${artist.image}" alt="${artist.name}">
      <h2>${artist.name}</h2>
      <button class="details-btn" data-id="${artist.id}">Details</button>
    `;

    container.appendChild(card);
  });
}

const touchedInputs = new Set();
export{ touchedInputs}
export function updateRangeValues() {
  const ranges = document.querySelectorAll('input[type="range"]');

  ranges.forEach((range) => {
    const label = document.getElementById(`${range.id}-value`);
    if (!label) return;

    // Initial value display
    label.textContent = range.value;

    // Update display + mark as touched
    range.addEventListener('input', () => {
      label.textContent = range.value;
      touchedInputs.add(range.name);
    });
  });
}


export function renderError(err) {
  document.body.innerHTML = "";

  const errorContent = document.createElement("div");
  errorContent.classList.add("error-box");

  errorContent.innerHTML = `
    <h1>Oooops 😬</h1>
    <p><strong>${err.status || "Error"}</strong> | ${err.message || "Something went wrong."}</p>
    ${err.details ? `<pre>${err.details}</pre>` : ""}
    <button class = "home-btn" id="back-home-btn">Back Home</button>
  `;

  document.body.appendChild(errorContent);
  document.getElementById("back-home-btn").addEventListener("click", loadHomePage);
}




