import { useState } from 'react'
import './app.css'
import { Button, Input, Space } from 'antd'
import { Api } from './shared'
import { useNavigate } from 'react-router-dom'

function LoginPage() {
  const navigate = useNavigate()
  const [password, setPassword] = useState("")
  const [username, setUsername] = useState("")

  return (
    <>
     <div className='app-container'>
     <Space direction="vertical" size="large" style={{ innerHeight: '100%' }}>
      <Space direction="vertical" size="small" style={{ innerHeight: '100%' }}>
      <div className='input-password'>
        <Input
          type="text"
          value={username}
          placeholder="请输入用户名"
          onChange={(e: React.ChangeEvent<HTMLInputElement>) => setUsername(e.target.value)}
        />
      </div>
      <div>
        <Input
            type="text"
            value={password}
            placeholder="请输入密码"
            onChange={(e: React.ChangeEvent<HTMLInputElement>) => setPassword(e.target.value)}
            className="password-input"
          />
      </div>
      </Space>
      
      {/* 登录按钮 */}
      <Button
      className='login-button'
      title="登陆"
      onClick={()=>{
        Api.userLoginPost(
          {
            loginReq:{
              username,
              password
            } 
          }
        ).then((res)=>{
          // 将 token 缓存到本地 
          // 现在更改为直接使用res.token 避免了异步操作导致的旧token被错误的延迟缓存从而没有被覆盖的问题
          localStorage.setItem("token",res.token)
          localStorage.setItem("userName",username)
          alert(res.token + "\n")
          // 路由跳转
          navigate('/user/postShow')
        }).catch((err) => {alert(err.message)})
      }}
      >
        登录
      </Button>
      </Space>
     </div>
    </>
  )
}

export default LoginPage
