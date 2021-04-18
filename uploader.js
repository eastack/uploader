onmessage = async e => {
    let hashString = await hash(e.data);
    postMessage("h", [e.data]);
}

async function hash(blob) {
    let hashBuffer = await crypto.subtle.digest('SHA-1', blob);
    let hashArray = Array.from(new Uint8Array(hashBuffer));
    return hashArray.map(b => b.toString(16).padStart(2, '0')).join('');
}