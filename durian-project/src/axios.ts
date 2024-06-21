import axios from "axios";

const axiosInstance = axios.create({
  baseURL: "http://localhost:10500/",
});

export default axiosInstance;
