import React from "react";
import BackButton from "../button/BackButton";
import SearchBar from "./SearchBar";

const TopNavbar = () => {
  return (
    <div className="flex gap-2">
      <BackButton />
      <SearchBar />
    </div>
  );
};

export default TopNavbar;
