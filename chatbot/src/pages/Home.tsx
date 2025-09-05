import React from "react";
import Navbar from "../components/Navbar";

const Home: React.FC = () => {
  return (
    <div className="p-4">
        <Navbar />
      <h1 className="text-2xl font-bold">Home Page</h1>
      <p>Welcome to the home page!</p>
    </div>
  );
};

export default Home;
