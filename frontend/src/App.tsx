import Layout from "./layout";
import { HashRouter, Routes, Route } from "react-router-dom";
import { IsAuthenticated } from "../wailsjs/go/main/App";
import Home from "./pages/home";
import { useEffect, useState } from "react";
import Login from "./pages/login";

function App() {
  const [isLoggedIn, setLoggedIn] = useState(false);

  // Handle initial check for login
  useEffect(() => {
    async function init() {
      setLoggedIn(await IsAuthenticated());
    }
    init();
  }, []);

  if (isLoggedIn) {
    return (
      <Layout>
        <HashRouter basename={"/"}>
          <Routes>
            <Route path="/" element={<Home />} />
          </Routes>
        </HashRouter>
      </Layout>
    );
  } else {
    return <Login setLoggedIn={setLoggedIn} />;
  }
}

export default App;
