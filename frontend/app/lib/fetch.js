const requestURL = 'http://127.0.0.1:13080/api'

async function sendRequest(method, url, body = null) {
    const headers = {
        'Content-Type':'application/json'
    }

    if (method !== "GET" && method !== "HEAD")
    {
        getRequestObject = {
            method: method,
            body: JSON.stringify(body),
            headers:headers
        }
    }
    else
    {
        getRequestObject = {
            method: method,
            headers:{'Content-Type':'text/html'}
        }
    }


    return await fetch(url, getRequestObject).then(async response => {
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

