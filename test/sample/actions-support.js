window.addEventListener("update", readyActions, false)
function readyActions() {
    const freeActions = document.getElementsByClassName("free-action")
    for (let i = 0; i < freeActions.length; i++) {
        const action = freeActions[i].value.split("::");
        const name = action.shift()
        const params = action
        switch (name) {
            case "ALERT":
                freeActions[i].addEventListener("click", free_alert, false)
                freeActions[i].params = params
                break;
            default:
                break;
        }
    }
}
// ACTIONS
function free_alert(ev) {
    alert(ev.currentTarget.params)
}