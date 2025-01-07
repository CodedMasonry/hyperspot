import Layout from "./layout";
import { HashRouter, Routes, Route } from "react-router-dom";
import { AuthenticateUser } from "../wailsjs/go/main/App";
import Home from "./pages/home";
import { useState } from "react";

function App() {
  const [isAuthenticated, setAuthenticated] = useState(false);

  if (isAuthenticated) {
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
    return (
      <button
        onClick={async () => {
          // Tries to authenticate
          setAuthenticated(await AuthenticateUser());
        }}
      >
        login
      </button>
    );
  }
}

export default App;
