const touchedInputs = new Set();

///////////////////////////////////////////
async function fetchArtists() {
  try {
    const res = await fetch("http://localhost:8080/api/artists");
    if (!res.ok) {
      // If backend returned an error with JSON body
      const errorBody = await res.json();
      console.log(errorBody);

      throw {
        status: res.status,
        message: errorBody.message || "Server error",
        details: errorBody.details || "",
      };
    }
    // At this point, we are sure it's a valid response
    const data = await res.json();
    renderArtists(data);
  } catch (err) {
    renderError({
      status: err.status || "Network Error",
      message: err.message || "Something went wrong",
      details: err.details || "",
    });
    return [];
  }
}
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
      <button class="details-btn" onclick="artistDetails(${artist.id})">Details</button>
    `;

    container.appendChild(card);
  });
}
function updateRangeValues() {
  const ranges = document.querySelectorAll('input[type="range"]');

  ranges.forEach((range) => {
    const label = document.getElementById(`${range.id}-value`);
    if (!label) return;

    // Initial value display
    label.textContent = range.value;

    // Update display + mark as touched
    range.addEventListener("input", () => {
      label.textContent = range.value;
      touchedInputs.add(range.name);
    });
  });
}
function renderError(err) {
  document.body.innerHTML = "";

  const errorContent = document.createElement("div");
  errorContent.classList.add("error-box");

  errorContent.innerHTML = `
    <h1>Oooops 😬</h1>
    <p><strong>${err.status || "Error"}</strong> | ${
    err.message || "Something went wrong."
  }</p>
    ${err.details ? `<pre>${err.details}</pre>` : ""}
    <button class="backhome-btn"><a href="/frontend/">Back Home</a></button>
  `;

  document.body.appendChild(errorContent);
}

fetchArtists();
///////////////////////////////////////////////////

async function artistDetails(id) {
  // const container = document.getElementById("container");
  // const buttons = document.body.querySelectorAll(".details-btn");
  // console.log(buttons);
  // buttons.forEach((btn) => {
  //   btn.addEventListener("click", async () => {
  //     const artistId = btn.dataset.id;
  //     console.log(artistId);
  try {
    const res = await fetch(
      `http://localhost:8080/api/artists/details?id=${id}`
    );
    const data = await res.json(); // always parse JSON

    if (!res.ok) {
      // // If backend returned an error with JSON body
      // const errorBody = await res.json()
      // console.log(errorBody);

      throw {
        status: res.status,
        message: data.message || "Server error",
        details: data.details || "",
      };
    }

    displayArtistDetails(data);
  } catch (err) {
    // network failure or invalid JSON
    renderError({
      status: err.status || "Network Error",
      message: err.message || "Something went wrong",
      details: err.details || "",
    });
  }
  // });
  // });
}

function displayArtistDetails(data) {
  const container = document.getElementById("container");
  const panel = document.getElementById("artist-details-panel");

  const { ArtistInfo, Locations, Dates, Relations } = data;

  const locations = Locations.locations
    .map((loc) => `<li>${loc.replace(/-/g, ", ").replace(/_/g, " ")}</li>`)
    .join("");
  const concertDates = Dates.dates
    .map((date) => `<li>${date.replace("*", "")}</li>`)
    .join("");
  const members = ArtistInfo.members.map((m) => `<li>${m}</li>`).join("");

  const datesByLocation = Object.entries(Relations.datesLocations)
    .map(([place, dates]) => {
      const prettyPlace = place.replace(/-/g, ", ").replace(/_/g, " ");
      const dateList = dates.map((d) => `<li>${d}</li>`).join("");
      return `<div><h4>${prettyPlace}</h4><ul>${dateList}</ul></div>`;
    })
    .join("");
  container.innerHTML = "";
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
      <button class="backhome-btn"><a href="/frontend/">Back Home</a></button>
    `;
}



//////////////////////////////////////////////////////////

function setupFilters() {
  const form = document.getElementById("filter-form");

  form.addEventListener("submit", async (e) => {
    e.preventDefault();

    const filters = {
      creationDateFrom: touchedInputs.has("creation-from")
        ? form.querySelector('input[name="creation-from"]').value
        : "0",

      creationDateTo: touchedInputs.has("creation-to")
        ? form.querySelector('input[name="creation-to"]').value
        : "0",

      firstAlbumFrom: touchedInputs.has("album-from")
        ? form.querySelector('input[name="album-from"]').value
        : "0",

      firstAlbumTo: touchedInputs.has("album-to")
        ? form.querySelector('input[name="album-to"]').value
        : "0",

      members:
        form.querySelector('input[name="members"]:checked')?.value || "0",

      concertDates: form.querySelector("#location-input")?.value.trim() || "",
    };

    try {
      const res = await fetch("http://localhost:8080/api/artists/filter", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(filters),
      });
      const data = await res.json();
      console.log(res.status);

      console.log(data);

      if (!res.ok) {
        throw {
          status: res.status,
          message: data.message || "Server error",
          details: data.details || "",
        };
      }
      renderArtists(filteredArtists);
    } catch (err) {
      // network failure or invalid JSON
      renderError({
        status: err.status || "Network Error",
        message: err.message || "Something went wrong",
        details: err.details || "",
      });
    }
  });
}
setupFilters();
