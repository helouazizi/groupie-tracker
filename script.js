async function fetchArtists() {
    try {
        let response = await fetch("http://localhost:8080/");
        if (!response.ok) {
            throw new Error(`Error ${response.status}: ${response.statusText}`);
        }
        let artists = await response.json();
        console.log(artists);
        displayArtists(artists)
    } catch (error) {
        console.error("Error fetching artists:", error);
        showErrorPage(error.message);
    }
}

function displayArtists(artists) {

    let grid = document.getElementById("artistsGrid");
    grid.innerHTML = "";
    artists.forEach(artist => {
        let card = document.createElement("div");
        card.classList.add("card");
        card.innerHTML = `
            <img src="${artist.image}" alt="${artist.name}">
            <div class="card-content">
                <h3 class="artist-name">${artist.name}</h3>
            <button onclick="artistDetails(${artist.id})">Details</button>
        `;
        grid.appendChild(card);
    });
}

// Function to filter based on user input from form controls
// function filterArtists() {
//     let creationDateRange = document.getElementById("creationDateRange").value.split("-");
//     let firstAlbumRange = document.getElementById("firstAlbumRange").value.split("-");
//     let memberCount = document.getElementById("memberCount").value;
//     let locationFilter = document.getElementById("locationFilter").value.toLowerCase();

//     let cards = document.querySelectorAll(".card");
    
//     cards.forEach(card => {
//         let creationDate = parseInt(card.getAttribute("data-creation-date"));
//         let firstAlbumDate = parseInt(card.getAttribute("data-first-album"));
//         let membersCount = parseInt(card.getAttribute("data-members"));
//         let locations = card.getAttribute("data-locations").toLowerCase();

//         let isVisible = true;

//         // Filter by creation date range
//         if (creationDate < parseInt(creationDateRange[0]) || creationDate > parseInt(creationDateRange[1])) {
//             isVisible = false;
//         }

//         // Filter by first album date range
//         if (firstAlbumDate < parseInt(firstAlbumRange[0]) || firstAlbumDate > parseInt(firstAlbumRange[1])) {
//             isVisible = false;
//         }

//         // Filter by number of members
//         if (memberCount && membersCount !== parseInt(memberCount)) {
//             isVisible = false;
//         }

//         // Filter by location
//         if (locationFilter && !locations.includes(locationFilter)) {
//             isVisible = false;
//         }

//         // Apply visibility
//         card.style.display = isVisible ? "block" : "none";
//     });
// }

// Call this function when filters change
// function applyFilters() {
//     filterArtists();
// }

// Ensure event listeners are applied correctly
// window.addEventListener("DOMContentLoaded", () => {
//     document.getElementById("creationDateRange").addEventListener("input", applyFilters);
//     document.getElementById("firstAlbumRange").addEventListener("input", applyFilters);
//     document.getElementById("memberCount").addEventListener("input", applyFilters);
//     document.getElementById("locationFilter").addEventListener("input", applyFilters);
// });


function toggleDarkMode() {
    document.body.classList.toggle("dark-mode");
}

function searchArtists() {
    let input = document.getElementById("search").value.toLowerCase();
    let cards = document.querySelectorAll(".card");
    let suggestionsBox = document.getElementById("suggestions");
    suggestionsBox.innerHTML = ""; // Clear previous suggestions

    let suggestions = [];

    cards.forEach(card => {
        let name = card.querySelector(".artist-name").innerText.toLowerCase();
        let firstAlbum = card.querySelector(".first-album").innerText.toLowerCase().replace("first album: ", "");
        let creationDate = card.querySelector(".creation-date").innerText.toLowerCase().replace("creation date: ", "");
        let members = Array.from(card.querySelectorAll(".members li")).map(member => member.innerText.toLowerCase());
        let locations = Array.from(card.querySelectorAll(".locations li")).map(location => location.innerText.toLowerCase());

        let matchedCategories = [];

        // Check matches for each category
        if (name.includes(input)) {
            matchedCategories.push({ text: card.querySelector(".artist-name").innerText, type: "artist/band" });
        }
        if (firstAlbum.includes(input)) {
            matchedCategories.push({ text: firstAlbum, type: "first album date" });
        }
        if (creationDate.includes(input)) {
            matchedCategories.push({ text: creationDate, type: "creation date" });
        }
        members.forEach(member => {
            if (member.includes(input)) {
                matchedCategories.push({ text: member, type: "member" });
            }
        });
        locations.forEach(location => {
            if (location.includes(input)) {
                matchedCategories.push({ text: location, type: "location" });
            }
        });

        if (matchedCategories.length > 0) {
            card.style.display = "block";
            matchedCategories.forEach(match => {
                let suggestionText = `${match.text} - ${match.type}`;
                if (!suggestions.includes(suggestionText)) {
                    suggestions.push(suggestionText);
                }
            });
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
                document.getElementById("search").value = suggestion.split(" - ")[0];
                suggestionsBox.innerHTML = ""; // Clear suggestions
                filterArtists(); // Filter based on selected suggestion
            };
            suggestionsBox.appendChild(suggestionItem);
        });
    }
}

