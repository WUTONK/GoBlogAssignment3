import { useEffect, useState } from 'react'
import './app.css'
import { Api } from './shared'
import { useNavigate } from 'react-router-dom'

function PostShowPage() {
  const navigate = useNavigate() //路由
  const [posts, setPosts] = useState<string[]>([])
  const [loading, setLoading] = useState(true) // 添加加载状态

  // 使用 Effect 保证只执行一遍
  useEffect(() => {
    const fetchPosts = async () => {
      try {
        setLoading(true)
        // 校验本地token
        const token = localStorage.getItem("token")
        if (!token) {
          alert("未找到登录token，请重新登录")
          navigate("/user/login")
          return
        }

        const response = await Api.userPostModifyPost({
          sqlReq: {
            mode: "get",
            token: token
          }
        })

        // 2. 分割并获取字符串
        const resPosts = response.context ? response.context.split("<slice>") : []
        setPosts(resPosts)
        console.log("获取到的文章列表:", resPosts)
      } catch (err: unknown) {
        console.error("获取文章失败:", err)
        const errorMessage = err instanceof Error ? err.message : "获取文章失败，请重新登录"
        alert(errorMessage)
        navigate("/user/login")
      } finally {
        setLoading(false)
      }
    }

    fetchPosts()
  }, [navigate]) // ✅ 添加依赖数组

  if (loading) {
    return <div>加载中...</div>
  }

  return (
    <>
      {/* 3. 显示为一行一个报文 */}
      {posts.length === 0 ? (
        <div>暂无文章</div>
      ) : (
        posts.map((post, idx) => (
          <div key={idx} style={{ margin: '10px 0', padding: '10px', border: '1px solid #ddd' }}>
            {post}
          </div>
        ))
      )}
    </>
  )
}

export default PostShowPage
