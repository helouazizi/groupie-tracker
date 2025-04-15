export function renderArtists(artists) {
  const container = document.getElementById("container");
  container.innerHTML = "";
  if (!artists) {
    container.innerHTML = "No Artist Found"
    return
  }
  artists.forEach((artist) => {
    const card = document.createElement("div");
    card.className = "card";
    card.innerHTML = `
        <img src="${artist.image}" alt="${artist.name}">
        <h2>${artist.name}</h2>
      `;
    container.appendChild(card);
  });
}


export function updateRangeValues() {
  const ranges = document.querySelectorAll('input[type="range"]')
  ranges.forEach((range) => {
    const label = document.getElementById(`${range.id}-value`)
    if (label) {
      // Set initial value
      label.textContent = range.value

      // Update on input
      range.addEventListener('input', () => {
        label.textContent = range.value
      })
    }
  })
}

