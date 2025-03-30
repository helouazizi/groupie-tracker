
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
                    <h3 class ="artist-name">${artist.name}</h3>
                    <p class ="first-album" style="display:none">First Album: ${artist.firstAlbum}</p>
                    <p class ="creation-date" style="display:none">Creation Date: ${artist.creationDate}</p>
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
    let suggestionsBox = document.getElementById("suggestions");
    suggestionsBox.innerHTML = ""; // Clear previous suggestions

    let suggestions = [];

    cards.forEach(card => {
        let name = card.querySelector("h3").innerText.toLowerCase();
        if (name.includes(input)) {
            card.style.display = "block";
            suggestions.push(name);
        } else {
            card.style.display = "none";
        }
    });

    // Show suggestions below input
    if (input.length > 0) {
        suggestions.forEach(suggestion => {
            let suggestionItem = document.createElement("div");
            suggestionItem.classList.add("suggestion-item");
            suggestionItem.innerText = suggestion;
            suggestionItem.onclick = function () {
                document.getElementById("search").value = suggestion;
                suggestionsBox.innerHTML = ""; // Clear suggestions
                filterArtists(); // Filter based on selected suggestion
            };
            suggestionsBox.appendChild(suggestionItem);
        });
    }
}



async function artistDeatils(id) {
    try {
        let response = await fetch(`http://localhost:8080/artist?id=${id}`);
        let data = await response.json();

        let artist = data.artist;
        let mumbers = artist.members
        let locationData = data.locationData;
        let concertData = data.concertData;
        let relationData = data.relationData;

        // Ensure locationData.locations is an array
        let locations = locationData.locations || [];

        let grid = document.getElementById("artistsGrid");
        grid.innerHTML = ""; // Clear the grid

        // Artist Card
        let artistCard = document.createElement("div");
        artistCard.classList.add("card");
        artistCard.innerHTML = `
            <img src="${artist.image}" alt="${artist.name}">
            <div class="card-content">
                <h3>${artist.name}</h3>
                <p>First Album: ${artist.firstAlbum}</p>
                <p>Creation Date: ${artist.creationDate}</p>
            </div>
        `;
        grid.appendChild(artistCard);

        // Locations Card
        if (mumbers.length > 0) {
            let mumbercard = document.createElement("div");
            mumbercard.classList.add("card");
            mumbercard.innerHTML = `
                <div class="card-content">
                    <h3>Members</h3>
                    <ul>
                        ${mumbers.map(mumber => `<li>${mumber}</li>`).join('')}
                    </ul>
                </div>
            `;
            grid.appendChild(mumbercard);
        }
        // Locations Card
        if (locations.length > 0) {
            let locationCard = document.createElement("div");
            locationCard.classList.add("card");
            locationCard.innerHTML = `
                <div class="card-content">
                    <h3>Locations</h3>
                    <ul>
                        ${locations.map(location => `<li>${location}</li>`).join('')}
                    </ul>
                </div>
            `;
            grid.appendChild(locationCard);
        }

        // Concert Dates Card
        let concertCard = document.createElement("div");
        concertCard.classList.add("card");
        concertCard.innerHTML = `
            <div class="card-content">
                <h3>Concert Dates</h3>
                <ul>
                    ${concertData.dates ? concertData.dates.map(concert => `<li>${concert}</li>`).join('') : '<li>No concert data available</li>'}
                </ul>
            </div>
        `;
        grid.appendChild(concertCard);

        // Relations Card
        let relationCard = document.createElement("div");
        relationCard.classList.add("card");
        relationCard.innerHTML = `
            <div class="card-content">
                <h3>Relations</h3>
                <ul>
                    ${relationData.datesLocations ? Object.keys(relationData.datesLocations).map(location => {
            return `<li>${location}: ${relationData.datesLocations[location].join(', ')}</li>`;
        }).join('') : '<li>No relation data available</li>'}
                </ul>
            </div>
        `;
        grid.appendChild(relationCard);
    } catch (error) {
        console.error("Error fetching artist details:", error);
    }
}


// // Function to toggle visibility of the details
// function toggleDetails(event) {
//     const details = event.target.closest('.card').querySelector('.details');
//     details.style.display = details.style.display === "none" ? "block" : "none";
// }





