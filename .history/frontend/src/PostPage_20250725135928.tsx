import { useState } from 'react'
import './app.css'
import { Api } from './shared'
import { useNavigate } from 'react-router-dom'

function InfoPage() {
  const navigate = useNavigate() //路由
  const [nicknames, setNicknames] = useState<string[]>([])

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
      const names = res.nickName ? res.nickName.split("<slice>") : []
      setNicknames(names)
    }
  ).catch((err) => {
      alert(err.message + "校验token失败 返回登录页面");
      navigate("/user/login");
  })
  
  return (
    <>      
       <p><strong>HI:</strong></p>
      {/* 3. 一行一个名字 */}
      {nicknames.map((name, idx) => (
      <div key={idx}>{name}</div>
      ))}
    </>
  )
}

export default InfoPage
