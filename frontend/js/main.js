import { fetchArtists } from "./api.js";
import { updateRangeValues, renderArtists, renderError } from "./dom.js";
import { setupFilters } from "./filters.js";
import { artistDetails } from "./details.js";

export async function loadHomePage() {
  try {
    const data = await fetchArtists();
    renderArtists(data);
    updateRangeValues();
    artistDetails(); // rebind buttons
    setupFilters();  // setup filters again
  } catch (err) {
    renderError(err);
  }
}
// status: err.status || "Error",
// message: err.message || "Failed to load artists.",
// details: err.details || "",

document.addEventListener("DOMContentLoaded", loadHomePage);
