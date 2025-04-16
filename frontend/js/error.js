export function renderError(err) {
    document.body.innerHTML = "";
  
    const errorContent = document.createElement("div");
    errorContent.classList.add("error-box");
  
    errorContent.innerHTML = `
      <h1>Oooops 😬</h1>
      <p><strong>${err.status || "Error"}</strong> | ${err.message || "Something went wrong."}</p>
      ${err.details ? `<pre>${err.details}</pre>` : ""}
      <button class="backhome-btn"><a href="/frontend/">Back Home</a></button>
    `;
  
    document.body.appendChild(errorContent);
  }
  