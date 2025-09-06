import React from "react";
import { Link } from "react-router-dom";

const Home: React.FC = () => {
  return (
    <div className="min-h-screen flex flex-col items-center justify-center bg-gray-900 text-white p-4">
      <h1 className="text-5xl font-bold mb-4">23</h1>
      <p className="text-xl mb-6">Welcome</p>
      <Link
        to="/chat"
        className="bg-blue-600 hover:bg-blue-700 px-6 py-3 rounded-lg text-white"
      >
        Let's Chat!
      </Link>
    </div>
  );
};

export default Home;
