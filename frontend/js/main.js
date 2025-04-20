import { fetchArtists } from "./api.js";
import { artistDetails } from "./details.js";
import { darckMode } from "./theme.js";
import { updateRangeValues } from "./dom.js";
import { setupFilters } from "./filters.js";
let globalErrorTriggered = false;

window.setGlobalError = () => {
    globalErrorTriggered = true;
};

window.hasGlobalError = () => globalErrorTriggered;

//document.addEventListener("DOMContentLoaded", async (e) => {
darckMode()
updateRangeValues()
if (!window.hasGlobalError?.()) {
    fetchArtists();
    setupFilters();
    //artistDetails();
}

//});





