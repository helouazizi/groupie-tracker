export function toggleDarkMode() {
    const button = document.getElementById('theme-toggle');
    document.body.classList.toggle('dark');
  
    // Toggle sun and moon
    if (document.body.classList.contains('dark')) {
      button.textContent = '🌙';  // Moon icon in dark mode
      localStorage.setItem('theme', 'dark');
    } else {
      button.textContent = '🌞';  // Sun icon in light mode
      localStorage.setItem('theme', 'light');
    }
  }
  
  export function applySavedTheme() {
    const saved = localStorage.getItem('theme');
    const button = document.getElementById('theme-toggle');
    if (saved === 'dark') {
      document.body.classList.add('dark');
      button.textContent = '🌙'; // Moon icon for dark mode
    } else {
      document.body.classList.remove('dark');
      button.textContent = '🌞'; // Sun icon for light mode
    }
  }
  