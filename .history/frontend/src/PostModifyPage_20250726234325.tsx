import { useEffect, useState } from 'react'
import './app.css'
import { Api } from './shared'
import Input from 'antd/es/input/Input'
import { useNavigate } from 'react-router-dom'

function PostModify() {
  const navigate = useNavigate() //路由
  const [post, setPosts] = useState("")

  return (
    <>
      <div style={{ maxWidth: 400, margin: "0 auto", padding: 20 }}>
        <h2>报文管理</h2>
        <Input
          type="text"
          value={input}
          onChange={(e: React.ChangeEvent<HTMLInputElement) => setInput(e.target.value)}
          placeholder="请输入报文"
          style={{ width: "100%", marginBottom: 10, padding: 8 }}
        />
        <div style={{ marginBottom: 10 }}>
          <button onClick={handleSubmit} style={{ marginRight: 8 }}>提交本条报文</button>
          <button onClick={handlePop} style={{ marginRight: 8 }}>清除最后一条报文</button>
          <button onClick={handleClear}>清除所有报文</button>
        </div>
        <div>
          <h4>已提交报文：</h4>
          {posts.length === 0 ? (
            <div>暂无报文</div>
          ) : (
            posts.map((msg, idx) => (
              <div key={idx}>{msg}</div>
            ))
          )}
        </div>
      </div>
    </>
  )
}

export default PostModify
