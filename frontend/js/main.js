import { fetchArtists } from "./api.js";
import { updateRangeValues } from "./dom.js";
import { setupFilters } from "./filters.js";
import { artistDetails } from "./details.js";

document.addEventListener("DOMContentLoaded", async () => {
  fetchArtists();
  updateRangeValues();
  artistDetails();
  setupFilters();
});
