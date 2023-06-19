const cpyURL = () => {
  const url = window.location.href
  const index = url.indexOf("room")
  const cutLen = 5
  const roomID = url.substring(index+cutLen, url.length)
  navigator.clipboard.writeText(roomID)
}

const CopyToClipboard = () => (
  <button className="rounded border border-[#00A29C] bg-[#00A29C] hover:bg-[#3ba8a4] text-white font-bold py-1 px-4 rounded ml-2" onClick={cpyURL}> 
    Copy Code to Clipboard
  </button>
)


export default CopyToClipboard
