let text = document.getElementById('result').innerHTML;
const copyAscii = async () => {
    try {
        await navigator.clipboard.writeText(text);
        } catch (err) {
        alert("Copy failed")
    }
}