import { useState } from 'react'
import './app.css'
import { Api } from './shared'
import Input from 'antd/es/input/Input'
import { useNavigate } from 'react-router-dom'

function PostModify() {
  const [post, setPosts] = useState("")

  return (
    <div style={{ maxWidth: 500, margin: "0 auto", padding: 20 }}>
      <h2 style={{ textAlign: "center" }}>报文管理</h2>
      {/* 输入框和提交按钮同一行 */}
      <div style={{ display: "flex", marginBottom: 16 }}>
        <Input
          type="text"
          value={post}
          onChange={(e: React.ChangeEvent<HTMLInputElement>) => setPosts(e.target.value)}
          placeholder="请输入报文"
          style={{ flex: 5, marginRight: 8, padding: 8 }}
        />
        <button
          onClick={() => { }}
          style={{ flex: 1, padding: "8px 0" }}
        >
          提交本条报文
        </button>
      </div>
      {/* 下面两个按钮一行 */}
      <div style={{ display: "flex", justifyContent: "center", gap: 16 }}>
        <button onClick={() => { }}>清除最后一条报文</button>
        <button onClick={() => { }}>清除所有报文</button>
      </div>
    </div>
  )
}

export default PostModify
