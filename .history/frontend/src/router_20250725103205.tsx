import LoginPage from "./loginPage";
import InfoPage from "./infoPage";
import { createBrowserRouter } from "react-router-dom";

export const router = createBrowserRouter(
   [
        {
            path: "/user/login",
            element: <LoginPage/>
        },

        {
            path: "/user/info",
            element: <InfoPage/>,
        }
   ]
)