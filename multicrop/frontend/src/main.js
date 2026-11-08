
document.addEventListener("DOMContentLoaded", (ev) => {
    window.addEventListener("drop", (ev) => {
        ev.preventDefault()
    })
    window.addEventListener("dragover", (ev) => {
        ev.preventDefault()
        ev.dataTransfer.dropEffect = 'copy'
    })
})

// got from wails issue #3985 - this should allow DnD events to propagate to backend
// on Windows 10
runtime.OnFileDrop(function (x, y, paths) {}, false);
