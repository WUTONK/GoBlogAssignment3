import { useEffect, useState } from 'react'
import './app.css'
import { Api } from './shared'
import { useNavigate } from 'react-router-dom'

function PostShowPage() {
  const navigate = useNavigate() //路由
  const [posts, setPosts] = useState<string[]>([])

  // 使用 Effect 保证只执行一遍
  useEffect(() => {
    // 校验本地token
    Api.userInfoGet(
      {}, // 或 { infoReq: { authorization: "" } }，看后端需不需要
      {
        headers: {
          Authorization: localStorage.getItem("token") || ""
        }
      }
    ).then(
      (res) => {
        // 2. 分割字符串
        const resPosts = res.nickwName ? res.nickName.split("<slice>") : []
        setPosts(resPosts)
        console.log(resPosts)
      }
    ).catch((err) => {
      alert(err.message + "校验token失败 返回登录页面");
      navigate("/user/login");
    })
  }, [])

  return (
    <>
      {/* 3. 一行一个名字 */}
      {posts.map((post, idx) => (
        <div key={idx}>{post}</div>
      ))}
    </>
  )
}

export default PostShowPage
