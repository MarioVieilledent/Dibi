
const i = document.getElementById('input');

const l = document.getElementById('list');

i.addEventListener('input', (e) => {
    let query = e.target.value;
    if (query.length > 10) {
        l.innerHTML = `<div class="word"><span>Trop de caractères dans la recherche</span></div>`
    } else if (query.length > 1) {
        fetch('http://' + location.hostname + ':3001/api/' + encodeURIComponent(e.target.value))
            .then((res) => res.json())
            .then((data) => {
                l.innerHTML = '';
                data.forEach(e => {
                    l.innerHTML += `<div class="word"><span>${e.dibi}</span> - <span>${e.french}</span> - <span>${e.english}</span></div>`
                });
            });
    }
});