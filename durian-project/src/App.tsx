import "./App.css";
import { Routes, Route } from "react-router-dom";
import Home from "./components/Home";
import Profile from "./components/Profile";
import About from "./components/About";
import Nav from "./components/Nav";
import PrivateRoute from "./routes/PrivateRoutes";

function App() {
  return (
    //auth provide to wrap around
    //private route for super user
    <>
      <Routes>
        {/* Public Routes */}
        <Route path="/" element={<Home />} />
        <Route path="/about" element={<About />} />
        <Route path="/profile" element={<Profile />} />
        <Route path="/" element={<Home />}></Route>
        {/* Private Routes */}
        <PrivateRoute path="/checkout" element={<Home />} />
      </Routes>
      <div>
        <Nav></Nav>
      </div>
    </>
  );
}

export default App;
