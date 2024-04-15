var pswd = ""
document.querySelector('input').addEventListener('input', (e) => {
    pswd = e.target.value
})
document.querySelector('button').addEventListener('click', () => {
    window.location.href = '/' + pswd
})
