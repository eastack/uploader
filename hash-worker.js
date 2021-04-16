onmessage = async function (e) {
    postMessage(await hash(e.data));
}

async function hash(blob) {
    const hashBuffer = await crypto.subtle.digest('SHA-256', blob);
    const hashArray = Array.from(new Uint8Array(hashBuffer));
    return hashArray.map(b => b.toString(16).padStart(2, '0')).join('');
}