import Navbar from "@/components/navigation/Navbar";
import React from "react";
import TopNavbar from "../navigation/TopNavbar";

const BackTemplate = ({ children }: { children: React.ReactNode }) => {
  return (
    <>
      <div className="flex flex-col gap-6 px-6 pt-6 py-24">
        <TopNavbar />
        {children}
      </div>
      <Navbar />
    </>
  );
};

export default BackTemplate;
