
function renderArtists(artists) {
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

function renderArtistDetails(data) {
  const container = document.getElementById('container')

  const { ArtistInfo, Locations, Dates, Relations } = data

  const locations = Locations.locations.map(loc => `<li>${loc.replace(/-/g, ', ').replace(/_/g, ' ')}</li>`).join("")
  const concertDates = Dates.dates.map(date => `<li>${date.replace('*', '')}</li>`).join("")
  const members = ArtistInfo.members.map(m => `<li>${m}</li>`).join("")

  const datesByLocation = Object.entries(Relations.datesLocations)
    .map(([place, dates]) => {
      const prettyPlace = place.replace(/-/g, ', ').replace(/_/g, ' ')
      const dateList = dates.map(d => `<li>${d}</li>`).join("")
      return `<div><h4>${prettyPlace}</h4><ul>${dateList}</ul></div>`
    }).join("")
  container.innerHTML = ""
  container.innerHTML = `
      <div class="card">
        <img src="${ArtistInfo.image}" alt="${ArtistInfo.name}" />
        <div>
          <h2>${ArtistInfo.name}</h2>
          <p>First Album: ${ArtistInfo.firstAlbum}</p>
          <p>Creation Year: ${ArtistInfo.creationDate}</p>
        </div>
      </div>
  
      <div class="card">
        <h3>Members</h3>
        <ul>${members}</ul>
      </div>
  
      <div class="card">
        <h3>Locations</h3>
        <ul>${locations}</ul>
      </div>
  
      <div class="card">
        <h3>Concert Dates</h3>
        <ul>${concertDates}</ul>
      </div>
      <div class="card">
        <h3>Concerts by Location</h3>
        ${datesByLocation}
      </div>
      
    `
  // <button class="home-btn" id="back-home-btn">Back Home</button>
  const btn = document.createElement("button")
  btn.innerHTML = "<a href='/frontend/'>Back Home</a>"
  btn.classList.add("home-btn")
  btn.setAttribute("id", "back-home-btn")
  const actions = document.getElementById("actions")
  actions.innerHTML = ""
  actions.append(btn)

  //document.getElementById("back-home-btn").addEventListener("click", loadHomePage);
}

const touchedInputs = new Set();
{ touchedInputs }
function updateRangeValues() {
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




function renderError(err) {
  if (window.hasGlobalError && window.hasGlobalError()) return;

  window.setGlobalError?.();

  document.body.innerHTML = "";

  const errorContent = document.createElement("div");
  errorContent.classList.add("error-box");

  errorContent.innerHTML = `
    <h1>Oooops 😬</h1>
    <p><strong>${err.status || "Error"}</strong> | ${err.message || "Something went wrong."}</p>
    <button class="home-btn" id="back-home-btn"><a href='/frontend/'>Back Home</a></button>
  `;

  document.body.appendChild(errorContent);
}



export { renderArtists, renderArtistDetails, renderError, updateRangeValues, touchedInputs }




