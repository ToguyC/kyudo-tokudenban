import axios from 'axios'

const getAxiosInstance = (version: string = "1") => {
  return axios.create({
    baseURL: `http://localhost:8080/api/v${version}`,
  })
}

export default getAxiosInstance