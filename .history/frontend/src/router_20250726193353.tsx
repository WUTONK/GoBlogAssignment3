import LoginPage from "./loginPage";
import PostModify from "./PostModifyPage";
import PostShowPage from "./PostShowPage";
import { createBrowserRouter } from "react-router-dom";

export const router = createBrowserRouter(
   [
        {
            path: "/user/login",
            element: <LoginPage/>
        },

        {
            path: "/user/postmodify",
            element: <PostModify/>
        },

        {
            path: "/user/postshow",
            element: <PostShowPage/>,
        }
   ]
)