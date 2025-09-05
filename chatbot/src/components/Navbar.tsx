import React from "react";
import { Link } from "react-router-dom";

const Navbar: React.FC = () => {
  return (
    <nav className="bg-gray-100 p-4 flex gap-4 shadow">
      <Link
        to="/"
        className="text-blue-600 font-semibold hover:text-blue-800"
      >
        Home
      </Link>
      <Link
        to="/chat"
        className="text-blue-600 font-semibold hover:text-blue-800"
      >
        Chat
      </Link>
    </nav>
  );
};

export default Navbar;
