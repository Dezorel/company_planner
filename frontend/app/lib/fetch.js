const requestURL = 'http://gateway:13080/api'

async function sendRequest(method, url, body = null) {
    const headers = {
        'Content-Type': 'application/json'
    }

    return await fetch(url, {
        method: method,
        body: JSON.stringify(body),
        headers:headers
    }).then(async response => {
        if (response.ok) {
            return await response.json()
        }

        return await response.json().then(error => {
            const e = new Error('Что-то пошло не так')
            e.data = error
            throw e
        })
    })
}

