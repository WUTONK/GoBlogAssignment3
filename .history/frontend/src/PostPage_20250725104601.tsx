import { useState } from 'react'
import './app.css'
import { Api } from './shared'
import { useNavigate } from 'react-router-dom'

function InfoPage() {
  const navigate = useNavigate() //路由
  const [nickname, setNickname] = useState("")

  // 校验本地token
  Api.userInfoGet(
    {}, // 或 { infoReq: { authorization: "" } }，看后端需不需要
    {
      headers: {
        Authorization: localStorage.getItem("token") || ""
      }
    }
  ).then(
    (res) => setNickname(res.nickName)
  ).catch(
    (err) => alert(err.message + "校验token失败 返回登录页面")
    navigate('/user/login')
  )

  return (
    <>      
      <p><strong>HI:</strong> {nickname}</p>
    </>
  )
}

export default InfoPage
