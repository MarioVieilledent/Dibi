// Elements du dom
const i = document.getElementById('input');
const l = document.getElementById('list');

// Autofocus sur barre de recherche
setTimeout(() => {
    i.focus()
}, 10);


// Options d'affichage
let display = {
    english: false,
    author: true
}

// À chaque recherche
i.addEventListener('input', (e) => {
    let query = e.target.value;
    if (query.length > 10) {
        l.innerHTML = `<div class="word"><span>Trop de caractères dans la recherche</span></div>`
    } else if (query.length > 1) {
        fetch('http://' + location.hostname + '/api/' + encodeURIComponent(e.target.value))
            .then((res) => res.json())
            .then((data) => {
                l.innerHTML = '';
                data.forEach(e => {
                    l.innerHTML += `<div class="word ${e.partOfSpeech}">
                    <div class="dibi ${e.partOfSpeech}">
                        <span class="dibiWord">${e.dibi}</span>
                        <span class="partOfSpeech">${partOfSpeechSimpler(e.partOfSpeech)}</span>
                    </div>
                    <div class="french">
                        <span>${e.french}</span>
                    </div>
                    ${display.english ? `
                        <div class="english">
                        <span>${e.english}</span>
                    </div>`: ''}
                    <div class="description">
                        <span>${e.description}</span>
                    </div>
                    ${display.author ? `
                        <div class="author">
                        <span>${e.author}</span>
                    </div>`: ''}
                    </div>`
                });
                l.scrollTop = 0;
            });
    }
});

// Affiche simplement le nature grammaticale
function partOfSpeechSimpler(pos) {
    switch (pos) {
        case 'Noun': return 'n.';
        case 'Pronoun': return 'pron.';
        case 'Verb': return 'v.';
        case 'Adjective': return 'adj.';
        case 'Adverb': return 'adv.';
        case 'Conjonction': return 'conj.';
        case 'FunctionParticule': return 'fonc.';
        case 'TransformationParticule': return 'trans.';
        case 'SpiritWord': return 'esp.';
        case 'Interjection': return 'intej.';
        default: return `Erreur : ${pos}`
    }
}