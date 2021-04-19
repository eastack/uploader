onmessage = async e => {
    let hashString = await hash(e.data);
    postMessage("h", [e.data]);
}

async function hash(blob) {
    let hashBuffer = await crypto.subtle.digest('SHA-1', blob);
    let hashArray = Array.from(new Uint8Array(hashBuffer));
    fetch("http://localohst:8080/upload", {
        headers: {
            'Content-Range': 'bytes 0-199/600',
        },
        body: blob,
        method: 'PATCH'
    }).then(console.log)
    return hashArray.map(b => b.toString(16).padStart(2, '0')).join('');
}