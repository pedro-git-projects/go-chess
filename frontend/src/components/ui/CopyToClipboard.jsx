const cpyURL = () => {
  const url = window.location.href
  const index = url.indexOf("room")
  const cutLen = 5
  const roomID = url.substring(index+cutLen, url.length)
  navigator.clipboard.writeText(roomID)
}

const CopyToClipboard = () => (
  <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded" onClick={cpyURL}> 
    Copy Code to Clipboard
  </button>
)


export default CopyToClipboard
