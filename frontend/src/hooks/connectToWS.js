const connectToWS = (url) =>
  new Promise((resolve, reject) => {
    const ws = new WebSocket(url)
    ws.onopen = () => {
      console.log("WebSocket connection established.")
      resolve(ws)
    }
    ws.onerror = (e) => {
      reject(e)
    }
  })

export default connectToWS
