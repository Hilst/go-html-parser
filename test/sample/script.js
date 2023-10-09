const URL = "http://localhost:8080"

let pageSelector = document.getElementById("pages-selector")
let goButton = document.getElementById("go-button")
let previousButton = document.getElementById("previous-button")
let nextButton = document.getElementById("next-button")
let pageNumber = document.getElementById("page-number-label")
let main = document.getElementsByTagName("main").item(0)
let tryMessageCollection = document.getElementsByClassName("try-message")

class HTMLS {
    items = []
    current = undefined
    constructor(list, errs) {
        var l = []
        errs.forEach((err, i) => {
            if (!err) {
                l = l.concat(list[i])
            }
        });
        this.items = l
        if (this.items.length > 0) {
            this.updateCurrent(0)
        }
    }
    updateCurrent(index) {
        this.current = index
        pageNumber.innerHTML = parseInt(index)
        if (index == 0) {
            previousButton.disabled = true
        } else {
            previousButton.disabled = false
        }
        if (index + 1 < this.size()) {
            nextButton.disabled = false
        } else {
            nextButton.disabled = true
        }
        this.sendCurrent()
    }
    size() {
        return this.items.length
    }
    sendCurrent() {
        const string = this.items[this.current]
        main.innerHTML = string
        sendUpdateEvent()
        changeTryVisibility(false)
    }
}
var model

function changeTryVisibility(shouldSee) {
    for (let i = 0; i < tryMessageCollection.length; i++) {
        const element = tryMessageCollection[i];
        element.style.visibility = shouldSee ? "visible" : "hidden"
    }
}

goButton.addEventListener("click", (function (_ev) {
    const hash = pageSelector.value
    const path = "/layout/" + hash
    const json = JSON.parse(http(path))
    if (json["layouts_htmls"] !== undefined) {
        model = new HTMLS(json["layouts_htmls"], json["errors_list"])
    } 
}))

function http(path, method) {
    try {
        var req = new XMLHttpRequest()
        req.open(method || "GET", URL + path, false)
        req.send()
        return req.response
    } catch (err) {
        alert("COULD NOT RECEIVE ANSWER FROM BACKEND")
        return JSON.stringify(err)
    }
}

nextButton.addEventListener("click", (function (_ev) {
    model.updateCurrent(model.current + 1)
}))
previousButton.addEventListener("click", (function (_ev) {
    model.updateCurrent(model.current - 1)
}))

pageSelector.addEventListener("change", (function (_ev) {
    changeTryVisibility(true)
}))

function sendUpdateEvent() {
    const ev = new Event("update")
    window.dispatchEvent(ev)
}