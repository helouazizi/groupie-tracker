import { renderError } from "./dom.js"

export async function fetchArtists() {
  try {
    const res = await fetch('http://localhost:8080/api/artists')
    if (!res.ok) {
      // If backend returned an error with JSON body
      const errorBody =  await res.json() 
      console.log(errorBody);
      
      throw {
        status: res.status,
        message: errorBody.message || "Server error",
        details: errorBody.details || ""
      }
    }
    // At this point, we are sure it's a valid response
    const data = await res.json()
    return data
  } catch (err) {
    renderError({
      status: err.status || "Network Error",
      message: err.message || "Something went wrong",
      details: err.details || ""
    })
    return []
  }
}
