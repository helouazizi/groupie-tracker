
  // Tema değiştirme fonksiyonu
 export function darckMode(params) {
    const themeButton = document.getElementById('theme');
    themeButton.addEventListener('click', () => {
        document.body.classList.toggle('dark-mode');
        const isDark = document.body.classList.contains('dark-mode');
        themeButton.textContent = isDark ? '🌞' : '🌙'; // Tema butonunun ikonunu değiştir
    });
 }