async function artistDetails(id) {
    try {
        let response = await fetch(`http://localhost:8080/artist?id=${id}`);

        // Handle non-OK responses immediately (e.g., 404 Not Found)
        if (!response.ok) {
            throw new Error(`Error ${response.status}: ${response.statusText}`);
        }

        // Parse JSON safely
        let data;
        try {
            data = await response.json();
        } catch (jsonError) {
            throw new Error("Invalid JSON response from server");
        }

        // If the response doesn't contain the expected data, stop execution
        if (!data) {
            throw new Error("Artist data is missing or invalid");
        }
        console.log(data);
        

        // Extracting data safely
        let artist = data.Artist;
        let members = data.Artist.members || [];
        let locationData = data.Locations || { Locations: [] };
        let concertData = data.Dates || { Dates: [] };
        let relationData = data.Relations || { Relations: {} };

        // Clear previous content
        let grid = document.getElementById("artistsGrid");
        grid.innerHTML = "";

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

        // Members Card
        if (members.length > 0) {
            let memberCard = document.createElement("div");
            memberCard.classList.add("card");
            memberCard.innerHTML = `
                <div class="card-content">
                    <h3>Members</h3>
                    <ul>
                        ${members.map(member => `<li>${member}</li>`).join('')}
                    </ul>
                </div>
            `;
            grid.appendChild(memberCard);
        }

        // Locations Card
        if (locationData.locations.length > 0) {
            let locationCard = document.createElement("div");
            locationCard.classList.add("card");
            locationCard.innerHTML = `
                <div class="card-content">
                    <h3>Locations</h3>
                    <ul>
                        ${locationData.locations.map(location => `<li>${location}</li>`).join('')}
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
                    ${concertData.dates.length > 0 ? concertData.dates.map(concert => `<li>${concert}</li>`).join('') : '<li>No concert data available</li>'}
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
                    ${Object.keys(relationData.datesLocations).length > 0 
                        ? Object.keys(relationData.datesLocations).map(location => {
                            return `<li>${location}: ${relationData.datesLocations[location].join(', ')}</li>`;
                          }).join('')
                        : '<li>No relation data available</li>'}
                </ul>
            </div>
        `;
        grid.appendChild(relationCard);
    } catch (error) {
        console.error("Error fetching artist details:", error);
        showErrorPage(error.message);
    }
}



function showErrorPage(message) {
    let grid = document.getElementById("artistsGrid");
    grid.innerHTML = `
        <div class="error">
            <h2>Oops! Something went wrong</h2>
            <p>${message}</p>
            <button><a href ="/">Back Home</a></button>
        </div>
    `;
    return
}


// Function to filter based on user input from form controls
// function filterArtists() {
//     let creationDateRange = document.getElementById("creationDateRange").value.split("-");
//     let firstAlbumRange = document.getElementById("firstAlbumRange").value.split("-");
//     let memberCount = document.getElementById("memberCount").value;
//     let locationFilter = document.getElementById("locationFilter").value.toLowerCase();

//     let cards = document.querySelectorAll(".card");
    
//     cards.forEach(card => {
//         let creationDate = parseInt(card.getAttribute("data-creation-date"));
//         let firstAlbumDate = parseInt(card.getAttribute("data-first-album"));
//         let membersCount = parseInt(card.getAttribute("data-members"));
//         let locations = card.getAttribute("data-locations").toLowerCase();

//         let isVisible = true;

//         // Filter by creation date range
//         if (creationDate < parseInt(creationDateRange[0]) || creationDate > parseInt(creationDateRange[1])) {
//             isVisible = false;
//         }

//         // Filter by first album date range
//         if (firstAlbumDate < parseInt(firstAlbumRange[0]) || firstAlbumDate > parseInt(firstAlbumRange[1])) {
//             isVisible = false;
//         }

//         // Filter by number of members
//         if (memberCount && membersCount !== parseInt(memberCount)) {
//             isVisible = false;
//         }

//         // Filter by location
//         if (locationFilter && !locations.includes(locationFilter)) {
//             isVisible = false;
//         }

//         // Apply visibility
//         card.style.display = isVisible ? "block" : "none";
//     });
// }

// Call this function when filters change
// function applyFilters() {
//     filterArtists();
// }

// Event listeners for filter inputs
// document.getElementById("creationDateRange").addEventListener("input", applyFilters);
// document.getElementById("firstAlbumRange").addEventListener("input", applyFilters);
// document.getElementById("memberCount").addEventListener("input", applyFilters);
// document.getElementById("locationFilter").addEventListener("input", applyFilters);









fetchArtists();