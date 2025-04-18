import { fetchArtists } from "./api.js";
import { updateRangeValues , renderArtists } from "./dom.js";
import { setupFilters } from "./filters.js";
import { artistDetails } from "./details.js";

document.addEventListener("DOMContentLoaded", async () => {
  const data = await fetchArtists();
  renderArtists(data)
  updateRangeValues();
  artistDetails();
  setupFilters();
});
