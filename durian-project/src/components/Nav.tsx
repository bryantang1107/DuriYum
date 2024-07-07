import { Link } from "react-router-dom";

const Nav = () => {
  return (
    <div className="sticky top-0 h-16 grid grid-cols-5 gap-4 bg-gray-600 p-4">
      <div className="col-span-1 bg-red-400">
        <h1>Logo</h1>
      </div>
      <div className="col-span-3 bg-green-500">
        <div className="flex justify-between">
          <Link to="/">Home</Link>
          <Link to="/about">About Us</Link>
          <Link to="/shop">Shop</Link>
          <Link to="/contact">Contact</Link>
          <Link to="/faq">FAQ</Link>
        </div>
      </div>
      <div className="col-span-1 bg-blue-300 flex justify-between">
        <h1>Cart</h1>
        <h1>Profile</h1>
      </div>
    </div>
  );
};

export default Nav;
