import { fetchArtists } from "./api.js";
import { renderArtists ,updateRangeValues } from "./dom.js";
import { setupFilters } from "./filters.js";


document.addEventListener("DOMContentLoaded", async () => {
  const data = await fetchArtists();
  updateRangeValues();
  renderArtists(data);
  setupFilters();
});

// async function fetchArtists() {
//     try {
//         let response = await fetch("http://localhost:8080/");
//         if (!response.ok) {
//             throw new Error(`Error ${response.status}: ${response.statusText}`);
//         }
//         let artists = await response.json();
//         console.log(artists);
//         displayArtists(artists)
//     } catch (error) {
//         console.error("Error fetching artists:", error);
//         showErrorPage(error.message);
//     }
// }

// function displayArtists(artists) {

//     let grid = document.getElementById("artistsGrid");
//     grid.innerHTML = "";
//     artists.forEach(artist => {
//         let card = document.createElement("div");
//         card.classList.add("card");
//         card.innerHTML = `
//             <img src="${artist.image}" alt="${artist.name}">
//             <div class="card-content">
//                 <h3 class="artist-name">${artist.name}</h3>
//             <button onclick="artistDetails(${artist.id})">Details</button>
//         `;
//         grid.appendChild(card);
//     });
// }

// function toggleDarkMode() {
//     document.body.classList.toggle("dark-mode");
// }

// function searchArtists() {
//     let input = document.getElementById("search").value.toLowerCase();
//     let cards = document.querySelectorAll(".card");
//     let suggestionsBox = document.getElementById("suggestions");
//     suggestionsBox.innerHTML = ""; // Clear previous suggestions

//     let suggestions = [];

//     cards.forEach(card => {
//         let name = card.querySelector(".artist-name").innerText.toLowerCase();
//         let firstAlbum = card.querySelector(".first-album").innerText.toLowerCase().replace("first album: ", "");
//         let creationDate = card.querySelector(".creation-date").innerText.toLowerCase().replace("creation date: ", "");
//         let members = Array.from(card.querySelectorAll(".members li")).map(member => member.innerText.toLowerCase());
//         let locations = Array.from(card.querySelectorAll(".locations li")).map(location => location.innerText.toLowerCase());

//         let matchedCategories = [];

//         // Check matches for each category
//         if (name.includes(input)) {
//             matchedCategories.push({ text: card.querySelector(".artist-name").innerText, type: "artist/band" });
//         }
//         if (firstAlbum.includes(input)) {
//             matchedCategories.push({ text: firstAlbum, type: "first album date" });
//         }
//         if (creationDate.includes(input)) {
//             matchedCategories.push({ text: creationDate, type: "creation date" });
//         }
//         members.forEach(member => {
//             if (member.includes(input)) {
//                 matchedCategories.push({ text: member, type: "member" });
//             }
//         });
//         locations.forEach(location => {
//             if (location.includes(input)) {
//                 matchedCategories.push({ text: location, type: "location" });
//             }
//         });

//         if (matchedCategories.length > 0) {
//             card.style.display = "block";
//             matchedCategories.forEach(match => {
//                 let suggestionText = `${match.text} - ${match.type}`;
//                 if (!suggestions.includes(suggestionText)) {
//                     suggestions.push(suggestionText);
//                 }
//             });
//         } else {
//             card.style.display = "none";
//         }
//     });

//     // Show suggestions below input
//     if (input.length > 0) {
//         suggestions.forEach(suggestion => {
//             let suggestionItem = document.createElement("div");
//             suggestionItem.classList.add("suggestion-item");
//             suggestionItem.innerText = suggestion;
//             suggestionItem.onclick = function () {
//                 document.getElementById("search").value = suggestion.split(" - ")[0];
//                 suggestionsBox.innerHTML = ""; // Clear suggestions
//                 filterArtists(); // Filter based on selected suggestion
//             };
//             suggestionsBox.appendChild(suggestionItem);
//         });
//     }
// }

// async function artistDetails(id) {
//     try {
//         // Show loading state
//         let grid = document.getElementById("artistsGrid");
//         grid.innerHTML = "<p>Loading artist details...</p>";

//         let response = await fetch(`http://localhost:8080/artist?id=${id}`);

//         if (!response.ok) {
//             throw new Error(`Error ${response.status}: ${response.statusText}`);
//         }

//         let data = await response.json();

//         if (!data || !data.Artist) {
//             throw new Error("Artist data is missing or invalid");
//         }

