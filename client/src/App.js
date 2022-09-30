import logo from './logo.svg';
import './App.css';
import { API, setAuthToken } from './config/api';
import { useState } from "react";
import { useMutation } from "react-query";

function App() {
  const [token, setToken] = useState()
  if (localStorage.token) {
    setAuthToken(localStorage.token)
  }

  const getToken = async () => {
    try {
      const response = await API.get("/token")
      console.log(response.data);
      localStorage.setItem("token", response.data)
      setToken(response.data)
    } catch (error) {
      console.log(error);
    }
  }

  const [form, setForm] = useState({
    image: "",
  });

  

  const handleChange = (e) => {
    setForm({
      ...form,
      [e.target.name]:
        e.target.type === "file" ? e.target.files : e.target.value,
    });
  }

  const handleSubmit = useMutation(async (e) => {
    try {
      e.preventDefault();

      const config = {
        headers: {
          "Content-type": "multipart/form-data",
        },
      };

      // Store data with FormData as object
      const formData = new FormData();
      formData.set("image", form.image[0], form.image[0].name);
  

      // Configuration

      // Insert product data
      const response = await API.post("/data", formData, config);
      console.log(response);

    } catch (error) {
      console.log(error);
    }
  });
  

  return (
    <div className="App">
      <button onClick={getToken}>get token</button>
      <input style={{width: "200px"}} value={token} />
      <form onSubmit={(e) => handleSubmit.mutate(e)}>
        <input type="file" name='image' onChange={handleChange} />
        <button type='submit' >
          submit
          </button>
      </form>
    </div>
  );
}

export default App;
