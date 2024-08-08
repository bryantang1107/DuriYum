import "./App.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./pages/Home";
import Profile from "./pages/Profile";
import About from "./pages/About";
import Nav from "./components/Nav";
import SignIn from "./pages/SignIn";
import SignUp from "./pages/SignUp";
import PrivateRoute from "./routes/PrivateRoutes";
import Footer from "./components/Footer";
import Error from "./components/Error";
import { useAuth } from "./context/AuthContext";

function App() {
  const { isAuthenticated } = useAuth();
  return (
    <div className="flex flex-col h-screen">
      <Router>
        <div className="flex flex-col flex-grow">
          <Nav></Nav>
        </div>
        <main className="flex-grow">
          <Routes>
            {/* Public Routes */}
            <Route path="/" element={<Home />} />
            <Route path="/about" element={<About />} />
            <Route path="/profile" element={<Profile />} />
            <Route path="/login" element={<SignIn />}></Route>
            <Route path="/register" element={<SignUp />}></Route>
            {/* Private Routes */}
            <Route element={<PrivateRoute isAuthenticated={isAuthenticated} />}>
              <Route path="/checkout" element={<Home />} />
            </Route>
            <Route path="*" element={<Error />}></Route>
          </Routes>
        </main>
        <Footer></Footer>
      </Router>
    </div>
  );
}

export default App;
