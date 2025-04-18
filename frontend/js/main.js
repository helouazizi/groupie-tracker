import { fetchArtists } from "./api.js";
// import { artistDetails } from "./details.js";
import { darckMode } from "./theme.js";
import { updateRangeValues, renderArtists } from "./dom.js";
import { setupFilters } from "./filters.js";

//document.addEventListener("DOMContentLoaded", async (e) => {
darckMode()
updateRangeValues()
fetchArtists()
setupFilters()
//});





