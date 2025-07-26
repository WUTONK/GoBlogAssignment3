import { useState } from 'react'
import './app.css'
import { Api } from './shared'
import Input from 'antd/es/input/Input'
import { useNavigate } from 'react-router-dom'

function PostModify() {
  const [post, setPosts] = useState("")

  return (
    <div className="post-modify-container">
      <h2 className="post-modify-title">报文管理</h2>
      {/* 输入框和提交按钮同一行 */}
      <div className="post-modify-row">
        <Input
          className="post-modify-input"
          type="text"
          value={post}
          onChange={(e: React.ChangeEvent<HTMLInputElement>) => setPosts(e.target.value)}
          placeholder="请输入报文"
        />
        <button
          onClick={() => { }}
          className="post-modify-submit-btn"
        >
          提交本条报文
        </button>
      </div>

      {/* 横线和说明文字 */}
      <div className="post-modify-divider">
        <hr />
        <span>其他功能</span>
        <hr />
      </div>

      {/* 下面两个按钮 */}
      <div className="post-modify-btn-row">
        <button className="post-modify-btn">清除最后一条报文</button>
        <button className="post-modify-btn">清除所有报文</button>
      </div>
    </div>
  )
}

export default PostModify
