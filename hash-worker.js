onmessage = async function (e) {
    console.log("calculate blob hash starting...")
    postMessage(await hash(e.data));
    console.log("calculate blob hash finished...")
}

async function hash(blob) {
    const arrayBuffer = await blob.arrayBuffer();
    const hashBuffer = await crypto.subtle.digest('SHA-256', arrayBuffer);
    const hashArray = Array.from(new Uint8Array(hashBuffer));
    return hashArray.map(b => b.toString(16).padStart(2, '0')).join('');
}