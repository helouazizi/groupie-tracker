export function renderArtists(artists) {
  const container = document.getElementById("container");
  container.innerHTML = "";

  artists.forEach((artist) => {
    const card = document.createElement("div");
    card.className = "card";
    card.innerHTML = `
        <h2>${artist.name}</h2>
        <img src="${artist.image}" alt="${artist.name}">
      `;
    container.appendChild(card);
  });
}
