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
          style={{
            flex: 1,
            minWidth: 120, // 最小宽度
            padding: "8px 0",
            whiteSpace: "nowrap", // 不换行
            fontSize: 18
          }}
        >
          提交本条报文
        </button>
      </div>

      {/* 横线和说明文字 */}
      <div style={{ display: "flex", alignItems: "center", margin: "24px 0 16px 0" }}>
        <hr style={{ flex: 1, border: "none", borderTop: "1px solid #888" }} />
        <span style={{ margin: "0 12px", color: "#888" }}>其他功能</span>
        <hr style={{ flex: 1, border: "none", borderTop: "1px solid #888" }} />
      </div>

      {/* 下面两个按钮 */}
      <div style={{ display: "flex", justifyContent: "center", gap: 16 }}>
       
        <button
          onClick={() => { }}
          style={{
            minWidth: 120,
            fontSize: 18,
            whiteSpace: "nowrap",
            padding: "8px 0"
          }}
        >
          清除最后一条报文
        </button>
        
        <button
          onClick={() => { }}
          style={{
            minWidth: 120,
            fontSize: 18,
            whiteSpace: "nowrap",
            padding: "8px 0"
          }}
        >
          清除所有报文
        </button>
      </div>
    </div>
  )
}

export default PostModify
