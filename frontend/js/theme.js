
export function toggleDarkMode() {
    document.body.classList.toggle("dark-mode");

    const btn = document.getElementById("theme-toggle");
    const isDark = document.body.classList.contains("dark-mode");
  
    btn.textContent = isDark ? "☀️" : "🌙"; // Show sun in dark mode, moon in light mode
}