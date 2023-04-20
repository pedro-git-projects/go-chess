const sendMessage = (ws, message) => 
  new Promise((resolve, reject) => {
    if (ws.readyState === WebSocket.OPEN) {
      ws.onmessage = (e) => {
        const response = JSON.parse(e.data)
        console.log("Received response:", response)
        resolve(response)
      }
      ws.send(message)
      console.log("Sent message:", message)
    } else {
      reject(new Error("WebSocket connection is not open."))
    }
  })


export default sendMessage
