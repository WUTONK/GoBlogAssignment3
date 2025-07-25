import { useState } from 'react'
import './app.css'
import { Api } from './shared'

function InfoPage() {
  const [nickname, setNickname] = useState("")

  console.log("调用infofget开始")
  Api.userInfoGet(
    {}, // 或 { infoReq: { authorization: "" } }，看后端需不需要
    {
      headers: {
        Authorization: localStorage.getItem("token") || ""
      }
    }
  ).then(
    (res) => setNickname(res.nickName)
  ).catch((err) => alert(err.message))
  console.log("调用infofget结束")

  return (
    <>      
      <p><strong>HI:</strong> {nickname}</p>
    </>
  )
}

export default InfoPage
