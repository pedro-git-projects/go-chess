import {useEffect, useState} from "react"

export const GetBoard = () => {
  const [coordinates, setCoordinates] = useState([])
  useEffect(() => {
    const fetchContent = async () => {
      const headers = new Headers();
      headers.append("Content-Type", "application/json")
      const requestOptions = {
        method: "GET",
        headers: headers,
      }
      const response = await fetch(`http://localhost:8080/board`, requestOptions)
        .catch(err => {
          console.log(err)
        })
      const data = await response.json()
      setCoordinates(data)
    }
    fetchContent()
  },[])

  return(
    <div>
      {coordinates.map((c) => (
        <p key={c.coordinate}>{c.coordinate}:{c.piece}</p>
      ))}
    </div>
  )
}
