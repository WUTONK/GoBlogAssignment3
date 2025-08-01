import { DefaultApi,Configuration } from "../api";

export const Api = new DefaultApi(new Configuration({
    basePath: import.meta.env.VITE_API_BASE_URL || "http://localhost:8080"
}))
