
document.addEventListener("DOMContentLoaded", (ev) => {
    window.addEventListener("drop", (ev) => {
        ev.preventDefault()
    })
    window.addEventListener("dragover", (ev) => {
        ev.preventDefault()
        ev.dataTransfer.dropEffect = 'copy'
    })
})
