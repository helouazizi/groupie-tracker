export function artistDetails() {

    const buttons = container.querySelectorAll(".details-btn");
    buttons.forEach((btn) => {
        btn.addEventListener("click", async (e) => {
            const artistId = btn.dataset.id;
            console.log(`http://localhost:8080/api/artists?id=${artistId}`);
            
            try {
                const res = await fetch(`http://localhost:8080/api/artists?id=${artistId}`);
                if (!res.ok) throw new Error("Failed to fetch artist details");
                const artistDetails = await res.json();
                //console.log("Artist Details:", artistDetails);
                displayArtistDetails(artistDetails)
                // You can display the artistDetails however you like:
                // showModal(artistDetails), renderDetailsPanel(artistDetails), etc.
                //alert(`Artist: ${artistDetails.name}\nFirst Album: ${artistDetails.firstAlbum}`);
            } catch (err) {
                console.error("Error fetching artist details:", err);
            }
        });
    });
}


function displayArtistDetails(data) {
    const container = document.getElementById('container')
    const panel = document.getElementById('artist-details-panel')
  
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
      <button></button>
    `
  }
  
