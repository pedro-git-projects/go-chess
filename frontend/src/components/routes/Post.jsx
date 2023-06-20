import { useState, useEffect } from "react"
import { useParams } from "react-router-dom"
import { Layout } from "../ui/Layout"

function Post() {
  const [post, setPost] = useState(null)
  const { id } = useParams()

  useEffect(() => {
    async function fetchPost() {
      const response = await fetch(`http://localhost:1337/post?id=${id}`)
      const data = await response.json()
      setPost(data)
      console.log("data:", data)
    }
    fetchPost()
  }, [id])

  if (!post) {
    return (
      <Layout>
        <div className="h-screen flex flex-col items-center justify-center">
			<div className="bg-white dark:bg-gray-900 w-full max-w-3xl rounded-lg shadow-md overflow-hidden">
            <div className="p-4">
				<div className="w-full h-48 dark:bg-gray-800 bg-gray-300"></div>
              <h1 className="text-3xl font-bold p-5 text-black"></h1>
              <p className="text-gray-700 leading-relaxed"></p>
            </div>
          </div>
        </div>
      </Layout>
    )
  }

  return (
    <Layout>
      <div className=" min-h-screen py-4 flex flex-col items-center justify-center">
		  <div className="bg-white dark:bg-gray-800 w-full max-w-3xl rounded-lg shadow-md overflow-hidden">
          <div className="p-4">
            <img src={`data:image/jpeg;base64,${post.image}`} />
			  <h1 className="text-3xl dark:text-white font-bold p-5 text-black">{post.title}</h1>
			  <p className="text-gray-700 dark:text-white leading-relaxed">{post.content}</p>
          </div>
        </div>
      </div>
    </Layout>
  )
}

export default Post
