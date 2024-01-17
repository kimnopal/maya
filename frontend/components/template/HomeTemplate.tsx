import Navbar from "@/components/navigation/Navbar";
import SearchBar from "@/components/navigation/SearchBar";
import React from "react";

const HomeTemplate = ({ children }: { children: React.ReactNode }) => {
  return (
    <>
      <div className="flex flex-col gap-6 px-6 pt-6 py-24">
        <SearchBar />
        {children}
      </div>
      <Navbar />
    </>
  );
};

export default HomeTemplate;
