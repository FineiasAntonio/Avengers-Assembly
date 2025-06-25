export function invocarModal(elemento) {
    elemento.style.display = "block";

    const handleClickOutside = function(event) {
        if (event.target === elemento) {
            elemento.style.display = "none";
            window.removeEventListener('click', handleClickOutside);
        }
    };
    
    setTimeout(() => {
        window.addEventListener('click', handleClickOutside);
    }, 0);
}

