import { useNavigate } from 'react-router-dom'

const NewRoom = () => {
  const navigate = useNavigate()

  const handleClick = async () => {
    // Send a GET request to create a new room with the given name
    const response = await fetch(`http://localhost:8080/room/new`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
    });

    if (!response.ok) {
      // Handle error response
      console.error(`Failed to create room: ${response.statusText}`);
      return;
    }

    const { clientID } = await response.json();

    // Navigate to the new room's URL
    navigate(`/room/${clientID}`);
    console.log(clientID)
  };

  return (
    <>
      <button type="button" onClick={handleClick}>Create room</button>
    </>
  )
}

export default NewRoom
