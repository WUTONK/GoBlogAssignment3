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
    Api.userPostModifyPost(
      { 
        // 1. 发送查询请求
        sqlReq: {
          userName: localStorage.getItem("userName") || "",
          mode: "get",
          token: localStorage.getItem("token") || "",
          appendText: "",
        }
      }
    ).then(
      (res) => {
        // 2. 分割并获取字符串
        const resPosts = res.context ? res.context.split("<slice>") : []
        setPosts(resPosts)
        console.log(resPosts)
      }
    ).catch((err) => {
      alert(err.message + "校验token失败 返回登录页面");
      navigate("/user/login");
    })
  },[navigate])

  return (
    <>
      {/* 3. 显示为一行一个报文 */}
      {posts.map((post, idx) => (
        <div key={idx}>{post}</div>
      ))}
    </>
  )
}

export default PostShowPage
