import "./App.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./pages/Home";
import Profile from "./pages/Profile";
import About from "./pages/About";
import Nav from "./components/Nav";
import PrivateRoute from "./routes/PrivateRoutes";
import { useAuth } from "./context/AuthContext";

function App() {
  const { isAuthenticated } = useAuth();
  return (
    <>
      <Router>
        <Nav></Nav>
        <Routes>
          {/* Public Routes */}
          <Route path="/" element={<Home />} />
          <Route path="/about" element={<About />} />
          <Route path="/profile" element={<Profile />} />
          <Route path="/" element={<Home />}></Route>
          {/* Private Routes */}
          <Route element={<PrivateRoute isAuthenticated={isAuthenticated} />}>
            <Route path="/checkout" element={<Home />} />
          </Route>
        </Routes>
      </Router>
    </>
  );
}

export default App;
