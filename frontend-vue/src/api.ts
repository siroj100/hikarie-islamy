export default {
    async v1QuranPage(layoutId: int, pageNumber: int) {
        console.log('v1QuranPage:', layoutId, pageNumber)
        const resp = await fetch(`http://localhost:5080/api/quran/v1/page/${layoutId}/${pageNumber}`)
        console.log('v1QuranPage:', resp)
        if (resp.ok) {
            return await resp.json()
        }
        return null;
    }
}