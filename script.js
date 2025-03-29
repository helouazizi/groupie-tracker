
async function fetchArtists() {
    try {
        let response = await fetch("http://localhost:8080/");
        let artists = await response.json();
        let grid = document.getElementById("artistsGrid");
        grid.innerHTML = "";
        artists.forEach(artist => {
            let card = document.createElement("div");
            card.classList.add("card");
            card.innerHTML = `
                <img src="${artist.image}" alt="${artist.name}">
                <div class="card-content">
                    <h3>${artist.name}</h3>
                </div>
            `;
            grid.appendChild(card);
        });
    } catch (error) {
        console.error("Error fetching artists:", error);
    }
}
fetchArtists();

function toggleDarkMode() {
    document.body.classList.toggle("dark-mode");
}

function filterArtists() {
    let input = document.getElementById("search").value.toLowerCase();
    let cards = document.querySelectorAll(".card");
    cards.forEach(card => {
        let name = card.querySelector("h3").innerText.toLowerCase();
        card.style.display = name.includes(input) ? "block" : "none";
    });
}
