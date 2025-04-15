import { fetchArtists } from "./api.js";
import { renderArtists ,updateRangeValues } from "./dom.js";
import { setupFilters } from "./filters.js";


document.addEventListener("DOMContentLoaded", async () => {
  const data = await fetchArtists();
  updateRangeValues();
  renderArtists(data);
  setupFilters();
});
