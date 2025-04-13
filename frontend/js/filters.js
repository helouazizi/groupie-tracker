import { renderArtists } from './dom.js'

export function setupFilters(data) {
  const filterSection = document.getElementById('filters')

  const input = document.createElement('input')
  input.placeholder = 'Search by name...'
  input.type = 'text'

  input.addEventListener('input', () => {
    const query = input.value.toLowerCase()
    const filtered = data.filter((a) =>
      a.Name.toLowerCase().includes(query)
    )
    renderArtists(filtered)
  })

  filterSection.appendChild(input)
}

