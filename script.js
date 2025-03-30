
async function fetchArtists() {
    try {
        let response = await fetch("http://localhost:8080/");
        let artists = await response.json();
        let grid = document.getElementById("artistsGrid");
        grid.innerHTML = "";
        artists.forEach(artist => {
            let card = document.createElement("div");
            card.classList.add("card");
            //card.setAttribute("id",artist.id)
            card.innerHTML = `
                <img src="${artist.image}" alt="${artist.name}">
                <div class="card-content">
                    <h3>${artist.name}</h3>
                </div>
                <button onclick="artistDeatils(${artist.id})">Details</button>
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


async function artistDeatils(id) {
    try {
        let response = await fetch(`http://localhost:8080/artist?id=${id}`);
        let data = await response.json();

        let artist = data.artist;
        let locationData = data.locationData;
        let concertData = data.concertData;
        let relationData = data.relationData;

        let locations = locationData.locations || [];
        if (!Array.isArray(locations)) {
            locations = [];
            console.warn("Location data is not an array, defaulting to an empty array.");
        }

        let grid = document.getElementById("artistsGrid");
        grid.innerHTML = ""; // Clear the grid

        // Main Artist Card
        let card = document.createElement("div");
        card.classList.add("card");
        card.innerHTML = `
            <div class="card-header">
                <img class="artist-image" src="${artist.image}" alt="${artist.name}">
                <h3 class="artist-name">${artist.name}</h3>
            </div>
            <div class="card-body">
                <button class="expand-details" onclick="toggleDetails(event)">Show Details</button>
                <div class="details" style="display:none;">
                    <div class="details-section">
                        <h4>First Album</h4>
                        <p>${artist.firstAlbum}</p>
                    </div>
                    <div class="details-section">
                        <h4>Creation Date</h4>
                        <p>${artist.creationDate}</p>
                    </div>
                    <div class="details-section">
                        <h4>Members</h4>
                        <ul>${artist.members.map(member => `<li>${member}</li>`).join('')}</ul>
                    </div>
                    <div class="details-section">
                        <h4>Locations</h4>
                        <ul>${locations.length > 0 ? locations.map(location => `<li>${location}</li>`).join('') : '<li>No locations available</li>'}</ul>
                    </div>
                    <div class="details-section">
                        <h4>Concert Dates</h4>
                        <ul>${concertData.dates ? concertData.dates.map(concert => `<li>${concert}</li>`).join('') : '<li>No concert data available</li>'}</ul>
                    </div>
                    <div class="details-section">
                        <h4>Relations</h4>
                        <ul>${relationData.datesLocations ? Object.keys(relationData.datesLocations).map(location => {
                            return `<li>${location}: ${relationData.datesLocations[location].join(', ')}</li>`;
                        }).join('') : '<li>No relation data available</li>'}</ul>
                    </div>
                </div>
            </div>
        `;
        grid.appendChild(card);
    } catch (error) {
        console.error("Error fetching artist details:", error);
    }
}

// Function to toggle visibility of the details
function toggleDetails(event) {
    const details = event.target.closest('.card').querySelector('.details');
    details.style.display = details.style.display === "none" ? "block" : "none";
}