//         let { Artist: artist, Locations: locationData, Dates: concertData, Relations: relationData } = data;

//         // Extract properties safely
//         let members = artist.members || [];
//         let locations = locationData?.locations || [];
//         let concerts = concertData?.dates || [];
//         let relations = relationData?.datesLocations || {};

//         // Clear previous content
//         grid.innerHTML = "";

//         // Create cards dynamically
//         let createCard = (title, content) => {
//             if (!content || content.length === 0) return ""; // Skip empty sections
//             return `
//                 <div class="card">
//                     <div class="card-content">
//                         <h3>${title}</h3>
//                         ${Array.isArray(content)
//                     ? `<ul>${content.map(item => `<li>${item}</li>`).join('')}</ul>`
//                     : `<p>${content}</p>`}
//                     </div>
//                 </div>
//             `;
//         };

//         // Append cards to grid
//         grid.innerHTML = `
//             <div class="card">
//                 <img src="${artist.image}" alt="${artist.name}">
//                 <div class="card-content">
//                     <h3>${artist.name}</h3>
//                     <p>First Album: ${artist.firstAlbum}</p>
//                     <p>Creation Date: ${artist.creationDate}</p>
//                 </div>
//             </div>
//             ${createCard("Members", members)}
//             ${createCard("Locations", locations)}
//             ${createCard("Concert Dates", concerts)}
//             ${createCard("Relations", Object.keys(relations).map(loc => `${loc}: ${relations[loc].join(', ')}`))}
//         `;
//     } catch (error) {
//         console.error("Error fetching artist details:", error);
//         showErrorPage(error.message);
//     }
// }

// function showErrorPage(message) {
//     let grid = document.getElementById("artistsGrid");
//     grid.innerHTML = `
//         <div class="error">
//             <h2>Oops! Something went wrong</h2>
//             <p>${message}</p>
//             <button><a href ="/">Back Home</a></button>
//         </div>
//     `;
//     return
// }

// function filter() {
//     document.getElementById("filters-form").addEventListener("submit", async (e) => {
//         e.preventDefault();

//         const submitButton = e.submitter;
//         submitButton.disabled = true;
//         submitButton.textContent = "Filtering...";

//         const creationDateFrom = document.getElementById("creation-date-from").value;
//         const creationDateTo = document.getElementById("creation-date-to").value;
//         const firstAlbumFrom = document.getElementById("first-album-from").value;
//         const firstAlbumTo = document.getElementById("first-album-to").value;
//         const concertDates = document.getElementById("concert-dates").value;
//         const membersRadio = document.querySelector('input[name="members"]:checked');
//         const members = membersRadio ? membersRadio.value : "";

//         // Validate numeric fields
//         const yearFields = [
//             { name: "Creation Date From", value: creationDateFrom },
//             { name: "Creation Date To", value: creationDateTo },
//             { name: "First Album From", value: firstAlbumFrom },
//             { name: "First Album To", value: firstAlbumTo },
//         ];

//         for (let field of yearFields) {
//             if (field.value && (isNaN(field.value) || parseInt(field.value) < 0)) {
//                 showErrorPage(`Invalid year for "${field.name}". Please enter a valid number.`);
//                 submitButton.disabled = false;
//                 submitButton.textContent = "Filter";
//                 return;
//             }
//         }

//         const filterData = {
//             creationDateFrom,
//             creationDateTo,
//             firstAlbumFrom,
//             firstAlbumTo,
//             members,
//             concertDates,
//         };

//         try {
//             const response = await fetch("http://localhost:8080/filter", {
//                 method: "POST",
//                 headers: {
//                     "Content-Type": "application/json",
//                 },
//                 body: JSON.stringify(filterData),
//             });

//             if (!response.ok) {
//                 const errorMessage = await response.text();
//                 showErrorPage(`Server Error: ${response.status} ${response.statusText} - ${errorMessage}`);
//                 return;
//             }

//             const result = await response.json();

//             if (!Array.isArray(result) || result.length === 0) {
//                 showErrorPage("No artists matched your filters.");
//                 return;
//             }

//             displayArtists(result);

//         } catch (error) {
//             console.error("Error filtering artists:", error);
//             showErrorPage("Something went wrong while contacting the server.");
//         } finally {
//             submitButton.disabled = false;
//             submitButton.textContent = "Filter";
//         }
//     });
// }

// fetchArtists();
// filter()
