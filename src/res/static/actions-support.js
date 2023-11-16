/**
 * Logs a text with the color passed
 * @param {string} color 
 * @param {string} text 
 */
window.customLog = function(color, text) {
    console.log(`%c${text}`, `color:${color}`)
}

/**
 * Change page from currentactive to nextactive
 * making the element with id page_currentactive hidden
 * and the element with id page_nextactive not hidden
 * @param {string} currentactive 
 * @param {string} nextactive 
 */
window.pagechange = function(currentactive, nextactive) {
    const p = "page_" + currentactive
    const n = "page_" + nextactive
    let current = document.getElementById(p)
    let next = document.getElementById(n)
    if (!current || !next) return
    current.hidden = true
    next.hidden = false
}