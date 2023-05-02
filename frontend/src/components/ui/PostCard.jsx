import { Link } from "react-router-dom"

const PostCard = ({ post }) => (
  <Link to={`/posts/${post.id}`} className="bg-white rounded-lg shadow-md p-4">
    <h2 className="text-xl font-bold mb-2">{post.title}</h2>
    <img src={`data:image/jpeg;base64,${post.image}`} alt="Image" />
    <p className="text-gray-700">{post.body}</p>
  </Link>
)

export default PostCard
