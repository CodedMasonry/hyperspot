import Layout from "./layout";
import { HashRouter, Routes, Route } from "react-router-dom";
//import { IsLoggedIn } from "../wailsjs/go/main/App";
import Login from "./pages/login";
import Home from "./pages/home";
import { useState } from "react";

function App() {
  //
  const [isAuthenticated, setAuthenticated] = useState(false);

  if (true) {
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
    return <Login />;
  }
}

export default App;
