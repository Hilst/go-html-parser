window.customLog = function(color, text) {
    console.log(`%c${text}`, `color:${color}`)
}

window.pagechange = function(currentactive, nextactive) {
    const p = "page_" + currentactive
    const n = "page_" + nextactive
    let current = document.getElementById(p)
    let next = document.getElementById(n)
    if (!current || !next) return
    current.hidden = true
    next.hidden = false
}